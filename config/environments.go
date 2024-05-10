package config

import "os"

var DbStringConnection = os.Getenv("DB_POSTGRES_CONNECTION")
