package config

type App struct {
	JwtSecret string
	AppUrl    string
	RunMode   string
	HttpPort  int
	ENV       string
}

var AppConfig = &App{}

type Database struct {
	Dialect         string
	User            string
	Password        string
	Host            string
	Name            string
	Port            int
	ConnMaxLifeTime int
}

var DBConfig = &Database{}

type Redis struct {
	Host     string
	Password string
}

var RedisConfig = &Redis{}
