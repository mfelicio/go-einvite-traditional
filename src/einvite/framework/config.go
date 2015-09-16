package framework

import (
	"encoding/json"
	"os"
	"log"
)

//More info here: http://golang.org/doc/articles/json_and_go.html

type config struct {
	cfg interface{}
}

var Config = &config{}

func (this *config) Init(configPath string) {

	configFile, _ := os.Open(configPath)
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	var cfg interface{}
	err := decoder.Decode(&cfg)

	if err != nil {
		log.Fatalf("Failed to load configuration file at %s", configPath)
		panic(err)
	}

	this.cfg = cfg
}

func (this *config) ReadInto(configSection string, v interface{}) {

	if this.cfg == nil {
		panic("config Init has not been called")
	}

	var cfgMap = this.cfg.(map[string]interface{})

	section := cfgMap[configSection]
	sectionJson, err := json.Marshal(section)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(sectionJson, &v)

}
