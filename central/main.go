package main

import (
	l "container/list"
	"fmt"
	"net"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"

	//rabbit
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type server struct {
	pb.UnimplementedCentralServiceServer
	mensaje string
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

	for situacion, _ = stream.Recv(); situacion.Resuelta == 0; situacion, _ = stream.Recv() {
		println("test 1")
		fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
		time.Sleep(5 * time.Second)
		stream.Send(&pb.SituacionReq{Peticion: 1})
		println("test 2")
	}
	fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
	fmt.Println("Retorno a Central Escuadra " + situacion.NroEscuadra + ", Conexion Laboratorio " + situacion.NroLab + " Cerrada")

	return nil
}

func EnviarAyuda(ip string, puerto string) {

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func rabbit(serv server) {
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

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		serv.mensaje = string(d.Body)
	}

	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

}

var queue *l.List

func main() {
	var forever chan struct{}
	queue = l.New()
	queue.PushBack("hola")
	queue.PushBack(2)
	fmt.Println(queue.Front().Value)

	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	nuevo_server := server{mensaje: "xd"}
	print(nuevo_server.mensaje + "aqui /n")
	go rabbit(nuevo_server)
	pb.RegisterCentralServiceServer(serv, &nuevo_server)
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	print("test /n")
	print(nuevo_server.mensaje + "otro /n")
	<-forever

	//Cambiar ciclo para que se repita hasta senal de termino
}

//test
