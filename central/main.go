package main

import (
	"fmt" // Para imprimir
	"net"
	"strconv"
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
)

var EquiposDisponibles = make([]string, 0)
var ColaEspera = make([]string, 0)
var Termino string
var LabsConectados int
var mu sync.Mutex
var mu2 sync.Mutex
var f *os.File
var ipToNumber = map[string]string{
	"10.6.46.47": "1",
	"10.6.46.48": "2",
	"10.6.46.49": "3",
	"10.6.46.50": "4",
}

type server struct {
	pb.UnimplementedCentralServiceServer
}

func rabbit() {
	conn, err := amqp.Dial("amqp://test:test@localhost:5670/")
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

	for d := range msgs {
		ColaEspera = enqueue(ColaEspera, ipToNumber[string(d.Body)])
		fmt.Println("Mensaje asincrono de laboratorio " + ipToNumber[string(d.Body)] + " leido")
	}

}

func (s *server) Terminar(stream pb.CentralService_TerminarServer) error {
	for Termino == "0" {
		time.Sleep(1 * time.Second)
	}
	time.Sleep(2 * time.Second)
	stream.Send(&pb.Termino{Termino: "1"})
	_, _ = stream.Recv()
	LabsConectados -= 1
	return nil
}

func (s *server) AbrirComunicacion(stream pb.CentralService_AbrirComunicacionServer) error {

	var situacion *pb.SituacionResp
	var nroLab string
	var nroEscuadra string
	var cantidadMensajes int

	LabsConectados += 1

	//Recibir mensaje de introduccion
	situacion, _ = stream.Recv()
	nroLab = ipToNumber[situacion.NroLab]

	for {
		//Esperar enviar mensaje hasta disponibilidad de equipos y que este Lab sea el primero en la cola
		for primeroEnCola(ColaEspera) != nroLab || len(EquiposDisponibles) == 0 {
			time.Sleep(1 * time.Second)
		}

		mu.Lock()
		nroEscuadra, EquiposDisponibles = dequeue(EquiposDisponibles)
		_, ColaEspera = dequeue(ColaEspera)
		mu.Unlock()

		//Enviar Ayuda
		if Termino == "1" {
			for {
				time.Sleep(30 * time.Second)
			}
		}
		fmt.Println("Se envia escuadron " + nroEscuadra + " a laboratorio " + nroLab)
		stream.Send(&pb.SituacionReq{NroEscuadra: nroEscuadra})

		//Realizando battalla. Esperar respuesta de situacion de lab
		cantidadMensajes = 0
		for situacion, _ = stream.Recv(); situacion.Resuelta == "NO LISTO"; situacion, _ = stream.Recv() {

			fmt.Println("Estatus Escuadra " + nroEscuadra + " : [" + situacion.Resuelta + "]")
			cantidadMensajes += 1
			time.Sleep(5 * time.Second)
			if Termino == "1" {
				for {
					time.Sleep(30 * time.Second)
				}
			}
			stream.Send(&pb.SituacionReq{NroEscuadra: nroEscuadra})
		}

		mu2.Lock()
		f.WriteString("Lab" + nroLab + ";" + strconv.Itoa(cantidadMensajes) + "\n")
		mu2.Unlock()

		//Equipo listo. Recibiendo al equipo
		fmt.Println("Estatus Escuadra " + nroEscuadra + " : [" + situacion.Resuelta + "]")
		fmt.Println("Retorno a Central Escuadra " + nroEscuadra + ", Conexion Laboratorio " + nroLab + " Cerrada")
		EquiposDisponibles = enqueue(EquiposDisponibles, nroEscuadra)
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
	ColaEspera = append(ColaEspera, element)
	return ColaEspera
}

func dequeue(ColaEspera []string) (string, []string) {
	element := ColaEspera[0]
	if len(ColaEspera) == 1 {
		var tmp = []string{}
		return element, tmp

	}
	return element, ColaEspera[1:]
}

func main() {
	go rabbit()
	Termino = "0"
	LabsConectados = 0

	EquiposDisponibles = append(EquiposDisponibles, "1")
	EquiposDisponibles = append(EquiposDisponibles, "2")

	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	f, _ = os.Create("SOLICITUDES")

	serv := grpc.NewServer()
	pb.RegisterCentralServiceServer(serv, &server{})

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Termino = "1"
		fmt.Println("Esperando confirmacion de cierre de los laboratorios")
		for LabsConectados != 0 {
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Se cerraron todos los laboratorios. Cerrando ejecucion")
		f.Close()

		os.Exit(1)
	}()

	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
}
