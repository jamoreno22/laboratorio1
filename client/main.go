package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"time"

	lab1 "github.com/jamoreno22/laboratorio1/pkg/proto/lab1"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client running s...")

	conn, err := grpc.Dial("10.10.28.17:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := lab1.NewEnvioPaquetesClient(conn)

	// Carga y envío de paquetes
	pymes := readCsvFile("data/pymes.csv")
	retail := readCsvFile("data/retail.csv")

	var paquetes []lab1.Paquete
	for i, v := range pymes {
		paquete := new(lab1.Paquete)
		paquete.Id = v[0]
		paquete.Tipo = v[1]
		paquete.Valor = v[2]
		paquete.Tienda = v[3]
		paquete.Destino = v[4]
		paquete.Prioritario = v[5]
		paquetes = append(paquetes, *paquete)
	}

	for i, v := range retail {
		paquete := new(lab1.Paquete)
		paquete.Id = v[0]
		paquete.Tipo = v[1]
		paquete.Valor = v[2]
		paquete.Tienda = v[3]
		paquete.Destino = v[4]
		paquetes = append(paquetes, *paquete)
	}
	log.Print(paquetes)
	var listaPaquetes = lab1.Paquetes{Paquete: *paquetes}

	//paquete := &lab1.Paquete{Id: "1", Tipo: "pymes", Valor: 30, Tienda: "abc", Destino: "mikasa", Prioritario: false}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetList(ctx, listaPaquetes)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.GetEnvio())
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
