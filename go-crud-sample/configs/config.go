package configs

type Config struct{
	DBDial string
}

func NewConfig()*Config{
	return &Config{
		DBDial: "server_user:server_pwd@tcp(localhost:3306)/server_db?parseTime=true",
	}
}