package postgresql

type Config struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	Dbname                string
	MaxConnections        string
	MaxConnectionIdleTime string
	UserName              any
	DbName                any
}
