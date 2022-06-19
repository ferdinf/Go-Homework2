package config

type Configuration struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   DB
}

type DB struct {
	User string `json:"ferdin"`
	Host string `json:"host"`
	Port string `json:"port"`
}
