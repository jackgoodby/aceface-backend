package config

import (
	"net/url"
	"sync"
	"time"

	"github.com/jackgoodby/aceface-backend/internal/common"
)

type ConfigStruct struct {
	Region string

	// Postgres
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	// Database related
	PrimaryTableName     string
	EmailSignInSortKey   string
	ResetPasswordSortKey string
	UserSortKey          string
	VerificationSortKey  string
	RefreshTokenSortKey  string
	VerificationTtl      time.Duration // To be used while adding to time

	// S3 related
	PrimaryBucketName string

	// SNS related
	PrimaryTopicArn string

	// Config from config.json
	CorsAllowOriginHeader      string // This is included in the config but is primarily used in the lambdaHelpers.LamdbaWrapper func
	TokenIssuer                string
	TokenExpirationTime        time.Duration // To be used while adding to time
	RefreshTokenExpirationTime time.Duration
	SourceEmailAddress         string
}

var Config *ConfigStruct
var onceConfig sync.Once

func GetConfig() *ConfigStruct {
	onceConfig.Do(func() {
		verificationTtl, verTtlParseErr := time.ParseDuration("15m")
		if verTtlParseErr != nil {
			panic(verTtlParseErr)
		}

		tokenExpiration, tokenExpParseErr := time.ParseDuration(common.GetEnv("TOKEN_EXPIRATION_TIME", "6h"))
		if tokenExpParseErr != nil {
			panic(tokenExpParseErr)
		}

		refreshTokenExpiration, refreshTokenExpParseErr := time.ParseDuration(common.GetEnv("REFRESH_TOKEN_EXPIRATION_TIME", "720h")) // 30 days
		if refreshTokenExpParseErr != nil {
			panic(refreshTokenExpParseErr)
		}

		Config = &ConfigStruct{
			Region:                     common.GetEnv("AWS_REGION", "eu-west-2"),
			DBHost:                     common.GetEnv("DB_HOST", ""),
			DBPort:                     common.GetEnv("DB_PORT", ""),
			DBUser:                     common.GetEnv("DB_USER", ""),
			DBPass:                     url.QueryEscape(common.GetEnv("DB_PASS", "")),
			DBName:                     common.GetEnv("DB_NAME", ""),
			PrimaryTableName:           common.GetEnv("PRIMARY_TABLE_NAME", ""),
			EmailSignInSortKey:         "email",
			ResetPasswordSortKey:       "reset",
			UserSortKey:                "user",
			VerificationSortKey:        "verification",
			RefreshTokenSortKey:        "refreshToken",
			VerificationTtl:            verificationTtl,
			PrimaryBucketName:          common.GetEnv("PRIMARY_BUCKET_NAME", ""),
			PrimaryTopicArn:            common.GetEnv("PRIMARY_SNS_TOPIC_ARN", ""),
			CorsAllowOriginHeader:      common.GetEnv("CORS_ALLOW_ORIGIN_HEADER", ""),
			TokenIssuer:                common.GetEnv("TOKEN_ISSUER", ""),
			TokenExpirationTime:        tokenExpiration,
			RefreshTokenExpirationTime: refreshTokenExpiration,
			SourceEmailAddress:         common.GetEnv("SOURCE_EMAIL_ADDRESS", ""),
		}
	})
	return Config
}
