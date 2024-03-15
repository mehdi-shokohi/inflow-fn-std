package conf

import (
	"os"

	"github.com/joho/godotenv"
)


	func GetPort() (port string) {
		port = Get_env_value("PORT")
		if port == "" {
			return "0.0.0.0:80"
		}
		return
	}


	func Get_env_value(key string) string {
		env:=".env"
		if appEnv:=os.Getenv("ENV");appEnv!=""{
			env=env+"."+appEnv
		}
	
		godotenv.Load(env)

		return os.Getenv(key)
	}