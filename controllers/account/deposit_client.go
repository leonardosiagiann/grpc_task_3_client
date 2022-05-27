package account

import (
	"context"
	"time"

	account "grpc_client/proto/account"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

type DepositClient struct {
	conn   grpc.ClientConn
	client account.DepositServiceClient
}

func NewDepositoClient(conn *grpc.ClientConn) *DepositClient {
	return &DepositClient{
		conn:   grpc.ClientConn{},
		client: account.NewDepositServiceClient(conn),
	}
}

func (d *DepositClient) Deposit(ctx context.Context, in interface{}) (interface{}, error) {
	var request *account.DepositRequest
	err := mapstructure.Decode(in, &request)
	if err != nil {
		return nil, err
	}

	ctxOutgoing, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := d.client.Deposit(ctxOutgoing, request)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DepositClient) GetDeposit(ctx context.Context) (interface{}, error) {
	ctxOutgoing, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := d.client.GetDeposit(ctxOutgoing, &account.GetDepositRequest{})
	if err != nil {
		return nil, err
	}

	return data, nil
}
