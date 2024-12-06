package commands

import (
	"github.com/spf13/cobra"
	"github.com/divultion/nevern/service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(cmd *cobra.Command) (*service.NevernServiceClient, error) {
	host, err := cmd.Flags().GetString("shost")
	if err != nil {
		return nil, err
	}
	port, err := cmd.Flags().GetString("sp")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := service.NewNevernServiceClient(conn)

	return &client, err
}
