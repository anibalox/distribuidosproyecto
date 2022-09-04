package main

import (
	"fmt"
	"net"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"

	//rabbit
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"strconv"
)

type server struct {
	pb.UnimplementedCentralServiceServer
	EquiposDisponibles int
	
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

<<<<<<< HEAD
func (s *server) EsperarAyuda(context.Context, *AyudaReq) (*AyudaResp, error) {
	for s.EquiposDisponibles == 0{}
	nro_equipo:= strconv.Itoa(s.EquiposDisponibles)
	s.EquiposDisponibles -= 1
	return &pb.AyudaResp{Escuadron: nro_equipo}, nil

func (s *server) AbrirComunicacion(stream pb.CentralService_AbrirComunicacionServer) error {

	var situacion *pb.SituacionResp

=======
func (s *server) AbrirComunicacion(stream pb.CentralService_AbrirComunicacionServer) error {

	var situacion *pb.SituacionResp

>>>>>>> 4d5c0447d0e777317aa805087e50399aaad7e6d8
	for situacion, _ = stream.Recv(); situacion.Resuelta == 0; situacion, _ = stream.Recv() {
		fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
		time.Sleep(5 * time.Second)
		stream.Send(&pb.SituacionReq{Peticion: 1})
	}
	fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
	fmt.Println("Retorno a Central Escuadra " + situacion.NroEscuadra + ", Conexion Laboratorio " + situacion.NroLab + " Cerrada")

	s.EquiposDisponibles += 1
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

type mensajes struct {
	ip     string
	puerto string
}

func main() {
<<<<<<< HEAD
	rabbit()
=======
>>>>>>> 4d5c0447d0e777317aa805087e50399aaad7e6d8

	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
<<<<<<< HEAD
	pb.RegisterCentralServiceServer(serv, &server{EquiposDisponibles: 2})
=======
	pb.RegisterCentralServiceServer(serv, &server{})
>>>>>>> 4d5c0447d0e777317aa805087e50399aaad7e6d8
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

	//Cambiar ciclo para que se repita hasta senal de termino
}

//test
