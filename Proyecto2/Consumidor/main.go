package main

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func RedisConnect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pong)

}

func main() {
	// set up kafka consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", ///my-cluster-kafka-bootstrap:9092
		"group.id":          "console-consumer-21564",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalln(err, "Error al crear el productor")
	}
	// suscribe to the topic
	err = c.SubscribeTopics([]string{"votos"}, nil)
	if err != nil {
		log.Fatalln(err, "Error al suscribirse al tópico")
	}
	fmt.Println("Consumidor de mensajes")
	//RedisConnect()
	// consume messages
	go func() {
		for {
			msg, err := c.ReadMessage(-1)
			if err == nil {
				log.Printf("Mensaje recibido: %s\n", string(msg.Value))
			} else {
				log.Printf("Error al recibir mensaje: %v\n", err)
			}
		}
	}()

	select {}
}
