package main

import (
	"context"
	"log"
	"net"

	lab1 "github.com/jamoreno22/laboratorio1/pkg/proto/lab1"
	"google.golang.org/grpc"
)

type server struct {
	lab1.UnimplementedEnvioPaquetesServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	log.Println("Server running ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	lab1.RegisterEnvioPaquetesServer(srv, &server{})
	log.Fatalln(srv.Serve(lis))

	// ***************** rabbitmq **********************
	/*
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"hello", // name
			false,   // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")

		body := "Hello World!"
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	*/
	//--------------------------------------------------
}

func (s *server) GetList(ctx context.Context, req *lab1.Paquetes) (*lab1.Respuesta, error) {
	return &lab1.Respuesta{Envio: "conected client"}, nil
}
func (s *server) Seguimiento(ctx context.Context, req *lab1.Paquete) (*lab1.Estado, error) {
	return &lab1.Estado{Estado: "conected"}, nil
}
func (s *server) HistorialCamion(ctx context.Context, req *lab1.EstadoCamion) (*lab1.Respuesta, error) {
	return &lab1.Respuesta{Envio: "conected"}, nil
}
