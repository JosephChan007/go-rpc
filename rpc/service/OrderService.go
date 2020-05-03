package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/JosephChan007/go-rpc/rpc/message"
)

func OrderHouse() *map[string]message.OrderInfo {
	orderMap := make(map[string]message.OrderInfo, 5)
	orderMap["1"] = message.OrderInfo{OrderId: "20200501001", OrderName: "橙子", OrderStatus: "待付款"}
	orderMap["2"] = message.OrderInfo{OrderId: "20200501002", OrderName: "苹果", OrderStatus: "已付款"}
	orderMap["3"] = message.OrderInfo{OrderId: "20200501003", OrderName: "梨", OrderStatus: "已提交"}
	orderMap["4"] = message.OrderInfo{OrderId: "20200501004", OrderName: "香蕉", OrderStatus: "退款中"}
	orderMap["5"] = message.OrderInfo{OrderId: "20200501005", OrderName: "火龙果", OrderStatus: "已退款"}
	return &orderMap
}

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

type OrderServiceImpl struct {
}

func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *message.OrderRequest) (res *message.OrderInfo, err error) {
	fmt.Printf("order service requst is: %v\n", req)

	orderMap := OrderHouse()
	for _, v := range *orderMap {
		if v.OrderId == req.OrderId {
			res = &v
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
