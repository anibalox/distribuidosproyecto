package main

import (
	"fmt" // Para imprimir
	"net"
	"sync"

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
var ColaEspera = make([]string, 0) //var ColaEspera [4]string
var Termino string
var LabsCerrados int
var mu sync.Mutex

type server struct {
	pb.UnimplementedCentralServiceServer
}

func rabbit() {
	conn, err := amqp.Dial("amqp://test:test@localhost:5670/") //conn, err := amqp.Dial("amqp://guest:guest@" + myIP() + ":5672/")
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

	//go func() {
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		println("iteracion rabbit")
		ColaEspera = enqueue(ColaEspera, string(d.Body))
	}
	//}()
}

func (s *server) Terminar(stream pb.CentralService_TerminarServer) error {
	for Termino == "0" {
		time.Sleep(1 * time.Second)
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
		for primeroEnCola(ColaEspera) != nroLab || EquiposDisponibles == 0 {
			time.Sleep(1 * time.Second)
		}

		nroEscuadra = strconv.Itoa(EquiposDisponibles)
		EquiposDisponibles -= 1
		mu.Lock()
		_, ColaEspera = dequeue(ColaEspera)
		mu.Unlock()

		//Enviar Ayuda
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
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func primeroEnCola(ColaEspera []string) string {
	if len(ColaEspera) == 0 {
		return "0"
	}
	return ColaEspera[0]
}

func enqueue(ColaEspera []string, element string) []string {
	ColaEspera = append(ColaEspera, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return ColaEspera
}

func dequeue(ColaEspera []string) (string, []string) {
	element := ColaEspera[0] // The first element is the one to be dequeued.
	if len(ColaEspera) == 1 {
		var tmp = []string{}
		return element, tmp

	}
	return element, ColaEspera[1:] // Slice off the element once it is dequeued.
}

func main() {
	//println(len(ColaEspera))
	//var forever chan struct{}
	go rabbit()
	Termino = "0"
	LabsCerrados = 0

	EquiposDisponibles = 2

	//ColaEspera = append(ColaEspera, "0") //ColaEspera[0] = "2"
	//fmt.Println(ColaEspera[0])
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
		Termino = "1"
		/*
			for LabsCerrados != 4 {
				time.Sleep(1 * time.Second)
			}
		*/
		os.Exit(1)
	}()

	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	//<-forever
}
