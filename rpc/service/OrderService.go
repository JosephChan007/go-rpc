package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/JosephChan007/go-rpc/rpc/message"
	"github.com/golang/protobuf/ptypes/timestamp"
	"log"
	"time"
)

var orderMap = make(map[string]*message.OrderInfo, 5)

func OrderHouse() *map[string]*message.OrderInfo {
	if len(orderMap) == 0 {
		orderMap["20200501001"] = &message.OrderInfo{OrderId: "20200501001", OrderName: "橙子", Status: message.OrderStatus_UnPay, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501002"] = &message.OrderInfo{OrderId: "20200501002", OrderName: "苹果", Status: message.OrderStatus_Paied, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501003"] = &message.OrderInfo{OrderId: "20200501003", OrderName: "梨", Status: message.OrderStatus_Commited, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501004"] = &message.OrderInfo{OrderId: "20200501004", OrderName: "香蕉", Status: message.OrderStatus_Refund, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
		orderMap["20200501005"] = &message.OrderInfo{OrderId: "20200501005", OrderName: "火龙果", Status: message.OrderStatus_Refunded, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
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

func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *message.OrderRequest) (res *message.OrderInfo, err error) {
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

func (s *OrderServiceImpl) GetOrderList(ctx context.Context, req *message.OrderSize) (res *message.OrderInfoList, err error) {
	log.Printf("order list service requst is: %v\n", req)

	orderMap := OrderHouse()
	list := make([]*message.OrderInfo, 0)
	for _, v := range *orderMap {
		list = append(list, v)
	}

	return &message.OrderInfoList{
		Orders: list,
	}, nil
}

func (s *OrderServiceImpl) NewOrder(ctx context.Context, req *message.OrderData) (res *message.OrderResponse, err error) {
	log.Printf("create a new order, data is: %v\n", req)

	orderMap := OrderHouse()
	if _, exist := (*orderMap)[req.OrderId]; exist {
		return &message.OrderResponse{
			OrderId: "-1",
			Message: "failed",
		}, nil
	}

	(*orderMap)[req.OrderId] = &message.OrderInfo{OrderId: req.OrderId, OrderName: req.OrderName, Status: req.Status, OrderTime: &timestamp.Timestamp{Seconds: time.Now().Unix()}}
	log.Printf("create a new order success, data is: %v\n", (*orderMap)[req.OrderId])
	return &message.OrderResponse{
		OrderId: (*orderMap)[req.OrderId].OrderId,
		Message: "ok",
	}, nil
}
