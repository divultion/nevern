package service

import (
	"encoding/hex"
	"fmt"

	"github.com/divultion/nevern/service/id"
	"github.com/divultion/nevern/service/shell"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func InternalToExternalError(id id.Id, err error) error {
	if err == nil {
		return nil
	}

	rawId := id.ToRaw()
	_, ok := err.(*shell.ConnectionNotFound)
	if ok {
		return NewNotFoundIdError(rawId[:])
	}
	if err == shell.PermanentDisconnect {
		return NewPermanentDisconnectError(rawId[:])
	}
	if err == shell.TemporaryDisconnect {
		return NewTemporaryDisconnectError(rawId[:])
	}

	st := status.New(codes.Unknown, err.Error())
	return st.Err()
}

type InvalidConnectionIdError struct {
	RawId       []byte `json:"id"`
	Description string `json:"description"`
}

func NewInvalidConnectionIdError(rawId []byte, description string) InvalidConnectionIdError {
	return InvalidConnectionIdError{RawId: rawId, Description: description}
}

func NewNotFoundIdError(rawId []byte) InvalidConnectionIdError {
	return InvalidConnectionIdError{RawId: rawId, Description: fmt.Sprintf("No connection with id %s", hex.EncodeToString(rawId))}
}

func (e InvalidConnectionIdError) Error() string {
	return fmt.Sprintf("[%s]: %s", hex.EncodeToString(e.RawId), e.Description)
}

func (e InvalidConnectionIdError) GRPCStatus() *status.Status {
	v := &errdetails.BadRequest_FieldViolation{
		Field:       "RawId",
		Description: e.Description,
	}
	br := &errdetails.BadRequest{}
	br.FieldViolations = append(br.FieldViolations, v)
	st, _ := status.New(codes.InvalidArgument, e.Error()).WithDetails(br)
	return st
}

type PermanentDisconnectError struct {
	RawId []byte `json:"id"`
}

func NewPermanentDisconnectError(rawId []byte) PermanentDisconnectError {
	return PermanentDisconnectError{RawId: rawId}
}

func (e PermanentDisconnectError) Error() string {
	return fmt.Sprintf("%s disconnected (0)", hex.EncodeToString(e.RawId))
}

func (e PermanentDisconnectError) GRPCStatus() *status.Status {
	return status.New(codes.Aborted, e.Error())
}

type TemporaryDisconnectError struct {
	RawId []byte `json:"id"`
}

func NewTemporaryDisconnectError(rawId []byte) TemporaryDisconnectError {
	return TemporaryDisconnectError{RawId: rawId}
}

func (e TemporaryDisconnectError) Error() string {
	return fmt.Sprintf("%s disconnected (temporary)", hex.EncodeToString(e.RawId))
}

func (e TemporaryDisconnectError) GRPCStatus() *status.Status {
	return status.New(codes.Unavailable, e.Error())
}
