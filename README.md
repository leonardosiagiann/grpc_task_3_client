# grpc_task_3_client

## Usage

### Run Client

##### Add Deposit

```bash
go run ./client/main.go addDeposit [deposit number value]
```
[deposit number value] must be positive int or float and bigger than 0

#### Get Deposit

```bash
go run ./client/main.go getBalance
```

### Client Example

##### Add Deposit
```bash
go run ./client/main.go addDeposit 1
```
```bash
go run ./client/main.go addDeposit
Please enter the amount
exit status 1
```
```bash
go run ./client/main.go addDeposit -1
Failed to send: rpc error: code = Unknown desc = deposit value must bigger than 0
exit status 1
```

#### Get Deposit
```bash
go run ./client/main.go getBalance
Current Balance: totalDeposit:[total deposit]
```

### Unit Test
```bash
go test -v ./controllers/account
=== RUN   TestDepositServiceClient_Deposit
=== RUN   TestDepositServiceClient_Deposit/Invalid_Request_With_Invalid_Amount
=== RUN   TestDepositServiceClient_Deposit/Valid_Request_With_Valid_Amount
--- PASS: TestDepositServiceClient_Deposit (0.00s)
    --- PASS: TestDepositServiceClient_Deposit/Invalid_Request_With_Invalid_Amount (0.00s)
    --- PASS: TestDepositServiceClient_Deposit/Valid_Request_With_Valid_Amount (0.00s)
=== RUN   TestDepositServiceClient_GetDeposit
=== RUN   TestDepositServiceClient_GetDeposit/Valid_Test_Get_Deposit
--- PASS: TestDepositServiceClient_GetDeposit (0.00s)
    --- PASS: TestDepositServiceClient_GetDeposit/Valid_Test_Get_Deposit (0.00s)
PASS
ok  	grpc_client/controllers/account	(cached)
```
