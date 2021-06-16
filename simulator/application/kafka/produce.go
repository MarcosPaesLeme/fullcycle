package kafka

import (
	"encoding/json"
	route2 "github.com/MarcosPaesLeme/imersaofsfc2-simulator/application/route"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/MarcosPaesLeme/imersaofsfc2-simulator/infra/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _,p := range positions {
		kafka.Publisher(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
} 