package main

import (
	"configTest/config"
	"fmt"
)

func main() {
	configer, err := config.NewConfig("json", "db.json")
	if err != nil {
		fmt.Println("err:", err)
	}
	ip := configer.String("dev.master.ip")
	fmt.Println("IP:", ip)
}
