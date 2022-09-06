package main

import (
	"fmt" // Para imprimir
	"net"

	//Para terminar los programas con un senal
	"os"
	"os/signal"
	"syscall"

	"time" // Para utilizar otras seed

	//Donde estan los datos del .proto
	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"

	//rabbit
	"log"
	//Tambien rabbit
	amqp "github.com/rabbitmq/amqp091-go"

	//Para convertir los EquiposDisponibles en string
	"strconv"
)

var EquiposDisponibles int
var ColaEspera [4]string
var Termino string
var LabsCerrados int

type server struct {
	pb.UnimplementedCentralServiceServer
}

func (s *server) Terminar(stream pb.CentralService_TerminarServer) error {
	for Termino == "0" {
		println("termino")
	}
	stream.Send(&pb.Termino{Termino: "1"})
	_, _ = stream.Recv()
	LabsCerrados += 1
	return nil
}

func (s *server) AbrirComunicacion(stream pb.CentralService_AbrirComunicacionServer) error {

	var situacion *pb.SituacionResp
	var nroLab string
	var nroEscuadra string

	//Recibir mensaje de introduccion
	situacion, _ = stream.Recv()
	nroLab = situacion.NroLab

	for {
		//Esperar enviar mensaje hasta disponibilidad de equipos y que este Lab sea el primero en la cola
		for ColaEspera[0] != nroLab || EquiposDisponibles == 0 {
		} // CAMBIAR ESTO DEPENDIENDO DE COMO FUNCIONE LA COLA
		EquiposDisponibles -= 1
		//Eliminar el dato de cabeza de la cola

		//Enviar Ayuda
		nroEscuadra = strconv.Itoa(EquiposDisponibles)
		stream.Send(&pb.SituacionReq{NroEscuadra: nroEscuadra})

		//Realizando battalla. Esperar respuesta de situacion de lab
		for situacion, _ = stream.Recv(); situacion.Resuelta == "NO LISTO"; situacion, _ = stream.Recv() {
			fmt.Println("Estatus Escuadra " + nroEscuadra + " : [" + situacion.Resuelta + "]")
			time.Sleep(5 * time.Second)
			stream.Send(&pb.SituacionReq{NroEscuadra: nroEscuadra})
		}
		//Equipo listo. Recibiendo al equipo
		fmt.Println("Estatus Escuadra " + nroEscuadra + " : [" + situacion.Resuelta + "]")
		fmt.Println("Retorno a Central Escuadra " + nroEscuadra + ", Conexion Laboratorio " + nroLab + " Cerrada")
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
	LabsCerrados = 0
	rabbit()
	EquiposDisponibles = 2
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterCentralServiceServer(serv, &server{})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		for LabsCerrados != 4 {
		}
		os.Exit(1)
	}()

	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
