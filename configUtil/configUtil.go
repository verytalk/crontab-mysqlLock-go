package configUtil

import (
	module "crontask/modile"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


var ConfigFile string;
var Config = new(module.Yaml)

func InitConfig() {
	// resultMap := make(map[string]interface{})
	conf := new(module.Yaml)
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	// err = yaml.Unmarshal(yamlFile, &resultMap)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf", conf)
	//log.Println(conf.Crontab.Period)
	//log.Println("conf", resultMap)
	Config = conf
}

