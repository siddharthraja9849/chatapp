package types

type ENV struct {
	AppDbUsername  string `mapstructure:"APP_DB_USERNAME"`
	AppDbPassword  string `mapstructure:"APP_DB_PASSWORD"`
	AppDbName      string `mapstructure:"APP_DB_NAME"`
	TestDbUsername string `mapstructure:"TEST_DB_USERNAME"`
	TestDbPassword string `mapstructure:"TEST_DB_PASSWORD"`
	TestDbName     string `mapstructure:"TEST_DB_NAME"`
}
