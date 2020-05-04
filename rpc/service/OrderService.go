package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/JosephChan007/go-rpc/rpc/message"
)

func OrderHouse() *map[string]*message.OrderInfo {
	orderMap := make(map[string]*message.OrderInfo, 5)
	orderMap["1"] = &message.OrderInfo{OrderId: "20200501001", OrderName: "橙子", Status: message.OrderStatus_UnPay}
	orderMap["2"] = &message.OrderInfo{OrderId: "20200501002", OrderName: "苹果", Status: message.OrderStatus_Paied}
	orderMap["3"] = &message.OrderInfo{OrderId: "20200501003", OrderName: "梨", Status: message.OrderStatus_Commited}
	orderMap["4"] = &message.OrderInfo{OrderId: "20200501004", OrderName: "香蕉", Status: message.OrderStatus_Refund}
	orderMap["5"] = &message.OrderInfo{OrderId: "20200501005", OrderName: "火龙果", Status: message.OrderStatus_Refunded}
	return &orderMap
}

// For rpc
/*
type OrderService struct {
}

func (s *OrderService) GetOrder(req *message.OrderRequest, res *message.OrderInfo) error {
	fmt.Printf("order service requst is: %v\n", req)

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
	fmt.Printf("order service requst is: %v\n", req)

	orderMap := OrderHouse()
	for _, v := range *orderMap {
		if v.Status == req.Status {
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
	fmt.Printf("order list service requst is: %v\n", req)

	orderMap := OrderHouse()
	return &message.OrderInfoList{
		Orders: []*message.OrderInfo{
			(*orderMap)["1"],
			(*orderMap)["3"],
			(*orderMap)["5"],
		},
	}, nil
}
