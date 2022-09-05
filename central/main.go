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
)

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

	for situacion, _ = stream.Recv(); situacion.Resuelta == 0; situacion, _ = stream.Recv() {
		fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
		time.Sleep(5 * time.Second)
		stream.Send(&pb.SituacionReq{Peticion: 1})
		fmt.Println("queue serv", queue)
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

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		queue = enqueue(queue, string(d.Body))
	}

	//log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) (string, []string) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []string{}
		return element, tmp

	}

	return element, queue[1:] // Slice off the element once it is dequeued.
}

var queue = make([]string, 0)

func main() {
	var forever chan struct{}
	/*
		queue = enqueue(queue, "10")
		fmt.Println("After pushing 10 ", queue)
		queue = enqueue(queue, "20")
		fmt.Println("After pushing 20 ", queue)
		ele, queue := dequeue(queue)
		fmt.Println("Queue After removing", ele, " :", queue)
	*/

	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	nuevo_server := server{}
	go rabbit()
	fmt.Println("After rabbit ", queue)
	pb.RegisterCentralServiceServer(serv, &nuevo_server)
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	<-forever

	//Cambiar ciclo para que se repita hasta senal de termino
}

//test
