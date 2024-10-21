package config

import "github.com/spf13/viper"

var AppConfig Config

type Config struct {
	Port int `mapstructure:"PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisExpires  int    `mapstructure:"REDIS_EXPIRES"`

	JWTSecret         string `mapstructure:"JWT_SECRET"`
	JWTIssuer         string `mapstructure:"JWT_ISSUER"`
	JWTExpires        int    `mapstructure:"JWT_EXPIRES"`
	JWTRefreshExpires int    `mapstructure:"JWT_REFRESH_EXPIRES"`

	MobizonBaseUrl string `mapstructure:"MOBIZON_BASE_URL"`
	MobizonApiKey  string `mapstructure:"MOBIZON_API_KEY"`

	AwsAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsAccessKeySecret string `mapstructure:"AWS_ACCESS_KEY_SECRET"`
	AwsBucketName      string `mapstructure:"AWS_BUCKET_NAME"`
	AwsBucketRegion    string `mapstructure:"AWS_BUCKET_REGION"`

	OneCBaseUrl  string `mapstructure:"ONE_C_BASE_URL"`
	OneCUsername string `mapstructure:"ONE_C_USERNAME"`
	OneCPassword string `mapstructure:"ONE_C_PASSWORD"`

	CdekBaseUrl      string `mapstructure:"CDEK_BASE_URL"`
	CdekGrantType    string `mapstructure:"CDEK_GRANT_TYPE"`
	CdekClientId     string `mapstructure:"CDEK_CLIENT_ID"`
	CdekClientSecret string `mapstructure:"CDEK_CLIENT_SECRET"`

	ForteBaseUrl  string `mapstructure:"FORTE_BASE_URL"`
	ForteUsername string `mapstructure:"FORTE_USERNAME"`
	FortePassword string `mapstructure:"FORTE_PASSWORD"`
}

func InitializeAppConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
