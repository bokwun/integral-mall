package integralrpcmodel

import (
	"context"
	"integral-mall/integral/protos"

	"github.com/yakaa/grpcx"
)

type (
	IntegralRpcModel struct {
		cli *grpcx.GrpcxClient
	}
)

func NewIntegralRpcModel(cli *grpcx.GrpcxClient) *IntegralRpcModel {
	return &IntegralRpcModel{
		cli: cli,
	}
}

func (m *IntegralRpcModel) AddIntegral(userId, integral int) error {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return err
	}
	clientIntegral := protos.NewIntegralRpcClient(conn)
	if _, err := clientIntegral.AddIntegral(context.Background(), &protos.AddIntegralRequest{
		UserId:   int64(userId),
		Integral: int64(integral),
	}); err != nil {
		return err
	}
	return nil
}

func (m *IntegralRpcModel) ConsumerIntegral(userId, consumerIntegral int) error {
	conn, err := m.cli.GetConnection()
	if err != nil {
		return err
	}
	clientIntegral := protos.NewIntegralRpcClient(conn)
	if _, err := clientIntegral.ConsumerIntegral(context.Background(), &protos.ConsumerIntegralRequest{
		UserId:           int64(userId),
		ConsumerIntegral: int64(consumerIntegral),
	}); err != nil {
		return err
	}
	return nil
}
