package errors

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes/any"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errMsgSep string = ":    "
)

// Error includes Code and Message
type Error struct {
	Code    int32
	Message string
	Service string
	s       *spb.Status
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	bs, _ := json.Marshal(e)
	return string(bs)
}

// grpcEncode : encodes to grpc error format
func (e *Error) grpcEncode() {
	grpcStatus := new(spb.Status)
	grpcStatus.Code = int32(e.Code)
	grpcStatus.Message = e.Message
	grpcStatus.Details = make([]*any.Any, 0)
	any := new(any.Any)
	if e.Service == "" {
		any.TypeUrl = "unknown Service"
	} else {
		any.TypeUrl = e.Service
	}
	grpcStatus.Details = append(grpcStatus.Details, any)
	e.s = grpcStatus
}

// New : return a new pointer of Error
func New(srvName string, Code int32, Message string) error {
	e := new(Error)
	e.Code = Code
	e.Message = Message
	e.Service = srvName
	e.grpcEncode()

	return status.FromProto(e.s).Err()
}

// WithMessage : annotates error with a new message
func WithMessage(err error, messages ...string) error {
	e := FromError(err)
	for _, msg := range messages {
		e.Message = msg + errMsgSep + e.Message
	}
	e.grpcEncode()
	return status.FromProto(e.s).Err()
}

// FromError : convert error to *Error
func FromError(err error) *Error {
	if _, ok := err.(*Error); ok {
		return err.(*Error)
	}

	e := new(Error)

	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		e.Code = int32(codes.Unknown)
		e.Message = err.Error()
		e.Service = "unknown Service"
	} else {
		e.Code = int32(st.Code())
		e.Message = st.Message()
		any := st.Proto().GetDetails()[0]
		e.Service = any.GetTypeUrl()
	}

	return e
}
