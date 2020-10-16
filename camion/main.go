package main

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	lab1 "github.com/jamoreno22/laboratorio1/pkg/proto/lab1"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running ...")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := lab1.NewEnvioPaquetesClient(conn)

	estadoCamion := &lab1.EstadoCamion{Id: "1", Tipo: "pymes", Valor: 30, Origen: "abc", Destino: "mikasa", Intentos: 1, Fecha: ptypes.TimestampNow()}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.HistorialCamion(ctx, estadoCamion)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.GetEnvio())
}
