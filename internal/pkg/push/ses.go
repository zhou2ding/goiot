package push

import (
	"CMAIOT/internal/pkg/conf"
	"CMAIOT/internal/pkg/logger"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SESConfig struct {
	Region          string `mapstructure:"region"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
}

var (
	gSESClient *ses.Client
	gSESConfig SESConfig
)

func GetSESClient() *ses.Client {
	return gSESClient
}

func InitSESClient() error {
	ctx := context.Background()
	if err := conf.Conf.UnmarshalKey("aws", &gSESConfig); err != nil {
		logger.Logger.Errorf("unmarshal ses config from config file error: %v", err)
		return err
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     gSESConfig.AccessKeyID,
				SecretAccessKey: gSESConfig.SecretAccessKey,
			},
		}),
		config.WithRegion(gSESConfig.Region),
	)
	if err != nil {
		logger.Logger.Errorf("load ses config error: %v", err)
		return err
	}

	gSESClient = ses.NewFromConfig(cfg)

	_, err = gSESClient.ListIdentities(ctx, &ses.ListIdentitiesInput{})
	if err != nil {
		logger.Logger.Errorf("connect to ses error: %v", err)
		return err
	}

	return nil
}

func IsEmailVerified(ctx context.Context, sender string) (bool, error) {
	result, err := gSESClient.ListIdentities(ctx, &ses.ListIdentitiesInput{IdentityType: types.IdentityTypeEmailAddress})
	if err != nil {
		return false, err
	}

	for _, identity := range result.Identities {
		if identity == sender {
			return true, nil
		}
	}

	return false, nil
}

func VerifyEmailAddress(ctx context.Context, sender string) error {
	_, err := gSESClient.VerifyEmailIdentity(ctx, &ses.VerifyEmailIdentityInput{EmailAddress: aws.String(sender)})
	return err
}
