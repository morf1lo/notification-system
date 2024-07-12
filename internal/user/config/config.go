package config

type GRPCServerConfig struct {
	Network string
	Addr    string
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}
