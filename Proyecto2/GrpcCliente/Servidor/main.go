package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	pb "modulo/proto"
	"net"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
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
	topickafka := "votos"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalln(err, "Error al crear el productor")
	}
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topickafka,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("{" + data.Name + "," + data.Album + "," + data.Year + "," + data.Rank + "}"),
	}, nil)
	if err != nil {
		log.Fatalln(err, "Error al producir el mensaje")
	}
	// insertMySQL(data)
	p.Flush(15 * 1000)
	p.Close()
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertMySQL(proyecto Data) {
	query := "INSERT INTO proyecto (name, album, year, ranks) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, proyecto.Name, proyecto.Album, proyecto.Year, proyecto.Rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
	}
}

func main() {
	//mysqlConnect()

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
