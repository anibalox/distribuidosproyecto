package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"

	//rabbit
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CalcularEstallido() string {
	var resultado string
	if rand.Float64() <= 0.8 {
		resultado = "ESTALLIDO"
	} else {
		resultado = "OK"
	}

	return resultado
}

func CalcularResolucion() string {
	var resultado string
	if rand.Float64() <= 0.6 {
		resultado = "LISTO"
	} else {
		resultado = "NO LISTO"
	}

	return resultado

}

func rabbit() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") //Escribir datos de la central
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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
func ComunicarseConCentral(client pb.CentralServiceClient, nro_lab string) {

	var resolucion string
	var nro_escuadron string
	var situacion *pb.SituacionReq
	var estallido string

	stream, _ := client.AbrirComunicacion(context.Background()) //stream, err := client.AbrirComunicacion(context.Background())

	//Mensaje de introduccion
	stream.Send(&pb.SituacionResp{Resuelta: "NO LISTO", NroLab: nro_lab})

	for {
		//Calculo de Estallido
		time.Sleep(5 * time.Second)
		for estallido = CalcularEstallido(); estallido == "OK"; estallido = CalcularEstallido() {
			fmt.Println("Analizando estado Laboratorio [" + estallido + "]")
			time.Sleep(5 * time.Second)
		}
		fmt.Println("Analizando estado Laboratorio [" + estallido + "]")
		fmt.Println("SOS Enviado a Central. Esperando respuesta...")

		//Envia mensaje
		rabbit()

		//Esperar recibir ayuda
		situacion, _ = stream.Recv()
		nro_escuadron = situacion.NroEscuadra
		fmt.Println("Llega escuadron " + nro_escuadron + ", conteniendo estallido")

		// Comienza batalla
		for resolucion = CalcularResolucion(); resolucion == "NO LISTO"; resolucion = CalcularResolucion() {
			fmt.Println("Revisando Estado Escuadron: [" + resolucion + "]")
			stream.Send(&pb.SituacionResp{Resuelta: resolucion}) // Puede que de problemas con el campo no incluido
			_, _ = stream.Recv()
		}
		//Termina batalla. Devolviendo equipo
		fmt.Println("Revisando Estado Escuadron: [" + resolucion + "]")
		stream.Send(&pb.SituacionResp{Resuelta: resolucion}) // Mismo problema de campo no incluido
		stream.CloseSend()
		fmt.Println("Estallido contenido. Escuadron " + nro_escuadron + " Retornando")
	}

}
func DarNumeroLab(ip string, puerto string) string {
	return "2"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano()) // iniciar semilla

	ip_Central := "localhost" //Colocar valores para esto
	port_Central := "50051"

	nro_lab := DarNumeroLab(ip_lab, port_lab) //Falta definir como le damos los nombres a los labs

	//Enviar mensaje con Rabbit. Esperar respuesta...

	conn, err := grpc.Dial(ip_Central+":"+port_Central, grpc.WithInsecure()) //grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCentralServiceClient(conn)

	go ComunicarseConCentral(serviceClient, nro_lab)

	stream, _ := serviceClient.AbrirComunicacion(context.Background())
	_, _ = stream.Recv()                   // Recibir senal de termino
	stream.Send(&pb.Termino{Termino: "1"}) // Enviar Confirmacion
	os.Exit(1)
}
