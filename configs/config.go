package configs

import (
	"github.com/spf13/viper"
)

type configuration struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	AWSS3AccessKeyID     string  `mapstructure:"AWS_S3_ACCESS_KEY_ID"`
	AWSS3SecretAccessKey string  `mapstructure:"AWS_S3_SECRET_ACCESS_KEY"`
	AWSS3Region          string  `mapstructure:"AWS_S3_REGION"`
	AWSS3AppClientID     string  `mapstructure:"AWS_S3_APP_CLIENT_ID"`
	AWSS3BaseEndpoint    *string `mapstructure:"AWS_S3_BASE_ENDPOINT"`
	AWSS3BucketName      string  `mapstructure:"AWS_S3_BUCKET_NAME"`

	AWSSQSAccessKeyID     string  `mapstructure:"AWS_SQS_ACCESS_KEY_ID"`
	AWSSQSSecretAccessKey string  `mapstructure:"AWS_SQS_SECRET_ACCESS_KEY"`
	AWSSQSRegion          string  `mapstructure:"AWS_SQS_REGION"`
	AWSSQSAppClientID     string  `mapstructure:"AWS_SQS_APP_CLIENT_ID"`
	AWSSQSBaseEndpoint    *string `mapstructure:"AWS_SQS_BASE_ENDPOINT"`

	AWSCognitoAccessKeyID     string  `mapstructure:"AWS_COGNITO_ACCESS_KEY_ID"`
	AWSCognitoSecretAccessKey string  `mapstructure:"AWS_COGNITO_SECRET_ACCESS_KEY"`
	AWSCognitoRegion          string  `mapstructure:"AWS_COGNITO_REGION"`
	AWSCognitoAppClientID     string  `mapstructure:"AWS_COGNITO_APP_CLIENT_ID"`
	AWSCognitoUserPoolID      string  `mapstructure:"AWS_COGNITO_USER_POOL_ID"`

	QueueProcessVideo           string `mapstructure:"QUEUE_PROCESS_VIDEO"`
	DeadLetterQueueProcessVideo string `mapstructure:"DEAD_LETTER_QUEUE_PROCESS_VIDEO"`
}

func LoadConfig(path string, fileName string) (*configuration, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg *configuration
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
