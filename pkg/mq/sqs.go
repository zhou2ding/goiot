package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
)

type SqsConfig struct {
	Region          string `mapstructure:"region"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	UserId          string `mapstructure:"userId"`
}

type (
	SQSPolicy struct {
		Version   string       `json:"Version"`
		Statement []*Statement `json:"Statement"`
	}
	Statement struct {
		Effect    string     `json:"Effect"`
		Principal any        `json:"Principal,omitempty"` // 支持json格式，也支持通配符格式
		Action    []string   `json:"Action"`
		Resource  []string   `json:"Resource"`
		Condition *Condition `json:"Condition,omitempty"`
	}
	Principal struct {
		AWS []string `json:"AWS,omitempty"`
	}
	Condition struct {
		ArnEquals map[string]string `json:"ArnEquals,omitempty"`
	}
)

var (
	gSqsClient  *sqs.Client
	gSqsConfig  SqsConfig
	queueUrlMap map[string]string
)

func InitSqsClient() error {
	ctx := context.Background()
	queueUrlMap = make(map[string]string)
	if err := conf.Conf.UnmarshalKey("aws", &gSqsConfig); err != nil {
		logger.Logger.Errorf("unmarshal sqs config from config file error: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     gSqsConfig.AccessKeyID,
				SecretAccessKey: gSqsConfig.SecretAccessKey,
			},
		}),
		config.WithRegion(gSqsConfig.Region),
	)
	if err != nil {
		logger.Logger.Errorf("load sqs config error: %v", err)
		return err
	}

	gSqsClient = sqs.NewFromConfig(cfg)

	_, err = gSqsClient.ListQueues(ctx, &sqs.ListQueuesInput{})
	if err != nil {
		logger.Logger.Errorf("connect to sqs error: %v", err)
		return err
	}

	return nil
}

func GetSqsClient() *sqs.Client {
	return gSqsClient
}

func NormalQueueExists(ctx context.Context, queueName, s3BucketArn string) (string, error) {
	resp, err := gSqsClient.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		var notExist *types.QueueDoesNotExist
		if errors.As(err, &notExist) {
			createResp, createErr := gSqsClient.CreateQueue(ctx, &sqs.CreateQueueInput{
				QueueName: &queueName,
			})
			if createErr != nil {
				return "", createErr
			}

			queueArn, err := getQueueArn(ctx, createResp.QueueUrl)
			if err != nil {
				return "", err
			}

			policy := SQSPolicy{
				Version: "2012-10-17",
				Statement: []*Statement{
					{
						Effect:    "Allow",
						Action:    []string{"sqs:SendMessage"},
						Resource:  []string{queueArn},
						Principal: "*",
						Condition: &Condition{
							ArnEquals: map[string]string{
								"aws:SourceArn": s3BucketArn,
							},
						},
					},
					{
						Effect:    "Allow",
						Action:    []string{"sqs:ReceiveMessage"},
						Resource:  []string{queueArn},
						Principal: "*",
					},
				},
			}
			policyJSON, err := json.Marshal(policy)
			if err != nil {
				return "", err
			}

			_, err = gSqsClient.SetQueueAttributes(ctx, &sqs.SetQueueAttributesInput{
				QueueUrl: createResp.QueueUrl,
				Attributes: map[string]string{
					"Policy": string(policyJSON),
				},
			})
			if err != nil {
				return "", err
			}

			return queueArn, nil
		}
		return "", err
	}

	return getQueueArn(ctx, resp.QueueUrl)
}

func getQueueArn(ctx context.Context, queueUrl *string) (string, error) {
	attrs, err := gSqsClient.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
		QueueUrl:       queueUrl,
		AttributeNames: []types.QueueAttributeName{"QueueArn"},
	})
	if err != nil {
		return "", fmt.Errorf("error retrieving sqs queue attributes: %w", err)
	}
	return attrs.Attributes["QueueArn"], nil
}

func FIFOQueueExists(ctx context.Context, client *sqs.Client, queueName string) (string, error) {
	resp, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		var notExist *types.QueueDoesNotExist
		if errors.As(err, &notExist) {
			createResp, createErr := client.CreateQueue(ctx, &sqs.CreateQueueInput{
				QueueName: &queueName,
				Attributes: map[string]string{
					"FifoQueue": "true",
				},
				Tags: nil,
			})
			if createErr != nil {
				return "", createErr
			}
			queueUrlMap[queueName] = *resp.QueueUrl
			queueArn, err := getQueueArn(ctx, createResp.QueueUrl)
			if err != nil {
				return "", err
			}
			return queueArn, nil
		}
		return "", err
	}
	queueUrlMap[queueName] = *resp.QueueUrl
	return getQueueArn(ctx, resp.QueueUrl)
}

func GetQueueUrl(queueName string) string {
	return queueUrlMap[queueName]
}
