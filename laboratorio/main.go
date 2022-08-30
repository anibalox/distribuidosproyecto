package main

import (
	"context"
	"fmt"
	"math/rand"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"
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

func transformarSituacion(str string) int {
	var resultado int
	if str == "LISTO" {
		resultado = 1
	} else {
		resultado = 0
	}
	return resultado

}

func ComunicarseConCentral(client pb.centralServiceClient, nro_escuadron string, nro_lab string) {

	var resolucion string
	stream, err := client.AbrirComunicacion(context.Background())

	for resolucion = CalcularResolucion(); resolucion == "NO LISTO"; resolucion = CalcularResolucion() {
		fmt.Println("Revisando Estado Escuadron: [" + resolucion + "]")
		stream.Send(&pb.SituacionResp{resuelta: transformarSituacion(resolucion), nro_escuadra: nro_escuadron, nro_lab: nro_lab})
		_, _ = stream.Recv()
	}
	stream.CloseSend(&pb.SituacionResp{resuelta: transformarSituacion(resolucion), nro_escuadra: nro_escuadron, nro_lab: nro_lab})
	fmt.Println("Estallido contenido. Escuadron " + nro_escuadron + " Retornando")
}

func main() {

	ip_Central := "1111" //Colocar valores para esto
	port_Central := "2222"

	//Enviar mensaje con Rabbit. Esperar respuesta...

	conn, err := grpc.Dial(ip_Central+":"+port_Central, grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	//FALTA COLOCAR LOOP IMPORTANTE!!!!!!!!!!!!!!!!!!!!!!!

	for estallido := CalcularEstallido(); estallido == "OK"; estallido = CalcularEstallido() {
		fmt.Println("Analizando estado Laboratorio [" + estallido + "]")
	}
	fmt.Println("SOS Enviado a Central. Esperando respuesta...")
	//Enviar mensaje con Rabbit. Esperar respuesta...

	//Hay que obtener el nro de lab y de escuadron
	nro_escuadron := "1"
	nro_lab := "2"

	//Leer mensaje con Rabbit. No se como funcione
	fmt.Println("Llega escuadron " + nro_escuadron + ", conteniendo estallido")

	//Comienzo del envio de mensajes
	ComunicarseConCentral(serviceClient, nro_escuadron, nro_lab)

	//HASTA ACA EL LOOP!!!!!!!!!!!!!!!!!!!!!!!!!!!

}
