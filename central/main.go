package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/anibalox/distribuidosproyecto/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCentralServiceServer
}

func (s *server) Ayuda(ctx context.Context, req *pb.AyudaReq) (*pb.AyudaResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterWishListServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}
