package confpack

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port        int      `yaml:"port"`
	DbURL       string   `yaml:"db_url"`
	JaegerURL   string   `yaml:"jaeger_url"`
	SentryURL   string   `yaml:"sentry_url"`
	KafkaBroker string   `yaml:"kafka_broker"`
	SomeAppID   string   `yaml:"some_app_id"`
	SomeAppKey  string   `yaml:"some_app_key"`
	SomeKeys    []string `yaml:"some_keys"`
	DeepSetting []struct {
		SubSett struct {
			Key1 string `yaml:"key1"`
			Key2 string `yaml:"key2"`
		} `yaml:"sub_sett"`
	} `yaml:"deep_setting"`
}

func GetConfigFromFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't open file: %v", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file: %v", err)
		}
	}()

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	//fmt.Println(string(bs))
	config := Config{}

	// Unmarshalling
	err = yaml.Unmarshal(bs, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Validation
	err = nil
	if config.Port <= 0 || 65535 < config.Port {
		err = errors.New("port must be in a range 1..65535")
	}
	return config, err
	//fmt.Println(config)
}
