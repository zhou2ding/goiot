package oss

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"goiot/internal/pkg/conf"
	"goiot/internal/pkg/logger"
)

type S3Config struct {
	Region          string `mapstructure:"region"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
}

// 从 2023 年 4 月开始，Amazon S3 将通过自动启用 S3 阻止公共访问和禁用所有新 S3 存储桶的 S3 访问控制列表 (ACL) 来引入两项新的默认存储段安全设置。因此修改桶权限不能通过直接设置ACL，需要设置桶策略。
type (
	BucketPolicy struct {
		Version   string       `json:"Version"`
		Statement []*Statement `json:"Statement"`
	}
	Statement struct {
		Effect    string   `json:"Effect"`
		Principal string   `json:"Principal,omitempty"`
		Action    []string `json:"Action"`
		Resource  []string `json:"Resource"`
	}
)

const (
	USBucketARN = "arn:aws:s3:::"
	CNBucketARN = "arn:aws-cn:s3:::"

	PackageBucket            = "goiot-package"       //aws的桶名全球唯一，因此需要结合账户id使用，下同
	DefaultUserBucket        = "goiot-default-user"  //用户上传文件时没有指定桶，统一放到默认桶
	DefaultDeviceImageBucket = "goiot-default-image" //设备上传图片时没有指定桶，统一放到默认桶
	DefaultDeviceVideoBucket = "goiot-default-video" //设备上传视频时没有指定桶，统一放到默认桶

	PrivateACL      = "private"
	PublicRead      = "public-read"
	PublicReadWrite = "public-read-write"

	readAction  = "s3:GetObject"
	writeAction = "s3:PutObject"
)

var (
	gS3Client        *s3.Client
	gS3PresignClient *s3.PresignClient
	gS3Config        S3Config
)

var (
	ImageSuffix = map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		".bmp": true, ".tiff": true, ".tif": true, ".webp": true,
	}
	VideoSuffix = map[string]bool{
		".mp4": true, ".mkv": true, ".avi": true, ".mov": true,
		".wmv": true, ".flv": true, ".webm": true, ".m4v": true, ".rmvb": true}
)

var eventMap = map[string][]types.Event{
	"upload": {
		types.EventS3ObjectCreatedPut,
		types.EventS3ObjectCreatedPost,
		types.EventS3ObjectCreatedCopy,
		types.EventS3ObjectCreatedCompleteMultipartUpload,
	},
	"delete": {
		types.EventS3ObjectRemovedDelete,
	},
}

func InitS3Client() error {
	ctx := context.Background()
	if err := conf.Conf.UnmarshalKey("aws", &gS3Config); err != nil {
		logger.Logger.Errorf("unmarshal s3 config from config file error: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     gS3Config.AccessKeyID,
				SecretAccessKey: gS3Config.SecretAccessKey,
			},
		}),
		config.WithRegion(gS3Config.Region),
	)
	if err != nil {
		logger.Logger.Errorf("load s3 config error: %v", err)
		return err
	}

	gS3Client = s3.NewFromConfig(cfg)
	gS3PresignClient = s3.NewPresignClient(gS3Client)

	_, err = gS3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		logger.Logger.Errorf("connect to s3 error: %v", err)
		return err
	}

	return nil
}

func GetS3Client() *s3.Client {
	return gS3Client
}

func GetS3PresignClient() *s3.PresignClient {
	return gS3PresignClient
}

func GetS3Config() *S3Config {
	return &gS3Config
}

func CreateBucket(ctx context.Context, bucket, acl string) error {
	create := s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(gS3Config.Region),
		},
	}
	if _, err := gS3Client.CreateBucket(ctx, &create); err != nil {
		return err
	}

	if acl != "" {
		if err := SetBucketACL(ctx, bucket, acl); err != nil {
			return err
		}
	}

	return nil
}

func SetBucketACL(ctx context.Context, bucket, acl string) error {
	var private bool
	if acl == PrivateACL {
		private = true
	}
	// 设置桶是否公共可访问
	_, err := gS3Client.PutPublicAccessBlock(ctx, &s3.PutPublicAccessBlockInput{
		Bucket: aws.String(bucket),
		PublicAccessBlockConfiguration: &types.PublicAccessBlockConfiguration{
			BlockPublicAcls:       aws.Bool(private),
			IgnorePublicAcls:      aws.Bool(private),
			BlockPublicPolicy:     aws.Bool(private),
			RestrictPublicBuckets: aws.Bool(private),
		},
	})
	if err != nil {
		return err
	}

	// 设置桶策略
	if !private {
		policy := BucketPolicy{
			Version:   "2012-10-17",
			Statement: make([]*Statement, 0),
		}
		statement := Statement{
			Effect:    "Allow",
			Principal: "*",
			Resource:  []string{fmt.Sprintf("%s%s/*", USBucketARN, bucket)},
		}
		switch acl {
		case PublicRead:
			statement.Action = []string{readAction}
		case PublicReadWrite:
			statement.Action = []string{readAction, writeAction}
		}
		policy.Statement = append(policy.Statement, &statement)

		policyJson, err := json.Marshal(policy)
		if err != nil {
			return err
		}
		_, err = gS3Client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
			Bucket: aws.String(bucket),
			Policy: aws.String(string(policyJson)),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func SetBucketNotice(ctx context.Context, bucket, queueArn, prefix, suffix string, events []string) error {
	queueCfg := types.QueueConfiguration{QueueArn: aws.String(queueArn)}
	for _, event := range events {
		queueCfg.Events = append(queueCfg.Events, eventMap[event]...)
	}

	var rules []types.FilterRule
	if prefix != "" {
		rules = append(rules, types.FilterRule{
			Name:  types.FilterRuleNamePrefix,
			Value: aws.String(prefix),
		})
	}
	if suffix != "" {
		rules = append(rules, types.FilterRule{
			Name:  types.FilterRuleNameSuffix,
			Value: aws.String(suffix),
		})
	}
	if len(rules) > 0 {
		queueCfg.Filter = &types.NotificationConfigurationFilter{
			Key: &types.S3KeyFilter{
				FilterRules: rules,
			},
		}
	}

	noticeCfg := s3.PutBucketNotificationConfigurationInput{
		Bucket: aws.String(bucket),
		NotificationConfiguration: &types.NotificationConfiguration{
			QueueConfigurations: []types.QueueConfiguration{queueCfg},
		},
	}

	_, err := gS3Client.PutBucketNotificationConfiguration(ctx, &noticeCfg)
	if err != nil {
		return err
	}
	return nil
}
