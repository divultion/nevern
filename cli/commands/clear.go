package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ClearCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clears connections of nevern@divultion",
	Long:  "Clears connections, forget disconnected of nevern@divultion",
	Run:   Clear,
}

func Clear(cmd *cobra.Command, args []string) {
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

	for _, connectionId := range ids {
		if !connectionId.Connected {
			fmt.Println("------------------------------")
			session := NewSession(client, connectionId)
			session.Run()
		}
	}
}
