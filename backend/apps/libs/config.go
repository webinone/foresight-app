package libs

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"flag"
)

type Configuration struct {
	MQTT 	mqttConfig
	REDIS   redisConfig
	INFO 	infoConfig
	RDB	rdbConfig
	AUTH	authConfig
}

type mqttConfig struct {
	LocalUrl string 	`json:"localUrl"`
	KfarmUrl string 	`json:"kfarmUrl"`
}

type redisConfig struct {
	Url		string 	`json:"url"`
	Password	string 	`json:"password"`
	Db		int 	`json:"db"`
}

type infoConfig struct {
	PrjctNo string 		`json:"prjctNo"`
	EndpntId string 	`json:"endpntId"`
	NodeId   int64		`json:"nodeId"`
}

type rdbConfig struct {
	Product string		`json:"product"`
	ConnectString string	`json:"connect_string"`
	Debug 	bool		`json:"debug"`
	Migrate bool		`json:"migrate"`
}

type authConfig struct {
	AuthKey		string 		`json:"auth_key"`
	JwtKey		string 		`json:"jwt_key"`
}

//var ConfigRoot string
var Config  = &Configuration{}

func LoadPathConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}
}

func LoadAutoConfig () {

	configRoot := flag.String("mode", "foo", "development mode")
	flag.Parse()

	var path string
	// 개발 환경 셋팅
	if *configRoot == "foo" {
		path = "app-dev.json"
	} else if *configRoot == "dev" {
		//ConfigRoot = *configRoot
		// 개발 환경
		path = "app-dev.json"
	} else if *configRoot == "prod" {
		//ConfigRoot = *configRoot
		// 운영환경 셋팅
		path = "app.json"
	} else {
		panic("Development mode error !!")
	}

	fmt.Println("path : ", path)

	LoadPathConfig(path)
}

