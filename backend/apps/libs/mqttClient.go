package libs

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	UUID "github.com/satori/go.uuid"
	"fmt"
	"strings"
	"strconv"
)

// 온실 내부 Local Client 만들기
func NewMqttLocalClient(isWillTopic bool) MQTT.Client {

	mqttUrl := Config.MQTT.LocalUrl
	mqttUrls := strings.Split(mqttUrl, ",")

	fmt.Println(">>>>>>>>>>>>>>>>>>>> mqttUrls : ", mqttUrls)

	opts := MQTT.NewClientOptions()
	for _, uri := range mqttUrls {
		opts.AddBroker(uri)
	}
	opts.SetClientID(UUID.NewV4().String())

	if isWillTopic {
		nodeId := strconv.FormatInt(Config.INFO.NodeId, 10)

		opts.SetWill("/kgw/v2/L/EVT/DATA/LWT/" + nodeId, "", 0, false)
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}