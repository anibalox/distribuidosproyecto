package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/anibalox/distribuidosproyecto/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLaboratorioServiceServer
}

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func (s *server) PedirSituacion(ctx context.Context, req *pb.SituacionReq) (*pb.SituacionResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}
func (s *server) Finalizar(ctx context.Context, req *pb.FinalizacionReq) (*pb.FinalizacionResp, error) {
	fmt.Println("creating the wish list " + req.WishList.Name)
	return &pb.CreateWishListResp{
		WishListId: req.WishList.Id,
	}, nil
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "my wishlist",
		},
	})

	if err != nil {
		panic("wishlist is not created  " + err.Error())
	}

	fmt.Println(res.WishListId)

}
