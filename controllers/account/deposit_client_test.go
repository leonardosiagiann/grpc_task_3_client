package account

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	proto "grpc_client/proto/account"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type mockDepositServiceServer struct {
	proto.UnimplementedDepositServiceServer
}

type Deposit struct {
	DepositAmmount float32
}

var Deposito Deposit

func dial() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	proto.RegisterDepositServiceServer(server, &mockDepositServiceServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

}

func (d *mockDepositServiceServer) Deposit(c context.Context, in *proto.DepositRequest) (*proto.DepositResponse, error) {
	if in.GetAmount() <= 0 {
		return nil, errors.New("deposit value must bigger than 0")
	}

	Deposito.DepositAmmount = Deposito.DepositAmmount + in.GetAmount()
	ok := true
	response := &proto.DepositResponse{
		Ok: ok,
	}
	return response, nil
}

func (d *mockDepositServiceServer) GetDeposit(ctx context.Context, in *proto.GetDepositRequest) (*proto.GetDepositResponse, error) {
	response := proto.GetDepositResponse{
		TotalDeposit: Deposito.DepositAmmount,
	}
	return &response, nil
}

func TestDepositServiceClient_Deposit(t *testing.T) {
	test := []struct {
		name   string
		amount float32
		res    bool
		err    string
	}{
		{
			"Invalid Request With Invalid Amount",
			-1,
			false,
			"deposit value must bigger than 0",
		},
		{
			"Valid Request With Valid Amount",
			1,
			true,
			"",
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			var value bool
			request := &proto.DepositRequest{Amount: tc.amount}
			ctx := context.Background()

			client := NewDepositoClient(conn)
			response, err := client.Deposit(ctx, request)

			if response != nil {
				val := response.(*proto.DepositResponse)
				value = val.Ok
			}

			if value != tc.res {
				t.Error("error : expected", tc.res, "received:", value)
			}
			if err != nil {
				if er, _ := status.FromError(err); er.Message() != tc.err {
					t.Error("error msg expected: ", tc.err, "received", er.Message())
				}
			}
		})
	}
}

func TestDepositServiceClient_GetDeposit(t *testing.T) {
	test := struct {
		name string
		res  *proto.GetDepositResponse
		err  string
	}{
		"Valid Test Get Deposit",
		&proto.GetDepositResponse{TotalDeposit: 1},
		"",
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dial()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	t.Run(test.name, func(t *testing.T) {
		dclient := NewDepositoClient(conn)
		response, err := dclient.GetDeposit(context.Background())
		val := response.(*proto.GetDepositResponse)
		if val.GetTotalDeposit() != test.res.GetTotalDeposit() {
			t.Error("error : expected", test.res, "received:", val)
		}
		if err != nil {
			if er, _ := status.FromError(err); er.Message() != test.err {
				t.Error("error msg expected: ", test.err, "received", er.Message())
			}
		}
	})
}
