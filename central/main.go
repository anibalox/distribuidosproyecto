package main

import (
	"fmt"
	"io"
	"net"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"

	//rabbit
	"log"

	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

var EquiposDisponibles int
var ColaEspera [4]string
var Termino string

type server struct {
	pb.UnimplementedCentralServiceServer
}

func transformarSituacion(nro int32) string {
	var resultado string
	if nro == 1 {
		resultado = "LISTO"
	} else {
		resultado = "NO LISTO"
	}
	return resultado

}

func (s *server) AbrirComunicacion(stream pb.CentralService_AbrirComunicacionServer) error {

	var situacion *pb.SituacionResp
	var nroLab string

	//Recibir mensaje de introduccion
	situacion, _ = stream.Recv()
	nroLab = situacion.NroLab

	for {
		//Esperar enviar mensaje hasta disponibilidad de equipos y que este Lab sea el primero en la cola
		for ColaEspera[0] != nroLab || EquiposDisponibles == 0 {
		} // CAMBIAR ESTO DEPENDIENDO DE COMO FUNCIONE LA COLA
		EquiposDisponibles -= 1

		//PUEDE HABER PROBLEMAS DE CONCURRENCIA ACA
		//Enviar Ayuda
		stream.Send(&pb.SituacionReq{NroEscuadra: strconv.Itoa(EquiposDisponibles), Termino: Termino})
		if Termino == "1" {
			break
		}

		//Falta hacer un delete del primer objeto aqui. Como un pull de una cola

		//Esperar respuesta de situacion de lab
		for situacion, _ = stream.Recv(); err != io.EOF; situacion, _ = stream.Recv() {
			fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
			time.Sleep(5 * time.Second)
			stream.Send(&pb.SituacionReq{Peticion: 1})
		}
		fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
		fmt.Println("Retorno a Central Escuadra " + situacion.NroEscuadra + ", Conexion Laboratorio " + situacion.NroLab + " Cerrada")
		EquiposDisponibles += 1
	}

	return nil
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever

}

func main() {

	Termino = "0"
	rabbit()

	EquiposDisponibles = 2
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterCentralServiceServer(serv, &server{})

	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

	//Cambiar ciclo para que se repita hasta senal de termino
}

//test
