package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/divultion/nevern/service/id"
	"github.com/divultion/nevern/service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DisconnectCommand = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect from a peer of nevern@divultion",
	Long:  "Disconnect from a peer of nevern@divultion",
	Run:   Disconnect,
}

func Disconnect(cmd *cobra.Command, args []string) {
	client, err := Connect(cmd)
	if err != nil {
		fmt.Printf("Error connecting to the nevern@divultion (Service) %s\n", err)
		return
	}

	if len(args) < 1 {
		fmt.Printf("No connection id passed.\n")
		return
	}

	id_, err := id.FromHex(args[0])
	if err != nil {
		fmt.Printf("Couldnt parse id due to %s\n", err)
		return
	}
	rawId := id_.ToRaw()

	_, err = (*client).DisconnectById(context.Background(), &service.ConnectionId{RawId: rawId[:]})
	if err != nil {
		code := status.Code(err)
		if code == codes.InvalidArgument {
			fmt.Printf("%s\n", err)
			return
		}

		if code == codes.Aborted || code == codes.Unavailable {
			return
		}

		fmt.Printf("Disconnect error %s\n", err)
		return
	}
}
