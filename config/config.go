package config

type Configuration struct {
	HttpServer HttpServerConfig `mapstructure:"http-server"`
	GrpcServer GrpcServerConfig `mapstructure:"grpc-server"`
	// UserGrpc   GrpcServerConfig `mapstructure:"user-grpc"`
	// SurveyGrpc GrpcServerConfig `mapstructure:"survey-grpc"`
	Database DatabaseConfig `mapstructure:"database"`
	// Mongo    MongoConfig    `mapstructure:"mongo"`
	Secrect SecrectConfig `mapstructure:"secrect"`
}

type HttpServerConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

type GrpcServerConfig struct {
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type MongoConfig struct {
	MONGO_URI           string `mapstructure:"mongo_uri"`
	DB                  string `mapstructure:"db"`
	RESPONSE_COLLECTION string `mapstructure:"response_collection"`
}

type SecrectConfig struct {
	ACCESS_SECRET string `mapstructure:"access_secret"`
}

var Config Configuration
