package main

import (
	"crontask/configUtil"
	"crontask/mysqlUtil"
	"flag"
	"fmt"
	"gopkg.in/robfig/cron.v2"
	"log"
)

func task() {
	log.Println("running ...")
	//fmt.Println(configUtil.Config.Crontab.Period)
	mysqlUtil.Run()
}


func main() {


	config := flag.String("config", "./config.yaml", "the config")
	flag.Parse()

	fmt.Println("config:", *config)


	configUtil.ConfigFile = *config
	configUtil.InitConfig()

	crontab := cron.New()
	crontab.AddFunc(configUtil.Config.Crontab.Period, task)
	crontab.Start()
	select {}
}
