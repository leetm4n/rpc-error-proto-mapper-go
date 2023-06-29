package encoder_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/leetm4n/rpc-error-proto-mapper-go/pkg/encoder"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
)

func getMarshalledStatus(
	errInfoDetails *errdetails.ErrorInfo,
	badRequestDetails *errdetails.BadRequest,
	retryInfoDetail *errdetails.RetryInfo,
	debugInfoDetails *errdetails.DebugInfo,
	quotaFailureDetails *errdetails.QuotaFailure,
	preconditionFailureDetails *errdetails.PreconditionFailure,
	resourceInfoDetails *errdetails.ResourceInfo,
) []byte {
	var details []*anypb.Any

	if errInfoDetails != nil {
		errDetailsAny, _ := anypb.New(errInfoDetails)
		details = append(details, errDetailsAny)
	}
	if badRequestDetails != nil {
		badRequestDetailsAny, _ := anypb.New(badRequestDetails)
		details = append(details, badRequestDetailsAny)
	}
	if retryInfoDetail != nil {
		retryInfoDetailAny, _ := anypb.New(retryInfoDetail)
		details = append(details, retryInfoDetailAny)
	}
	if debugInfoDetails != nil {
		debugInfoDetailsAny, _ := anypb.New(debugInfoDetails)
		details = append(details, debugInfoDetailsAny)
	}
	if quotaFailureDetails != nil {
		quotaFailureDetailsAny, _ := anypb.New(quotaFailureDetails)
		details = append(details, quotaFailureDetailsAny)
	}
	if preconditionFailureDetails != nil {
		preconditionFailureDetailsAny, _ := anypb.New(preconditionFailureDetails)
		details = append(details, preconditionFailureDetailsAny)
	}
	if resourceInfoDetails != nil {
		resourceInfoDetailsAny, _ := anypb.New(resourceInfoDetails)
		details = append(details, resourceInfoDetailsAny)
	}

	payload, _ := proto.Marshal(&status.Status{
		Code:    int32(code.Code_INTERNAL),
		Details: details,
		Message: "error",
	})
	return payload
}

func getMarshalledStatusWithCustomDetail(customDetail proto.Message) []byte {
	any, _ := anypb.New(customDetail)

	payload, _ := proto.Marshal(&status.Status{
		Code: int32(code.Code_INTERNAL),
		Details: []*anypb.Any{
			any,
		},
		Message: "error",
	})
	return payload
}

func getMarshalledStatusWithCustomAnyType(customAnyType string) []byte {
	any := &anypb.Any{
		TypeUrl: customAnyType,
		Value:   []byte("test"),
	}

	payload, _ := proto.Marshal(&status.Status{
		Code: int32(code.Code_INTERNAL),
		Details: []*anypb.Any{
			any,
		},
		Message: "error",
	})
	return payload
}

func protoValueOfErrorInfo(errDetails *errdetails.ErrorInfo) []byte {
	errDetailsAny, _ := anypb.New(errDetails)

	return errDetailsAny.Value
}

func protoValueOfBadRequest(errDetails *errdetails.BadRequest) []byte {
	errDetailsAny, _ := anypb.New(errDetails)

	return errDetailsAny.Value
}

