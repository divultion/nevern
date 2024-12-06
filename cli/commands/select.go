package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/divultion/nevern/service/id"
	"github.com/divultion/nevern/service/service"
)

var SelectCommand = &cobra.Command{
	Use:   "select",
	Short: "Select one connection of nevern@divultion",
	Long:  "Select one connection of nevern@divultion Reverse Shell",
	Run:   Select,
}

func FindConnectionById(connections []*service.ConnectionData, targetId id.Id) *service.ConnectionData {
	for _, connection := range connections {
		connId, err := id.FromRaw(connection.Id.RawId)
		if err != nil {
			continue
		}

		if connId == targetId {
			return connection
		}
	}
	return nil
}

func Select(cmd *cobra.Command, args []string) {
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

	connections, err := ListConnectionIdsRequest(client)
	if err != nil {
		fmt.Printf("Error listing connections from nevern@divultion (Service) %s\n", err)
		return
	}
	connection := FindConnectionById(connections, id_)
	if connection == nil {
		fmt.Printf("No connection with id: '%s'\n", id_.ToHex())
		return
	}

	session := NewSession(client, connection)
	session.Run()
}
