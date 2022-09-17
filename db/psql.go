package db_config

type psql_config struct {
	PORT     string
	HOST     string
	PASSWORD string
}

func NewPsqlConfig(host string, port string, password string) *psql_config {
	return &psql_config{
		PORT:     port,
		PASSWORD: password,
		HOST:     host,
	}
}
