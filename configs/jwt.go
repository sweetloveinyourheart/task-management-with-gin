package configs

import "os"

var JwtSecret = os.Getenv("JWT_SECRET")
