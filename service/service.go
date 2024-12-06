package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/divultion/nevern/service/service"
	"github.com/divultion/nevern/service/shell"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "nevern@divultion",
	Short: "Nevern is multi-client reverse shell",
	Long:  `This is a reverse shell service for nevern@divultion. Nevern is a multi-client, multi-threaded reverse shell written in Go. We are aiming for backcompitability with all existing NetCat reverse shells.`,
	Run: func(cmd *cobra.Command, args []string) {
		service_host, err := cmd.Flags().GetString("shost")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		service_port, err := cmd.Flags().GetString("sp")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		port, err := cmd.Flags().GetString("p")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		ch, err := cmd.Flags().GetString("ch")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		channLength, err := strconv.Atoi(ch)
		if err != nil {
			fmt.Printf("--ch flag must be an int: %s\n", err)
			return
		}

		buf, err := cmd.Flags().GetString("buf")
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		bufferLength, err := strconv.Atoi(buf)
		if err != nil {
			fmt.Printf("--buf flag must be an int: %s\n", err)
			return
		}

		fmt.Printf("Running nevern@divultion Reverse Shell at %s:%s.\n", host, port)
		reverseShell, err := shell.ShellTCPListen(host, port, channLength, bufferLength)
		if err != nil {
			fmt.Printf("Error running Reverse Shell %s\n", err)
			return
		}

		go reverseShell.ListenForConnections()

		fmt.Printf("Running nevern@divultion Service at %s:%s.\n\n", service_host, service_port)
		listener, err := net.Listen("tcp", service_host+":"+service_port)
		if err != nil {
			fmt.Printf("Error running Service %s\n", err)
			return
		}

		s := grpc.NewServer()
		service.RegisterNevernServiceServer(s, service.NewService(reverseShell))

		fmt.Printf("Use nevern-cli@divultion to access connected peers.\n")
		if err := s.Serve(listener); err != nil {
			fmt.Printf("Error running Service %s\n", err)
			return
		}
	},
}

func main() {
	rootCmd.PersistentFlags().String("host", "0.0.0.0", "Set up Host where the reverse shell program will be running")
	rootCmd.PersistentFlags().String("shost", "0.0.0.0", "Set up Host where the service program will be running")
	rootCmd.PersistentFlags().String("sp", "3303", "Set up Port where the service program will be running")
	rootCmd.PersistentFlags().String("p", "8080", "Set up Port where the reverse shell program will be running")
	rootCmd.PersistentFlags().String("ch", "99", "Set up channel length for connection")
	rootCmd.PersistentFlags().String("buf", "1024", "Set up buffer length for connection")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
