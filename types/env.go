package types

type ENV struct {
	PORT           string `mapstructure:"PORT"`
	AppDbHost      string `mapstructure:"APP_DB_HOST"`
	AppDbPort      string `mapstructure:"APP_DB_PORT"`
	AppDbUsername  string `mapstructure:"APP_DB_USERNAME"`
	AppDbPassword  string `mapstructure:"APP_DB_PASSWORD"`
	AppDbName      string `mapstructure:"APP_DB_NAME"`
	AppDbSslMode   string `mapstructure:"APP_DB_SSL_MODE"`
	TestDbUsername string `mapstructure:"TEST_DB_USERNAME"`
	TestDbPassword string `mapstructure:"TEST_DB_PASSWORD"`
	TestDbName     string `mapstructure:"TEST_DB_NAME"`
}
