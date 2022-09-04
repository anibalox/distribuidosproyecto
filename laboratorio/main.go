package main

import (
	"context"
	"fmt"
	"math/rand"
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

func transformarSituacion(str string) int32 {
	var resultado int32
	if str == "LISTO" {
		resultado = 1
	} else {
		resultado = 0
	}
	return resultado

}

func ComunicarseConCentral(client pb.CentralServiceClient, nro_escuadron string, nro_lab string) {

	var resolucion string
	stream, _ := client.AbrirComunicacion(context.Background()) //stream, err := client.AbrirComunicacion(context.Background())

	for resolucion = CalcularResolucion(); resolucion == "NO LISTO"; resolucion = CalcularResolucion() {
		fmt.Println("Revisando Estado Escuadron: [" + resolucion + "]")
		stream.Send(&pb.SituacionResp{Resuelta: transformarSituacion(resolucion), NroEscuadra: nro_escuadron, NroLab: nro_lab})
		_, _ = stream.Recv()
	}
<<<<<<< HEAD
	fmt.Println("Revisando Estado Escuadron: [" + resolucion + "]")
=======
>>>>>>> 4d5c0447d0e777317aa805087e50399aaad7e6d8
	stream.Send(&pb.SituacionResp{Resuelta: transformarSituacion(resolucion), NroEscuadra: nro_escuadron, NroLab: nro_lab})
	stream.CloseSend()
	fmt.Println("Estallido contenido. Escuadron " + nro_escuadron + " Retornando")
}

func EsperarAyuda(client pb.centralServiceClient) string {
	var respuesta *pb.AyudaResp
	respuesta, _ = client.EsperarAyuda(context.Background(), &pb.AyudaReq{Ayuda: "1"})
	return respuesta.Escuadron
}

func DarNumeroLab() string {
	return "2"
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func rabbit() {
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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

func main() {
	rabbit()

	ip_Central := "localhost" //Colocar valores para esto
	port_Central := "50051"

	nro_lab := "2"

	var estallido string
	//Enviar mensaje con Rabbit. Esperar respuesta...

	conn, err := grpc.Dial(ip_Central+":"+port_Central, grpc.WithInsecure()) //grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewCentralServiceClient(conn)

	//FALTA COLOCAR LOOP IMPORTANTE!!!!!!!!!!!!!!!!!!!!!!!

	time.Sleep(5 * time.Second)
	for estallido = CalcularEstallido(); estallido == "OK"; estallido = CalcularEstallido() {
		fmt.Println("Analizando estado Laboratorio [" + estallido + "]")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Analizando estado Laboratorio [" + estallido + "]")
	fmt.Println("SOS Enviado a Central. Esperando respuesta...")
	//Enviar mensaje con Rabbit. Esperar respuesta...

	nro_escuadron := EsperarAyuda(serviceClient)

	//Leer mensaje con Rabbit. No se como funcione
	fmt.Println("Llega escuadron " + nro_escuadron + ", conteniendo estallido")

	//Comienzo del envio de mensajes
	ComunicarseConCentral(serviceClient, nro_escuadron, nro_lab)

	//HASTA ACA EL LOOP!!!!!!!!!!!!!!!!!!!!!!!!!!!

	//Para cerrar conn.Close()

}
