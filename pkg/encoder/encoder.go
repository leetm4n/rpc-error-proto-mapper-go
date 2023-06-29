package encoder

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Details struct {
	Code                       code.Code
	ErrorInfoDetails           *errdetails.ErrorInfo
	RetryInfoDetails           *errdetails.RetryInfo
	BadRequestDetails          *errdetails.BadRequest
	DebugInfoDetails           *errdetails.DebugInfo
	QuotaFailureDetails        *errdetails.QuotaFailure
	PreconditionFailureDetails *errdetails.PreconditionFailure
	ResourceInfoDetails        *errdetails.ResourceInfo
}

func Encode(code code.Code, details []*anypb.Any) ([]byte, error) {
	marshalled, err := proto.Marshal(&status.Status{
		Code:    int32(code),
		Details: details,
		Message: "error",
	})
	if err != nil {
		return nil, err
	}

	return marshalled, nil
}

func Decode(payload []byte) (*Details, error) {
	s := &status.Status{}
	err := proto.Unmarshal(payload, s)
	if err != nil {
		return nil, err
	}

	var errorInfoDetails *errdetails.ErrorInfo
	var badRequestDetails *errdetails.BadRequest
	var retryInfoDetails *errdetails.RetryInfo
	var debugInfoDetails *errdetails.DebugInfo
	var quotaFailureDetails *errdetails.QuotaFailure
	var preconditionFailureDetails *errdetails.PreconditionFailure
	var resourceInfoDetails *errdetails.ResourceInfo

	for _, detail := range s.Details {
		unmarshalled, err := anypb.UnmarshalNew(detail, proto.UnmarshalOptions{})
		if err != nil {
			return nil, err
		}

		switch v := unmarshalled.(type) {
		case *errdetails.ErrorInfo:
			errorInfoDetails = v
		case *errdetails.BadRequest:
			badRequestDetails = v
		case *errdetails.RetryInfo:
			retryInfoDetails = v
		case *errdetails.DebugInfo:
			debugInfoDetails = v
		case *errdetails.QuotaFailure:
			quotaFailureDetails = v
		case *errdetails.PreconditionFailure:
			preconditionFailureDetails = v
		case *errdetails.ResourceInfo:
			resourceInfoDetails = v
		default:
			return nil, fmt.Errorf("unknown error detail: %s", v)
		}
	}

	return &Details{
		Code:                       code.Code(s.Code),
		ErrorInfoDetails:           errorInfoDetails,
		BadRequestDetails:          badRequestDetails,
		DebugInfoDetails:           debugInfoDetails,
		QuotaFailureDetails:        quotaFailureDetails,
		RetryInfoDetails:           retryInfoDetails,
		PreconditionFailureDetails: preconditionFailureDetails,
		ResourceInfoDetails:        resourceInfoDetails,
	}, nil
}
