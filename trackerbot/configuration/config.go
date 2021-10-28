package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Prefix 		string 	`json:"prefix"`
	Token 		string 	`json:"token"`
	OwnerID 	string 	`json:"owner_id"`
	UseSharding bool 	`json:"use_sharding"`
	ShardID 	int 	`json:"shard_id"`
}

func LoadConfig(fileName string) *Config {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("error loading config,", err)
		return nil
	}
	var conf Config

	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("error loading json in %v,", fileName), err)
		return nil
	}

	return &conf
}
