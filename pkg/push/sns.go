package push

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
)

type SNSConfig struct {
	Region             string `mapstructure:"region"`
	AccessKeyID        string `mapstructure:"accessKeyID"`
	SecretAccessKey    string `mapstructure:"secretAccessKey"`
	IOSTopicArn        string `mapstructure:"iosTopicArn"`
	IOSPlatformArn     string `mapstructure:"iosPlatformArn"`
	AndroidTopicArn    string `mapstructure:"androidTopicArn"`
	AndroidPlatformArn string `mapstructure:"androidPlatformArn"`
}

const (
	ApplePlatform   = "ios"
	AndroidPlatform = "android"
)

var (
	gSNSClient *sns.Client
	gSNSConfig SNSConfig
)

func InitSnsClient() error {
	ctx := context.Background()
	if err := conf.Conf.UnmarshalKey("aws", &gSNSConfig); err != nil {
		logger.Logger.Errorf("unmarshal SNS config from config file error: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     gSNSConfig.AccessKeyID,
				SecretAccessKey: gSNSConfig.SecretAccessKey,
			},
		}),
		config.WithRegion(gSNSConfig.Region),
	)
	if err != nil {
		logger.Logger.Errorf("load SNS config error: %v", err)
		return err
	}

	gSNSClient = sns.NewFromConfig(cfg)

	_, err = gSNSClient.ListTopics(ctx, &sns.ListTopicsInput{})
	if err != nil {
		logger.Logger.Errorf("connect to SNS error: %v", err)
		return err
	}

	return nil
}

func GetSNSClient() *sns.Client {
	return gSNSClient
}

func SendSNSMessage(ctx context.Context, deviceType, endPointArn string, msg any) error {
	input := sns.PublishInput{
		MessageStructure: aws.String("json"),
		TargetArn:        aws.String(endPointArn),
	}

	snsMsg := map[string]any{}
	if deviceType == ApplePlatform {
		snsMsg["APNS"] = msg
	} else if deviceType == AndroidPlatform {
		snsMsg["GCM"] = msg
	}

	msgJson, err := json.Marshal(snsMsg)
	if err != nil {
		logger.Logger.Errorf("marshal sns message error: %v", err)
		return err
	}
	input.Message = aws.String(string(msgJson))

	_, err = gSNSClient.Publish(ctx, &input)
	if err != nil {
		return err
	}
	return nil
}
