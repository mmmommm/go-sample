package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// 本番環境ではk8sのconfingmapで環境変数をいれる
type Config struct {
	Env string `envconfig:"ENV" default:"local"`

	ServerHost string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	ServerPort int    `envconfig:"SERVER_PORT" default:"9090"`

	MLServerHost string `envconfig:"ML_SERVER_HOST" default:"localhost"`
	MLServerPort string `envconfig:"ML_SERVER_PORT" default:":8000"`

	DBEngine                string `envconfig:"DB_ENGINE" default:"mysql"`
	DBUser                  string `envconfig:"DB_USER" default:"sampleuser"`
	DBPass                  string `envconfig:"DB_PASSWORD" default:"password"`
	DBHost                  string `envconfig:"DB_ADDR" default:"dev_db"`
	DBPort                  int    `envconfig:"DB_PORT" default:"3306"`
	DBName                  string `envconfig:"DB_NAME" default:"dev_db"`
	RedisAddress            string `envconfig:"REDIS_ADDRESS" default:"dev_redis"`
	RedisPASS               string `envconfig:"REDIS_PASSWORD" default:""`
	RedisPort               int    `envconfig:"REDIS_PORT" default:"6379"`
	RedisClusterModeEnabled bool   `envconfig:"REDIS_CLUSTER_MODE_ENABLED" default:"false"`

	SQSRegion   string `envconfig:"SQS_REGION" default:"ap-northeast-1"`
	SQSEndPoint string `envconfig:"SQS_ENDPOINT" default:"http://dev-sqs:9324"`
	SQSQueueURL string `envconfig:"SQS_QUEUE_URL" default:"http://dev-sqs:9324/000000000000/timeline_update_queue"`
	// dbから受け取ったDateTimeをGo側でどのタイムゾーンで表示するか
	// とりあえずTokyoに
	DBTimeZone string `envconfig:"DB_TZ" default:"Asia%2FTokyo"`

	ImgS3BucketName   string `envconfig:"IMG_S3_BUCKET_NAME" default:""`
	ImgS3BucketRegion string `envconfig:"IMG_S3_BUCKET_REGION" default:"ap-northeast-1"`
	ImgBaseURI        string `envconfig:"IMG_BASE_URI" default:"https://media.22fresh-c.ca-developers.io"`
}

func ProvideConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("go-sample", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c Config) ServerAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort)
}

func (c Config) DBAddr() string {
	return fmt.Sprintf("%s:%d", c.DBHost, c.DBPort)
}

// func (c Config) RedisAddr() string {
// 	return fmt.Sprintf("%s:%d", c.RedisAddress, c.RedisPort)
// }

// func (c Config) MLServerAddr() string {
// 	return fmt.Sprintf("%s%s", c.MLServerHost, c.MLServerPort)
// }
