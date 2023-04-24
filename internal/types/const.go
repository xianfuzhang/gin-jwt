package types

import "time"

// export DATABASE_URL="postgres://postgres:postgres@192.168.58.10:5432/postgres?sslmode=require&pool_max_conns=10"
const (
	DB_host     = "localhost"
	DB_port     = 5432
	DB_user     = "postgres"
	DB_password = "123456"
	DB_name     = "postgres"
	DB_max_conn = 10
	DB_timeout  = 1 * time.Nanosecond

	UserResetPwd = "123456"
)
