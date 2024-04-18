package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	pb "modulo/proto" // Replace with your proto package name
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var db *sql.DB

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

// Kafka configurations
var (
	kafkaBrokers  = "localhost:9092"      // Replace with your Kafka broker address
	kafkaTopic    = "your-topic-name"     // Replace with your topic name
	consumerGroup = "your-consumer-group" // Replace with your consumer group ID
)

func mysqlConnect() {
	dsn := "root:tarea@tcp(34.85.187.123:3306)/tarea4"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(data)
	//insertMySQL(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertMySQL(proyecto Data) {
	query := "INSERT INTO proyecto (name, album, year, ranks) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, proyecto.Name, proyecto.Album, proyecto.Year, proyecto.Rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
	}
}

func handleKafkaMessages() {
	// Kafka consumer configuration
	config := kafka.ReaderConfig{
		Brokers:  []string{kafkaBrokers},
		GroupID:  consumerGroup,
		Topic:    kafkaTopic,
		MinBytes: 1024,
		MaxBytes: 10e6,
	}

	// Create a Kafka reader
	reader := kafka.NewReader(config)
	fmt.Println("Listening to Kafka topic:", kafkaTopic)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading Kafka message:", err)
			continue
		}

		// Unmarshal the message value (assuming JSON format)
		var data Data
		err = json.Unmarshal(msg.Value, &data) // Modify if needed
		if err != nil {
			log.Println("Error unmarshalling Kafka message:", err)
			continue
		}

		// Use the data object for your application logic
		insertMySQL(data)
	}
}

func main() {
	// Connect to MySQL
	mysqlConnect()

	// Start Kafka consumer in a separate goroutine
	go handleKafkaMessages()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
