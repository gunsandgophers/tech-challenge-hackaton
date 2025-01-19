package config

var (
	AWS_ACCESS_KEY_ID     = GetEnv("AWS_ACCESS_KEY_ID", "000000000000")
	AWS_SECRET_ACCESS_KEY = GetEnv("AWS_SECRET_ACCESS_KEY", "000000000000")
	AWS_REGION            = GetEnv("AWS_REGION", "us-east-1")
	AWS_BUCKERT_NAME      = GetEnv("AWS_BUCKERT_NAME", "tech-challenge-hackaton")
	AWS_BASE_ENDPOINT     = GetEnvOrNil("AWS_BASE_ENDPOINT")
)
