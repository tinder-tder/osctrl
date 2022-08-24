package logging

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jmpsec/osctrl/settings"
	"github.com/jmpsec/osctrl/types"
	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// LoggerS3 will be used to log data using S3
type LoggerS3 struct {
	S3Config  types.S3Configuration
	AWSConfig aws.Config
	Enabled   bool
}

// CreateLoggerS3 to initialize the logger
func CreateLoggerS3(s3Config types.S3Configuration) (*LoggerS3, error) {
	ctx := context.Background()
	creds := credentials.NewStaticCredentialsProvider(s3Config.AccessKey, s3Config.SecretAccessKey, "")
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(creds), config.WithRegion(s3Config.Region),
	)
	if err != nil {
		return nil, err
	}
	l := &LoggerS3{
		S3Config:  s3Config,
		AWSConfig: cfg,
		Enabled:   true,
	}
	return l, nil
}

// CreateLoggerS3File to initialize the logger with a filename
func CreateLoggerS3File(s3File string) (*LoggerS3, error) {
	s3Config, err := LoadS3(s3File)
	if err != nil {
		return nil, err
	}
	return CreateLoggerS3(s3Config)
}

// LoadS3 - Function to load the S3 configuration from JSON file
func LoadS3(file string) (types.S3Configuration, error) {
	var _s3Cfg types.S3Configuration
	log.Printf("Loading %s", file)
	// Load file and read config
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err != nil {
		return _s3Cfg, err
	}
	cfgRaw := viper.Sub(settings.LoggingS3)
	if err := cfgRaw.Unmarshal(&_s3Cfg); err != nil {
		return _s3Cfg, err
	}
	// No errors!
	return _s3Cfg, nil
}

// Settings - Function to prepare settings for the logger
func (logS3 *LoggerS3) Settings(mgr *settings.Settings) {
	log.Printf("No s3 logging settings\n")
}

// Send - Function that sends JSON logs to S3
func (logS3 *LoggerS3) Send(logType string, data []byte, environment, uuid string, debug bool) {
	ctx := context.Background()
	if debug {
		log.Printf("DebugService: Sending %d bytes to S3 for %s - %s", len(data), environment, uuid)
	}
	client := s3.NewFromConfig(logS3.AWSConfig)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(logS3.S3Config.Bucket),
		Key:           aws.String(logType + ":" + environment + ":" + uuid + ":" + strconv.FormatInt(time.Now().UnixMilli(), 10) + ".json"),
		Body:          bytes.NewBuffer(data),
		ContentLength: int64(len(data)),
		ContentType:   aws.String(http.DetectContentType(data)),
	})
	if err != nil {
		log.Printf("Error sending data to s3 %s", err)
	}
	if debug {
		log.Printf("DebugService: S3 Upload %+v", result)
	}
}
