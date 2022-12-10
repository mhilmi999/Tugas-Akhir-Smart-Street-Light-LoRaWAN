package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/config"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/app/database"
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/models"
	repository "github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/home/repository"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var saveMsg models.MqttSubsMsg
	var saveDataMsg models.ConnectionDat
	var dataMqtt string
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	json.Unmarshal(msg.Payload(), &saveMsg)
	// fmt.Println("Ini payload yang disave ke struct: \n", saveMsg)
	Antares_Device_Id := strings.Replace(saveMsg.Pertama.Ketiga.First.Pi, "/antares-cse/cnt-", "", -1)
	dataMqtt = string(saveMsg.Pertama.Ketiga.First.Con)
	json.Unmarshal([]byte(dataMqtt), &saveDataMsg)
	fmt.Println("Ini payload data yang disave ke struct: \n", saveDataMsg)

	repository.NewRepository(db).BindSensorData(saveDataMsg, Antares_Device_Id, 0.0)
	fmt.Println("Sukses (teorinya)")
	// homeparse.Channel <- msg
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func MqttRun() {
	var broker = "mqtt.antares.id"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetKeepAlive(60 * 2 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for true {
		sub(client)
		time.Sleep(time.Second * 5)
	}
	//sub(client)
	//publish(client)

	client.Disconnect(250)
}

// func publish(client mqtt.Client) {
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		text := fmt.Sprintf("Message %d", i)
// 		token := client.Publish("topic/test", 0, false, text)
// 		token.Wait()
// 		time.Sleep(time.Second)
// 	}
// }

func sub(client mqtt.Client) {
	topic := "/oneM2M/resp/antares-cse/01fe7c50a39803d0:93a1cf61893c1605/json"
	// topic := "/oneM2M/req/antares-cse/01fe7c50a39803d0:93a1cf61893c1605/json"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	//fmt.Printf("Subscribed to topic: %s\n", topic)
}
