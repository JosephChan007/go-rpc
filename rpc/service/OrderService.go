package service

import (
	"context"
	"errors"
	"fmt"
	. "github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io"
	"log"
	"time"
)

var orderMap = make(map[string]*OrderInfo, 5)

func OrderHouse() *map[string]*OrderInfo {
	if len(orderMap) == 0 {
		orderMap["20200501001"] = &OrderInfo{OrderId: "20200501001", OrderName: "橙子", Status: OrderStatus_UnPay, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501002"] = &OrderInfo{OrderId: "20200501002", OrderName: "苹果", Status: OrderStatus_Paied, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501003"] = &OrderInfo{OrderId: "20200501003", OrderName: "梨", Status: OrderStatus_Commited, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501004"] = &OrderInfo{OrderId: "20200501004", OrderName: "香蕉", Status: OrderStatus_Refund, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501005"] = &OrderInfo{OrderId: "20200501005", OrderName: "火龙果", Status: OrderStatus_Refunded, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
	}
	return &orderMap
}

// For rpc
/*
type OrderService struct {
}

func (s *OrderService) GetOrder(req *message.OrderRequest, res *message.OrderInfo) error {
	log.Printf("order service requst is: %v\n", req)

	orderMap := OrderHouse()
	for _, v := range *orderMap {
		if v.OrderId == req.OrderId {
			*res = v
			fmt.Printf("order service info is: %v\n", res)
			break
		}
	}

	if res.OrderId == "" {
		return errors.New("no data")
	}

	return nil
}
*/

// For gprc-go

type OrderServiceImpl struct {
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *OrderRequest) (res *OrderInfo, err error) {
	log.Printf("order service requst is: %v\n", req)

	orderMap := OrderHouse()
	for _, v := range *orderMap {
		if v.OrderId == req.OrderId {
			res = v
			fmt.Printf("order service info is: %v\n", res)
			break
		}
	}

	if res.OrderId == "" {
		err = errors.New("no data")
		return nil, err
	}

	return res, nil
}

func (s *OrderServiceImpl) GetOrderList(ctx context.Context, req *OrderSize) (res *OrderInfoList, err error) {
	log.Printf("order list service requst is: %v\n", req)

	orderMap := OrderHouse()
	list := make([]*OrderInfo, 0)
	for _, v := range *orderMap {
		list = append(list, v)
	}

	return &OrderInfoList{
		Orders: list,
	}, nil
}

func (s *OrderServiceImpl) NewOrder(ctx context.Context, req *OrderData) (res *OrderResponse, err error) {
	log.Printf("create a new order, data is: %v\n", req)

	orderMap := OrderHouse()
	if _, exist := (*orderMap)[req.OrderId]; exist {
		return &OrderResponse{
			OrderId: "-1",
			Message: "failed",
		}, nil
	}

	(*orderMap)[req.OrderId] = &OrderInfo{OrderId: req.OrderId, OrderName: req.OrderName, Status: req.Status, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
	log.Printf("create a new order success, data is: %v\n", (*orderMap)[req.OrderId])
	return &OrderResponse{
		OrderId: (*orderMap)[req.OrderId].OrderId,
		Message: "ok",
	}, nil
}

func (s *OrderServiceImpl) GetOrderListByServerStream(req *OrderStatusRequest, stream OrderService_GetOrderListByServerStreamServer) error {
	log.Printf("query status order, data is: %v\n", req)

	status := req.Status
	orderMap := OrderHouse()
	orders := make([]*OrderInfo, 0)
	for _, v := range *orderMap {
		for _, state := range status {
			if state != v.Status {
				continue
			}
			orders = append(orders, v)
			if len(orders)%2 == 0 && len(orders) > 0 {
				err := stream.Send(&OrderInfoList{
					Orders: orders,
				})
				if err != nil {
					return err
				}
				log.Printf("query status order, data is: %v\n", orders)
				time.Sleep(time.Second)
				orders = (orders)[0:0]
			}
		}
	}

	if len(orders) > 0 {
		err := stream.Send(&OrderInfoList{
			Orders: orders,
		})
		if err != nil {
			return err
		}
		log.Printf("query status order, data is: %v\n", orders)
		orders = (orders)[0:0]
	}

	return nil
}

func (s *OrderServiceImpl) GetOrderListByClientStream(stream OrderService_GetOrderListByClientStreamServer) error {

	orderMap := OrderHouse()
	orders := make([]*OrderInfo, 0)
	for {
		request, err := stream.Recv()
		log.Printf("query status order, data is: %v\n", request)
		if err == io.EOF {
			log.Printf("query status order, data is: %v\n", orders)
			return stream.SendAndClose(&OrderInfoList{
				Orders: orders,
			})
		}
		if err != nil {
			return err
		}
		for _, v := range *orderMap {
			for _, state := range request.Status {
				if v.Status == state {
					orders = append(orders, v)
				}
			}
		}
		time.Sleep(time.Second * 1)
	}

	return nil
}

func (s *OrderServiceImpl) GetOrderListByDoubleStream(stream OrderService_GetOrderListByDoubleStreamServer) error {

	orderMap := OrderHouse()
	orders := make([]*OrderInfo, 0)
	for {
		request, err := stream.Recv()
		log.Printf("query status order, data is: %v\n", request)
		if err == io.EOF {
			log.Printf("query status order, data is: EOF")
			break
		}
		if err != nil {
			return err
		}
		for _, v := range *orderMap {
			for _, state := range request.Status {
				if v.Status == state {
					orders = append(orders, v)
				}
			}
		}

		err = stream.Send(&OrderInfoList{
			Orders: orders,
		})
		if err != nil {
			log.Printf("Server stream send error: %v", err)
		}
		orders = (orders)[0:0]

		time.Sleep(time.Second * 1)
	}

	return nil
}
