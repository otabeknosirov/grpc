package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"grpc/generated-proto/order"
	"grpc/model"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	server := Server{}
	grpcServer := grpc.NewServer()
	order.RegisterOrderManagementServer(grpcServer, &server)
	fmt.Println("grpc://localhost:50051 is in process")
	grpcServer.Serve(lis)
}

type Server struct {
}

// AddOrderAndRec bidirectional streaming
func (s *Server) AddOrderAndRec(server order.OrderManagement_AddOrderAndRecServer) error {
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func() {
		defer wait.Done()
		for _, v := range mapDB {
			order := &order.Order{
				Id:          v.Id,
				Items:       v.Items,
				Description: v.Description,
				Price:       v.Price,
			}
			err := server.Send(order)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		defer wait.Done()
		for {
			mapDb, err := server.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
				return
			}
			mapDB[mapDb.OrderType] = model.Order{
				Id:          mapDb.Order.Id,
				Items:       mapDb.Order.Items,
				Description: mapDb.Order.Description,
				Price:       mapDb.Order.Price,
			}
		}
	}()
	wait.Wait()
	return nil
}

func (s *Server) GetAllOrders(ctx context.Context, empty *emptypb.Empty) (*order.Orders, error) {
	orders := make([]*order.Order, 0)
	for _, o := range mapDB {
		var t = &order.Order{
			Id:          o.Id,
			Items:       o.Items,
			Description: o.Description,
			Price:       o.Price,
		}
		orders = append(orders, t)
	}
	return &order.Orders{Orders: orders}, nil
}

// AddOrder client side streaming
func (s *Server) AddOrder(serverStream order.OrderManagement_AddOrderServer) error {
	var countAddedOrders int
	for {
		order, err := serverStream.Recv()
		if err == io.EOF {
			return serverStream.SendAndClose(&wrapperspb.StringValue{Value: "overall added orders:" + strconv.Itoa(countAddedOrders)})
		}
		countAddedOrders++
		mapDB[order.OrderType] = model.Order{
			Id:          order.Order.Id,
			Items:       order.Order.Items,
			Description: order.Order.Description,
			Price:       order.Order.Price,
		}
	}
	return nil
}

// manually-filled map database
var mapDB map[string]model.Order = map[string]model.Order{
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
		Id:          "3",
		Items:       []string{"soup", "shampoo", "cream", "body cream", "cologne", "varnish", "hair dye", "ear sticker"},
		Description: "Perfume",
		Price:       47800000,
	},
}

// GetOrder unary
func (s *Server) GetOrder(ctx context.Context, ProductType *order.ProductType) (*order.Order, error) {
	if mapDB == nil {
		return nil, errors.New("server error")
	}
	or, ok := mapDB[ProductType.ProductType]
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

// GetOrderItems server side streaming (when i am in my local pc)
func (s *Server) GetOrderItems(itemName *wrapperspb.StringValue, server order.OrderManagement_GetOrderItemsServer) error {
	for _, val := range mapDB {
		for i := range val.Items {
			item := val.Items[i]
			if strings.Contains(item, itemName.Value) {
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
