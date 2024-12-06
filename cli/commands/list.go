package commands

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/divultion/nevern/service/id"
	"github.com/divultion/nevern/service/service"
)

var ListConnectionsCommand = &cobra.Command{
	Use:   "list",
	Short: "List all Peers connected to the Reverse Shell",
	Long:  "List all Peers connected to the running instance of nevern@divultion Reverse Shell",
	Run:   ListConnections,
}

func ListConnectionIdsRequest(client *service.NevernServiceClient) (ids []*service.ConnectionData, err error) {
	in := service.Empty{}

	stream, err := (*client).ListConnectionIds(context.Background(), &in)
	if err != nil {
		return ids, err
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return ids, err
		}
		ids = append(ids, resp)
	}

	return ids, nil
}

func ListConnections(cmd *cobra.Command, args []string) {
	client, err := Connect(cmd)
	if err != nil {
		fmt.Printf("Error connecting to the nevern@divultion (Service) %s\n", err)
		return
	}

	ids, err := ListConnectionIdsRequest(client)
	if err != nil {
		fmt.Printf("Error listing connections from nevern@divultion (Service) %s\n", err)
		return
	}

	DebugConnections(ids)
}

func DebugConnections(ids []*service.ConnectionData) {
	fmt.Println("-----------------------------")
	fmt.Println("Id\tAddress\t\tStatus")

	for _, connectionId := range ids {
		fmt.Println(DebugConnection(connectionId))
	}
}

func DebugConnection(connectionId *service.ConnectionData) string {
	actualId, err := id.FromRaw(connectionId.Id.RawId)

	var messages string
	if connectionId.MessagesAvailable > 0 {
		messages = fmt.Sprintf(" (%d+)", connectionId.MessagesAvailable)
	}
	status := "connected"
	if !connectionId.Connected {
		status = "dis" + status
	}
	var hexId string
	if err != nil {
		hexId = "[invalid]"
	} else {
		hexId = actualId.ToHex()
	}

	return fmt.Sprintf("%s\t%s%s\t\t%s", hexId, connectionId.Address, messages, status)
}
