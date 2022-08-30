package main

import (
	"fmt"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
)

type server struct {
	pb.UnimplementedCentralServiceServer
}

func transformarSituacion(nro int) string {
	var resultado string
	if nro == 1 {
		resultado = "LISTO"
	} else {
		resultado = "NO LISTO"
	}
	return resultado

}

func (s *server) AbrirComunicacion(stream *pb.CentralService_AbrirComunicacionServer) error {

	var situacion pb.SituacionResp

	for situacion, _ = stream.Recv(); situacion.resuelta == 0; situacion = stream.Recv() {
		time.Sleep(5 * time.Second)
		stream.Send(1)
		fmt.Println("Estatus Escuadra " + situacion.nro_escuadra + " : [" + transformarSituacion(situacion.resuelta) + "]")
	}
	fmt.Println("Retorno a Central Escuadra " + situacion.nro_escuadra + ", Conexion Laboratorio " + situacion.nro_lab + " Cerrada")

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

	//Cambiar ciclo para que se repita hasta senal de termino

	for mensaje := range lista_mensajes {
		puerto = mensaje.puerto
		ip = mensaje.ip
	}

}
