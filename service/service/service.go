package service

import (
	context "context"
	"fmt"

	"github.com/divultion/nevern/service/id"
	"github.com/divultion/nevern/service/shell"
)

type Service struct {
	reverse_shell *shell.ShellTCP
}

func NewService(reverse_shell *shell.ShellTCP) Service {
	return Service{reverse_shell: reverse_shell}
}

// mustEmbedUnimplementedNevernServiceServer implements NevernServiceServer.
func (s Service) mustEmbedUnimplementedNevernServiceServer() {
	panic("unimplemented")
}

func (s Service) ListConnectionIds(in *Empty, srv NevernService_ListConnectionIdsServer) error {
	for _, id_ := range s.reverse_shell.GetAllConnectionIds() {
		rawId := id_.ToRaw()

		address, err := s.reverse_shell.GetConnectionAddressById(id_)
		if err != nil {
			fmt.Printf("GetConnectionAddressById err %s\n", err)
			continue
		}
		messagesAvailable, err := s.reverse_shell.GetConnectionMessagesAvailableById(id_)
		if err != nil {
			fmt.Printf("GetConnectionMessagesAvailableById err %s\n", err)
			continue
		}
		connected, err := s.reverse_shell.GetConnectedById(id_)
		if err != nil {
			if err == shell.PermanentDisconnect {
				continue
			}
			fmt.Printf("GetConnectedById err %s\n", err)
			continue
		}
		resp := ConnectionData{
			Id:                &ConnectionId{RawId: rawId[:]},
			Address:           address,
			MessagesAvailable: uint32(messagesAvailable),
			Connected:         connected,
		}
		if err := srv.Send(&resp); err != nil {
			fmt.Printf("send err %s\n", err)
		}
	}

	return nil
}

func (s Service) TryReadOutputById(ctx context.Context, connectionId *ConnectionId) (*Output, error) {
	id, err := id.FromRaw(connectionId.RawId)
	if err != nil {
		return nil, NewInvalidConnectionIdError(connectionId.RawId, err.Error())
	}

	data, ok, err := s.reverse_shell.TryReadOutputById(id)
	err = InternalToExternalError(id, err)
	if err != nil {
		if ok {
			return &Output{Data: data, Ok: ok}, nil
		}
		return nil, err
	}

	return &Output{Data: data, Ok: ok}, nil
}

func (s Service) WriteInputById(ctx context.Context, in *Input) (*WriteInputByIdResponse, error) {
	id, err := id.FromRaw(in.Id.RawId)
	if err != nil {
		return nil, NewInvalidConnectionIdError(in.Id.RawId, err.Error())
	}

	i, err := s.reverse_shell.WriteInputById(in.Data, id)
	err = InternalToExternalError(id, err)
	if err != nil {
		return nil, err
	}

	return &WriteInputByIdResponse{DataWritten: int64(i)}, nil
}

func (s Service) ForgetById(ctx context.Context, in *ConnectionId) (*Empty, error) {
	id, err := id.FromRaw(in.RawId)
	if err != nil {
		return nil, NewInvalidConnectionIdError(in.RawId, err.Error())
	}

	err = InternalToExternalError(id, s.reverse_shell.ForgetById(id))
	if err != nil {
		return nil, err
	}

	return &Empty{}, nil
}

func (s Service) DisconnectById(ctx context.Context, in *ConnectionId) (*Empty, error) {
	id, err := id.FromRaw(in.RawId)
	if err != nil {
		return nil, NewInvalidConnectionIdError(in.RawId, err.Error())
	}

	err = InternalToExternalError(id, s.reverse_shell.DisconnectById(id))
	if err != nil {
		return nil, err
	}

	return &Empty{}, nil
}
