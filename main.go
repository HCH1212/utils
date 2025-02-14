package main

import (
	"fmt"
	"github.com/HCH1212/utils/conf"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()
	fmt.Println(conf.GetConf().MySQL.DSN)
	fmt.Println(os.Getenv("GO_ENV"))
}
