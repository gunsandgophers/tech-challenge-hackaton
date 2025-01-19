package config

var (
	DB_HOST     = GetEnv("DB_HOST", "db")
	DB_POST     = GetEnv("DB_PORT", "5432")
	DB_USER     = GetEnv("DB_USER", "tech-challenge-hackaton")
	DB_PASSWORD = GetEnv("DB_PASSWORD", "tech-challenge-hackaton")
	DB_NAME     = GetEnv("DB_NAME", "tech-challenge-hackaton")
)