func TestEncode(t *testing.T) {
	type args struct {
		code    code.Code
		details []*anypb.Any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "should encode with empty details",
			args: args{
				code:    code.Code_INTERNAL,
				details: []*anypb.Any{},
			},
			want:    getMarshalledStatus(nil, nil, nil, nil, nil, nil, nil),
			wantErr: false,
		},
		{
			name: "should encode with nil details",
			args: args{
				code:    code.Code_INTERNAL,
				details: nil,
			},
			want:    getMarshalledStatus(nil, nil, nil, nil, nil, nil, nil),
			wantErr: false,
		},
		{
			name: "should encode with multiple details",
			args: args{
				code: code.Code_INTERNAL,
				details: []*anypb.Any{
					{
						TypeUrl: "type.googleapis.com/google.rpc.ErrorInfo",
						Value:   protoValueOfErrorInfo(&errdetails.ErrorInfo{Reason: "test", Domain: "test"}),
					},
					{
						TypeUrl: "type.googleapis.com/google.rpc.BadRequest",
						Value:   protoValueOfBadRequest(&errdetails.BadRequest{FieldViolations: []*errdetails.BadRequest_FieldViolation{{Field: "test", Description: "test"}}}),
					},
				},
			},
			want: getMarshalledStatus(
				&errdetails.ErrorInfo{
					Reason: "test",
					Domain: "test",
				},
				&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{
							Field:       "test",
							Description: "test",
						},
					},
				},
				nil, nil, nil, nil, nil,
			),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encoder.Encode(tt.args.code, tt.args.details)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *encoder.Details
		wantErr bool
	}{
		{
			name: "should be able to decode empty payload",
			args: args{
				payload: []byte{},
			},
			want:    &encoder.Details{},
			wantErr: false,
		},
		{
			name: "should fail if payload is not decodable to status",
			args: args{
				payload: []byte{1, 3},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should fail if error detail is not within the list of error details supported",
			args: args{
				payload: getMarshalledStatusWithCustomDetail(&errdetails.Help{}),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should fail if detail is not known for any type",
			args: args{
				payload: getMarshalledStatusWithCustomAnyType("type.googleapis.com/unknown"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should be able to decode non empty, partially filled payload",
			args: args{
				payload: getMarshalledStatus(&errdetails.ErrorInfo{
					Reason: "test",
					Domain: "test",
					Metadata: map[string]string{
						"test": "test",
					},
				}, nil, nil, nil, nil, nil, nil),
			},
			want: &encoder.Details{
				Code: code.Code_INTERNAL,
				ErrorInfoDetails: &errdetails.ErrorInfo{
					Reason: "test",
					Domain: "test",
					Metadata: map[string]string{
						"test": "test",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should be able to decode non empty, fully filled payload",
			args: args{
				payload: getMarshalledStatus(
					&errdetails.ErrorInfo{
						Reason: "test",
						Domain: "test",
						Metadata: map[string]string{
							"test": "test",
						},
					},
					&errdetails.BadRequest{FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{Field: "test", Description: "test"},
					}},
					&errdetails.RetryInfo{
						RetryDelay: &durationpb.Duration{
							Seconds: 1,
						},
					},
					&errdetails.DebugInfo{Detail: "test"},
					&errdetails.QuotaFailure{Violations: []*errdetails.QuotaFailure_Violation{
						{Subject: "test", Description: "test"},
					}},
					&errdetails.PreconditionFailure{Violations: []*errdetails.PreconditionFailure_Violation{
						{Type: "test", Subject: "test", Description: "test"},
					}},
					&errdetails.ResourceInfo{ResourceType: "test", ResourceName: "test", Owner: "test", Description: "test"},
				),
			},
			want: &encoder.Details{
				Code: code.Code_INTERNAL,
				ErrorInfoDetails: &errdetails.ErrorInfo{
					Reason: "test",
					Domain: "test",
					Metadata: map[string]string{
						"test": "test",
					},
				},
				RetryInfoDetails: &errdetails.RetryInfo{
					RetryDelay: &durationpb.Duration{
						Seconds: 1,
					},
				},
				BadRequestDetails: &errdetails.BadRequest{FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{Field: "test", Description: "test"},
				}},
				DebugInfoDetails: &errdetails.DebugInfo{Detail: "test"},
				QuotaFailureDetails: &errdetails.QuotaFailure{Violations: []*errdetails.QuotaFailure_Violation{
					{Subject: "test", Description: "test"},
				}},
				PreconditionFailureDetails: &errdetails.PreconditionFailure{Violations: []*errdetails.PreconditionFailure_Violation{
					{Type: "test", Subject: "test", Description: "test"},
				}},
				ResourceInfoDetails: &errdetails.ResourceInfo{ResourceType: "test", ResourceName: "test", Owner: "test", Description: "test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encoder.Decode(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !isDetailsEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isDetailsEqual(a, b *encoder.Details) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)

	return reflect.DeepEqual(aJSON, bJSON)
}
