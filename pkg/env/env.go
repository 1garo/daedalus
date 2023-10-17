package env

import "github.com/joho/godotenv"

var Env map[string]string

func GetEnv(key, value string) string {
	if val, ok := Env[key]; ok {
		return val
	}

	return value
}

func SetupEnvFile() {
	var err error
	if Env, err = godotenv.Read(".env"); err != nil {
		panic(err)
	}
}
