module plugins/todo

go 1.21.1

require (
	github.com/shimabukuromeg/connect-go-todo-example v0.0.0-20230923104853-c6549f378c5e
	google.golang.org/grpc v1.58.2
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace google.golang.org/grpc v1.58.2 => google.golang.org/grpc v1.58.1
