package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"grpc/generated-proto/order"
	"grpc/model"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	clientOrder := order.NewOrderManagementClient(conn)
	fmt.Println("1st:")
	order, err := GetOrder(clientOrder, "food")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result1:", order)
	fmt.Println("2nd:")
	GetOrders(clientOrder, "hair dye")
	fmt.Println("3rd:")
	result, err := AddOrder(clientOrder)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("result3:%v\n", result)

	fmt.Println("4th:")
	AddAndReceive(clientOrder)
}

func GetOrder(client order.OrderManagementClient, productType string) (*order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	order, err := client.GetOrder(ctx, &order.ProductType{ProductType: productType})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order, nil
}

// GetOrders server side streaming
func GetOrders(clientOrder order.OrderManagementClient, item string) {
	orderStream, err := clientOrder.GetOrderItems(context.Background(), &wrapperspb.StringValue{Value: item})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("result2:")
	for {
		searchOrder, err := orderStream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(searchOrder)
	}
}

var mapDB map[string]model.Order = map[string]model.Order{
	"national-food": {
		Id:          "4",
		Items:       []string{"soup", "shurva", "osh", "shashlik"},
		Description: "Food",
		Price:       33400000,
	},
	"stationary": {
		Id:          "5",
		Items:       []string{"pen", "pencil", "notebook"},
		Description: "Staff",
		Price:       41204000,
	},
}

func AddOrder(clientOrder order.OrderManagementClient) ([]model.Order, error) {
	clientStream, err := clientOrder.AddOrder(context.Background())
	if err != nil {
		return nil, err
	}
	for k, o := range mapDB {
		var t = &order.Order{
			Id:          o.Id,
			Items:       o.Items,
			Description: o.Description,
			Price:       o.Price,
		}
		err = clientStream.Send(&order.MapDb{
			OrderType: k,
			Order:     t,
		})
		if err != nil {
			return nil, err
		}
	}
	countAddedOrders, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Println("cannot get orders")
		return nil, err
	}
	fmt.Println(countAddedOrders.Value)
	orders, err := clientOrder.GetAllOrders(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println("no data found")
		return nil, err
	}
	result := make([]model.Order, len(orders.Orders))
	for i := range orders.Orders {
		result[i].Id = orders.Orders[i].Id
		result[i].Items = orders.Orders[i].Items
		result[i].Description = orders.Orders[i].Description
		result[i].Price = orders.Orders[i].Price
	}
	return result, nil
}

func AddAndReceive(client order.OrderManagementClient) {
	var mapDb map[string]model.Order = map[string]model.Order{
		"toys": {
			Id:          "6",
			Items:       []string{"barbi", "pistols"},
			Description: "Baby",
			Price:       63400000,
		},
		"clothes": {
			Id:          "7",
			Items:       []string{"jeans", "t-shirt", "socks", "smoking"},
			Description: "Clothes",
			Price:       7120440520,
		},
	}
	done := make(chan struct{})
	stream, err := client.AddOrderAndRec(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		for {
			order, err := stream.Recv()
			if err == io.EOF {
				close(done)
				break
			}
			fmt.Println(order)
		}
	}()
	for k, v := range mapDb {
		err = stream.Send(&order.MapDb{
			OrderType: k,
			Order: &order.Order{
				Id:          v.Id,
				Items:       v.Items,
				Description: v.Description,
				Price:       v.Price,
			},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	stream.CloseSend()
	<-done
}
