package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongo() string {
    err := godotenv.Load()
	
    if err != nil {
        log.Fatal("Error loading enviroment file")
    }
  
    return os.Getenv("MONGOURI")
}