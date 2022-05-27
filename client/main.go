package main

import (
	"context"
	"log"
	"os"
	"strconv"

	DepositClient "grpc_client/controllers/account"
	pAccount "grpc_client/proto/account"

	"google.golang.org/grpc"
)

const address = "localhost:8080"

func main() {

	if len(os.Args) < 2 {
		log.Println("You must be provide a valid command")
		log.Fatal("Use get to retrieve account balance, and send to send amount")
		return
	}

	if os.Args[1] != "getBalance" && os.Args[1] != "addDeposit" {
		log.Println("You must be provide a valid command")
		log.Fatal("Use get to retrieve account balance, and send to send amount")
		return
	}

	if os.Args[1] == "addDeposit" {
		if len(os.Args) <= 2 {
			log.Println("You must be provide a valid amount")
			return
		} else if len(os.Args) >= 3 {
			_, err := strconv.ParseFloat(os.Args[2], 32)
			if err != nil {
				log.Fatal("Invalid Amount Args, Must be float or int")
				return
			}
		}
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	client := DepositClient.NewDepositoClient(conn)

	if os.Args[1] == "addDeposit" {
		amount, _ := strconv.ParseFloat(os.Args[2], 32)
		res, err := client.Deposit(context.Background(), pAccount.DepositRequest{
			Amount: float32(amount),
		})
		if err != nil {
			log.Fatalf("Failed to send: %v", err)
			return
		}
		if res == true {
			log.Println("Balance Added")
			return
		}

	}
	if os.Args[1] == "getBalance" {
		res, err := client.GetDeposit(context.Background())
		if err != nil {
			log.Fatalf("Failed to get deposit ammount: %v", err)
			return
		}
		log.Printf("Current Balance: %v", res)
	}
}
