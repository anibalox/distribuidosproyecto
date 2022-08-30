package main

import (
	"fmt"
	"net"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"
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

	time.Sleep(5 * time.Second)
	for situacion, _ = stream.Recv(); situacion.Resuelta == 0; situacion, _ = stream.Recv() {
		stream.Send(&pb.SituacionReq{Peticion: 1})
		fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Estatus Escuadra " + situacion.NroEscuadra + " : [" + transformarSituacion(situacion.Resuelta) + "]")
	fmt.Println("Retorno a Central Escuadra " + situacion.NroEscuadra + ", Conexion Laboratorio " + situacion.NroLab + " Cerrada")

	return nil
}

func EnviarAyuda(ip string, puerto string) {

}

type mensajes struct {
	ip     string
	puerto string
}

func main() {

	var lista_mensajes []mensajes // Definir mensajes con Rabbit
	cantidad_equipos := 2
	var puerto, ip string

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

	for mensaje := range lista_mensajes {
		puerto = mensaje.puerto
		ip = mensaje.ip
	}

}

//test
