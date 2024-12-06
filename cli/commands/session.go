package commands

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/divultion/nevern/service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Session struct {
	service    *service.NevernServiceClient
	connection *service.ConnectionData
}

func NewSession(service *service.NevernServiceClient, connection *service.ConnectionData) Session {
	fmt.Printf("Setting up session with %s.\n", connection.Address)
	return Session{service: service, connection: connection}
}

func (session *Session) Run() {
	var wg sync.WaitGroup
	killswitch := make(chan struct{})
	wg.Add(1)
	go OutputLogger(session.service, session.connection, killswitch, &wg)

	if !session.connection.Connected {
		wg.Wait()
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Couldnt readline %s\n", err)
			continue
		}

		commandStr := strings.TrimSpace(command)

		in := service.Input{Data: commandStr, Id: session.connection.Id}
		_, err = (*session.service).WriteInputById(context.Background(), &in)
		if err != nil {
			if status.Code(err) == codes.Unavailable {
				wg.Wait()
				return
			}
			fmt.Printf("> %s\n", err)
			continue
		}
	}
}

func OutputLogger(client *service.NevernServiceClient, connection *service.ConnectionData, killswitch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if IsKsDone(killswitch) {
			return
		}

		in := connection.Id
		output, err := (*client).TryReadOutputById(context.Background(), in)
		if err != nil {
			code := status.Code(err)
			if code == codes.Aborted || code == codes.Unavailable || code == codes.InvalidArgument {
				return
			}

			fmt.Printf("try read output error %s: %s\n", code.String(), err)
			return
		}
		if !output.Ok {
			continue
		}
		fmt.Printf("%s", output.Data)
	}
}

func IsKsDone(killswitch <-chan struct{}) bool {
	select {
	case <-killswitch:
		return true
	default:
		return false
	}
}
