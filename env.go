/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/6 上午10:02
 * @note:
 */

package tuna

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

// GetEnv get environment variables
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
