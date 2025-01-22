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

	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSBucketName     string `mapstructure:"AWS_BUCKET_NAME"`
	AWSBaseEndpoint    *string `mapstructure:"AWS_BASE_ENDPOINT"`

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

