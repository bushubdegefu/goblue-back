package manager

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"semay.com/bluerpc"
	"semay.com/config"
)

var (
	startrpc = &cobra.Command{
		Use:   "rpcserve",
		Short: "Start RPC server from the app at the provied Port",
		Long:  `Start RPC server from the app at the provied Port`,
		Run: func(cmd *cobra.Command, args []string) {
			RpcServe()
		},
	}

	rpcclient = &cobra.Command{
		Use:   "rpcclient",
		Short: "Make  RPC call server from the app at the provied Port. For testing",
		Long:  `Make RPC call to server from the app at the provied Port through .env variable`,
		Run: func(cmd *cobra.Command, args []string) {
			RpcClient()
		},
	}
)

func RpcServe() {
	lis, err := net.Listen("tcp", "0.0.0.0:"+config.Config("RPC_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v\n", err, config.Config("RPC_PORT"))
	}

	blueserver := bluerpc.BlueRPCServer{}
	grpcServer := grpc.NewServer()

	// bluerpc.RegisterChatServiceServer(grpcServer, &blueserver)
	bluerpc.RegisterBlueServiceServer(grpcServer, &blueserver)

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
	fmt.Println("Started RPC Server for BLUE")

}

func RpcClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.Config("RPC_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := bluerpc.NewBlueServiceClient(conn)

	message := bluerpc.BlueAppID{
		AppId: uuid.New().String(),
	}

	response, err := c.GetSalt(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling Getting Salt: %s", err)
	}

	log.Printf("Response from Server:\n %s\n", response)

}

func init() {
	goBlueCmd.AddCommand(startrpc)
	goBlueCmd.AddCommand(rpcclient)

}
