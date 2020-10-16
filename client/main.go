package main

import (
	"context"
	"log"
	"time"

	lab1 "github.com/jamoreno22/laboratorio1/pkg/proto/lab1"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running ...")

	conn, err := grpc.Dial("10.10.28.17:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := lab1.NewEnvioPaquetesClient(conn)

	paquete := &lab1.Paquete{Id: "1", Tipo: "pymes", Valor: 30, Tienda: "abc", Destino: "mikasa", Prioritario: false}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetList(ctx, paquete)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.GetEnvio())
}
