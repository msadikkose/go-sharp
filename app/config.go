package app

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Port             string
	ConnectionString string
	DbDriver         string
	ReleaseMode      bool
	//RedisHost         string
	//RedisPassword     string
	//RedisDuration     int
	//LoggerDestination string
	//LoggerHost        string
	//LoggerToken       string
	//LogLevel          string

}

//AppConfig configuration

var (
	Config     *configuration
	configOnce sync.Once
)

//InitConfig initialize the config
func InitConfig() {
	if Config == nil {
		configOnce.Do(func() {
			loadAppConfig()
		})
	}

	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("appSettings.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = &configuration{}
	err = decoder.Decode(Config)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}
