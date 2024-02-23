package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"grpc/generated-proto/order"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8083))
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	server := Server{}
	grpcServer := grpc.NewServer()
	order.RegisterOrderManagementServer(grpcServer, &server)
	fmt.Println("grpc://localhost:8083 is in process")
	grpcServer.Serve(lis)
}

type Order struct {
	Id          string
	Items       []string
	Description string
	Price       uint64
}

type Server struct {
}

// manually-filled map database
var mapDB map[string]Order = map[string]Order{
	"perfume-rose": {
		Id:          "1",
		Items:       []string{"soup", "shampoo", "cream", "body cream", "cologne", "varnish", "hair dye"},
		Description: "Perfume",
		Price:       45000000,
	},
	"food": {
		Id:          "2",
		Items:       []string{"hot-dog", "hamburger", "cheeseburger", "chicken", "lavish"},
		Description: "Fast Food",
		Price:       2300000,
	},
	"perfume-kudrat": {
		Id:          "1",
		Items:       []string{"soup", "shampoo", "cream", "body cream", "cologne", "varnish", "hair dye", "ear sticker"},
		Description: "Perfume",
		Price:       47800000,
	},
}

// GetOrder unary
func (s *Server) GetOrder(ctx context.Context, ProductId *order.ProductId) (*order.Order, error) {
	if mapDB == nil {
		return nil, errors.New("server error")
	}
	or, ok := mapDB[ProductId.Id]
	if !ok {
		return nil, errors.New("no such order found")
	}
	return &order.Order{
		Id:          or.Id,
		Items:       or.Items,
		Description: or.Description,
		Price:       or.Price,
	}, nil
}

func (s *Server) GetOrderItems(itemType *order.Item, server order.OrderManagement_GetOrderItemsServer) error {
	for _, val := range mapDB {
		for i := range val.Items {
			item := val.Items[i]
			if strings.Contains(item, itemType.GetItem()) {
				if err := server.Send(&order.Order{
					Id:          val.Id,
					Items:       val.Items,
					Description: val.Description,
					Price:       val.Price,
				}); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func (s *Server) mustEmbedUnimplementedOrderManagementServer() {
	//TODO implement me
	panic("implement me")
}
