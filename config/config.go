package config

type IParams struct {
	ParseTime string
}

type IDatabase struct {
	User     string
	Password string
	Net      string
	Addr     string
	DBName   string
	Params   IParams
}

type IServer struct {
	Address string
}

type IConfig struct {
	Database IDatabase
	Server   IServer
}

var C IConfig = IConfig{
	Database: IDatabase{
		User:     "gustavo",
		Password: "123456789",
		Net:      "localhost",
		Addr:     "3306",
		DBName:   "prueba",
		Params: IParams{
			ParseTime: "True",
		},
	},
	Server: IServer{
		Address: "8080",
	},
}
