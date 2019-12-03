package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when opeing env")
	}

	fmt.Printf("Environment: %s\n", os.Getenv("ENV"))
	fmt.Printf("DB username: %s\n", os.Getenv("DB_USER"))
	fmt.Printf("DB password: %s\n", os.Getenv("DB_PASS"))
	fmt.Printf("SECRET: %s\n", os.Getenv("SECRET"))
}