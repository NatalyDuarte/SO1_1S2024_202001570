package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client
var redisChannel = "votos_procesados" // Ajustar según el nombre del canal deseado

func RedisConnect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.2:6379", // Replace with "redis-stack:6379" if using the alternative container
		Password: "",                // no password set
		DB:       0,                 // use default DB
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
		"bootstrap.servers": "localhost:9092", // Adjust if necessary
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
	RedisConnect()

	// Create a channel to signal termination
	done := make(chan bool)

	// Handle termination (e.g., user interrupt)
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM) // Add SIGTERM for termination signal

		select {
		case <-sigchan:
			log.Println("Terminating consumer...")
			done <- true
		}
	}()

	// Consume messages
	go func() {
		for {
			log.Println("Probando si entra")
			msg, err := c.ReadMessage(-1)
			log.Printf("Mensaje recibido: %s\n", string(msg.Value))
			if err != nil { // Check for errors in ReadMessage
				// Handle errors from ReadMessage
				if err == io.EOF {
					log.Println("Consumer closed due to termination")
					break
				} else {
					log.Printf("Error al recibir mensaje: %v\n", err)
				}
				continue
			}
			var mensaje = string(msg.Value)
			sinllaves := strings.Trim(mensaje, "{}")
			splitArray := strings.Split(sinllaves, ",")
			fmt.Println("Split array:", splitArray)
			clave := splitArray[0] + splitArray[1] + splitArray[2]
			ran := splitArray[3]
			log.Println(clave)
			log.Println(ran)
			claves := strings.Join([]string{splitArray[0], splitArray[1], splitArray[2]}, ":")
			redisc, rediscan := context.WithTimeout(context.Background(), 1*time.Minute)
			defer rediscan()
			erra := rdb.HIncrBy(redisc, "votosalbum", claves, 1).Err()
			if erra != nil {
				panic(err)
			} else {
				log.Println("Se inserto correctamente")
			}

		}
	}()

	select {
	case <-done:
		log.Println("Consumer terminated")

	}
	c.Close()
}
