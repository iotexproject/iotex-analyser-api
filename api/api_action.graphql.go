// Code generated by proroc-gen-graphql, DO NOT EDIT.
package api

import (
	"context"

	"github.com/graphql-go/graphql"
	pagination "github.com/iotexproject/iotex-analyser-api/api/pagination"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_XrcInfo                                    *graphql.Object      // message XrcInfo in api_action.proto
	gql__type_EvmTransfersByAddressResponse_EvmTransfer  *graphql.Object      // message EvmTransfersByAddressResponse.EvmTransfer in api_action.proto
	gql__type_EvmTransfersByAddressResponse              *graphql.Object      // message EvmTransfersByAddressResponse in api_action.proto
	gql__type_EvmTransfersByAddressRequest               *graphql.Object      // message EvmTransfersByAddressRequest in api_action.proto
	gql__type_EvmTransferInfo                            *graphql.Object      // message EvmTransferInfo in api_action.proto
	gql__type_ActionResponse                             *graphql.Object      // message ActionResponse in api_action.proto
	gql__type_ActionRequest                              *graphql.Object      // message ActionRequest in api_action.proto
	gql__type_ActionInfo                                 *graphql.Object      // message ActionInfo in api_action.proto
	gql__type_ActionByTypeResponse                       *graphql.Object      // message ActionByTypeResponse in api_action.proto
	gql__type_ActionByTypeRequest                        *graphql.Object      // message ActionByTypeRequest in api_action.proto
	gql__type_ActionByHashResponse_EvmTransfers          *graphql.Object      // message ActionByHashResponse.EvmTransfers in api_action.proto
	gql__type_ActionByHashResponse                       *graphql.Object      // message ActionByHashResponse in api_action.proto
	gql__type_ActionByHashRequest                        *graphql.Object      // message ActionByHashRequest in api_action.proto
	gql__type_ActionByDatesResponse                      *graphql.Object      // message ActionByDatesResponse in api_action.proto
	gql__type_ActionByDatesRequest                       *graphql.Object      // message ActionByDatesRequest in api_action.proto
	gql__type_ActionByAddressResponse                    *graphql.Object      // message ActionByAddressResponse in api_action.proto
	gql__type_ActionByAddressRequest                     *graphql.Object      // message ActionByAddressRequest in api_action.proto
	gql__input_XrcInfo                                   *graphql.InputObject // message XrcInfo in api_action.proto
	gql__input_EvmTransfersByAddressResponse_EvmTransfer *graphql.InputObject // message EvmTransfersByAddressResponse.EvmTransfer in api_action.proto
	gql__input_EvmTransfersByAddressResponse             *graphql.InputObject // message EvmTransfersByAddressResponse in api_action.proto
	gql__input_EvmTransfersByAddressRequest              *graphql.InputObject // message EvmTransfersByAddressRequest in api_action.proto
	gql__input_EvmTransferInfo                           *graphql.InputObject // message EvmTransferInfo in api_action.proto
	gql__input_ActionResponse                            *graphql.InputObject // message ActionResponse in api_action.proto
	gql__input_ActionRequest                             *graphql.InputObject // message ActionRequest in api_action.proto
	gql__input_ActionInfo                                *graphql.InputObject // message ActionInfo in api_action.proto
	gql__input_ActionByTypeResponse                      *graphql.InputObject // message ActionByTypeResponse in api_action.proto
	gql__input_ActionByTypeRequest                       *graphql.InputObject // message ActionByTypeRequest in api_action.proto
	gql__input_ActionByHashResponse_EvmTransfers         *graphql.InputObject // message ActionByHashResponse.EvmTransfers in api_action.proto
	gql__input_ActionByHashResponse                      *graphql.InputObject // message ActionByHashResponse in api_action.proto
	gql__input_ActionByHashRequest                       *graphql.InputObject // message ActionByHashRequest in api_action.proto
	gql__input_ActionByDatesResponse                     *graphql.InputObject // message ActionByDatesResponse in api_action.proto
	gql__input_ActionByDatesRequest                      *graphql.InputObject // message ActionByDatesRequest in api_action.proto
	gql__input_ActionByAddressResponse                   *graphql.InputObject // message ActionByAddressResponse in api_action.proto
	gql__input_ActionByAddressRequest                    *graphql.InputObject // message ActionByAddressRequest in api_action.proto
)

func Gql__type_XrcInfo() *graphql.Object {
	if gql__type_XrcInfo == nil {
		gql__type_XrcInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_XrcInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"from": &graphql.Field{
					Type: graphql.String,
				},
				"to": &graphql.Field{
					Type: graphql.String,
				},
				"quantity": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
				"contract": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_XrcInfo
}

func Gql__type_EvmTransfersByAddressResponse_EvmTransfer() *graphql.Object {
	if gql__type_EvmTransfersByAddressResponse_EvmTransfer == nil {
		gql__type_EvmTransfersByAddressResponse_EvmTransfer = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EvmTransfersByAddressResponse_EvmTransfer",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"blkHash": &graphql.Field{
					Type: graphql.String,
				},
				"sender": &graphql.Field{
					Type: graphql.String,
				},
				"recipient": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_EvmTransfersByAddressResponse_EvmTransfer
}

func Gql__type_EvmTransfersByAddressResponse() *graphql.Object {
	if gql__type_EvmTransfersByAddressResponse == nil {
		gql__type_EvmTransfersByAddressResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EvmTransfersByAddressResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
				"evmTransfers": &graphql.Field{
					Type: graphql.NewList(Gql__type_EvmTransfersByAddressResponse_EvmTransfer()),
				},
			},
		})
	}
	return gql__type_EvmTransfersByAddressResponse
}

func Gql__type_EvmTransfersByAddressRequest() *graphql.Object {
	if gql__type_EvmTransfersByAddressRequest == nil {
		gql__type_EvmTransfersByAddressRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EvmTransfersByAddressRequest",
			Fields: graphql.Fields{
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_EvmTransfersByAddressRequest
}

func Gql__type_EvmTransferInfo() *graphql.Object {
	if gql__type_EvmTransferInfo == nil {
		gql__type_EvmTransferInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_EvmTransferInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"blkHash": &graphql.Field{
					Type: graphql.String,
				},
				"from": &graphql.Field{
					Type: graphql.String,
				},
				"to": &graphql.Field{
					Type: graphql.String,
				},
				"quantity": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_EvmTransferInfo
}

func Gql__type_ActionResponse() *graphql.Object {
	if gql__type_ActionResponse == nil {
		gql__type_ActionResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
				"actionList": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionInfo()),
				},
				"evmTransferList": &graphql.Field{
					Type: graphql.NewList(Gql__type_EvmTransferInfo()),
				},
				"xrcList": &graphql.Field{
					Type: graphql.NewList(Gql__type_XrcInfo()),
				},
			},
		})
	}
	return gql__type_ActionResponse
}

func Gql__type_ActionRequest() *graphql.Object {
	if gql__type_ActionRequest == nil {
		gql__type_ActionRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionRequest",
			Fields: graphql.Fields{
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_ActionRequest
}

func Gql__type_ActionInfo() *graphql.Object {
	if gql__type_ActionInfo == nil {
		gql__type_ActionInfo = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionInfo",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
				"blkHash": &graphql.Field{
					Type: graphql.String,
				},
				"actType": &graphql.Field{
					Type: graphql.String,
				},
				"sender": &graphql.Field{
					Type: graphql.String,
				},
				"recipient": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
				"timestamp": &graphql.Field{
					Type: graphql.Int,
				},
				"gasFee": &graphql.Field{
					Type: graphql.String,
				},
				"blkHeight": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ActionInfo
}

func Gql__type_ActionByTypeResponse() *graphql.Object {
	if gql__type_ActionByTypeResponse == nil {
		gql__type_ActionByTypeResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByTypeResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"actions": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionInfo()),
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ActionByTypeResponse
}

func Gql__type_ActionByTypeRequest() *graphql.Object {
	if gql__type_ActionByTypeRequest == nil {
		gql__type_ActionByTypeRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByTypeRequest",
			Fields: graphql.Fields{
				"type": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_ActionByTypeRequest
}

func Gql__type_ActionByHashResponse_EvmTransfers() *graphql.Object {
	if gql__type_ActionByHashResponse_EvmTransfers == nil {
		gql__type_ActionByHashResponse_EvmTransfers = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByHashResponse_EvmTransfers",
			Fields: graphql.Fields{
				"sender": &graphql.Field{
					Type: graphql.String,
				},
				"recipient": &graphql.Field{
					Type: graphql.String,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_ActionByHashResponse_EvmTransfers
}

func Gql__type_ActionByHashResponse() *graphql.Object {
	if gql__type_ActionByHashResponse == nil {
		gql__type_ActionByHashResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByHashResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"actionInfo": &graphql.Field{
					Type: Gql__type_ActionInfo(),
				},
				"evmTransfers": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionByHashResponse_EvmTransfers()),
				},
			},
		})
	}
	return gql__type_ActionByHashResponse
}

func Gql__type_ActionByHashRequest() *graphql.Object {
	if gql__type_ActionByHashRequest == nil {
		gql__type_ActionByHashRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByHashRequest",
			Fields: graphql.Fields{
				"actHash": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_ActionByHashRequest
}

func Gql__type_ActionByDatesResponse() *graphql.Object {
	if gql__type_ActionByDatesResponse == nil {
		gql__type_ActionByDatesResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByDatesResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"actions": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionInfo()),
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ActionByDatesResponse
}

func Gql__type_ActionByDatesRequest() *graphql.Object {
	if gql__type_ActionByDatesRequest == nil {
		gql__type_ActionByDatesRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByDatesRequest",
			Fields: graphql.Fields{
				"startDate": &graphql.Field{
					Type: graphql.Int,
				},
				"endDate": &graphql.Field{
					Type: graphql.Int,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_ActionByDatesRequest
}

func Gql__type_ActionByAddressResponse() *graphql.Object {
	if gql__type_ActionByAddressResponse == nil {
		gql__type_ActionByAddressResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByAddressResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"actions": &graphql.Field{
					Type: graphql.NewList(Gql__type_ActionInfo()),
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ActionByAddressResponse
}

func Gql__type_ActionByAddressRequest() *graphql.Object {
	if gql__type_ActionByAddressRequest == nil {
		gql__type_ActionByAddressRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ActionByAddressRequest",
			Fields: graphql.Fields{
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"pagination": &graphql.Field{
					Type: pagination.Gql__type_Pagination(),
				},
			},
		})
	}
	return gql__type_ActionByAddressRequest
}

func Gql__input_XrcInfo() *graphql.InputObject {
	if gql__input_XrcInfo == nil {
		gql__input_XrcInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_XrcInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"from": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"to": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"contract": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_XrcInfo
}

func Gql__input_EvmTransfersByAddressResponse_EvmTransfer() *graphql.InputObject {
	if gql__input_EvmTransfersByAddressResponse_EvmTransfer == nil {
		gql__input_EvmTransfersByAddressResponse_EvmTransfer = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EvmTransfersByAddressResponse_EvmTransfer",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"sender": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"recipient": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_EvmTransfersByAddressResponse_EvmTransfer
}

func Gql__input_EvmTransfersByAddressResponse() *graphql.InputObject {
	if gql__input_EvmTransfersByAddressResponse == nil {
		gql__input_EvmTransfersByAddressResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EvmTransfersByAddressResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"evmTransfers": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_EvmTransfersByAddressResponse_EvmTransfer()),
				},
			},
		})
	}
	return gql__input_EvmTransfersByAddressResponse
}

func Gql__input_EvmTransfersByAddressRequest() *graphql.InputObject {
	if gql__input_EvmTransfersByAddressRequest == nil {
		gql__input_EvmTransfersByAddressRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EvmTransfersByAddressRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"address": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_EvmTransfersByAddressRequest
}

func Gql__input_EvmTransferInfo() *graphql.InputObject {
	if gql__input_EvmTransferInfo == nil {
		gql__input_EvmTransferInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_EvmTransferInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"from": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"to": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"quantity": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_EvmTransferInfo
}

func Gql__input_ActionResponse() *graphql.InputObject {
	if gql__input_ActionResponse == nil {
		gql__input_ActionResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"actionList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionInfo()),
				},
				"evmTransferList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_EvmTransferInfo()),
				},
				"xrcList": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_XrcInfo()),
				},
			},
		})
	}
	return gql__input_ActionResponse
}

func Gql__input_ActionRequest() *graphql.InputObject {
	if gql__input_ActionRequest == nil {
		gql__input_ActionRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"address": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_ActionRequest
}

func Gql__input_ActionInfo() *graphql.InputObject {
	if gql__input_ActionInfo == nil {
		gql__input_ActionInfo = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionInfo",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"actType": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"sender": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"recipient": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"timestamp": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"gasFee": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blkHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ActionInfo
}

func Gql__input_ActionByTypeResponse() *graphql.InputObject {
	if gql__input_ActionByTypeResponse == nil {
		gql__input_ActionByTypeResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByTypeResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"actions": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionInfo()),
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ActionByTypeResponse
}

func Gql__input_ActionByTypeRequest() *graphql.InputObject {
	if gql__input_ActionByTypeRequest == nil {
		gql__input_ActionByTypeRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByTypeRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"type": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_ActionByTypeRequest
}

func Gql__input_ActionByHashResponse_EvmTransfers() *graphql.InputObject {
	if gql__input_ActionByHashResponse_EvmTransfers == nil {
		gql__input_ActionByHashResponse_EvmTransfers = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByHashResponse_EvmTransfers",
			Fields: graphql.InputObjectConfigFieldMap{
				"sender": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"recipient": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_ActionByHashResponse_EvmTransfers
}

func Gql__input_ActionByHashResponse() *graphql.InputObject {
	if gql__input_ActionByHashResponse == nil {
		gql__input_ActionByHashResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByHashResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"actionInfo": &graphql.InputObjectFieldConfig{
					Type: Gql__input_ActionInfo(),
				},
				"evmTransfers": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionByHashResponse_EvmTransfers()),
				},
			},
		})
	}
	return gql__input_ActionByHashResponse
}

func Gql__input_ActionByHashRequest() *graphql.InputObject {
	if gql__input_ActionByHashRequest == nil {
		gql__input_ActionByHashRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByHashRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"actHash": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_ActionByHashRequest
}

func Gql__input_ActionByDatesResponse() *graphql.InputObject {
	if gql__input_ActionByDatesResponse == nil {
		gql__input_ActionByDatesResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByDatesResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"actions": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionInfo()),
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ActionByDatesResponse
}

func Gql__input_ActionByDatesRequest() *graphql.InputObject {
	if gql__input_ActionByDatesRequest == nil {
		gql__input_ActionByDatesRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByDatesRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startDate": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"endDate": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_ActionByDatesRequest
}

func Gql__input_ActionByAddressResponse() *graphql.InputObject {
	if gql__input_ActionByAddressResponse == nil {
		gql__input_ActionByAddressResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByAddressResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"actions": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ActionInfo()),
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ActionByAddressResponse
}

func Gql__input_ActionByAddressRequest() *graphql.InputObject {
	if gql__input_ActionByAddressRequest == nil {
		gql__input_ActionByAddressRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ActionByAddressRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"address": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.InputObjectFieldConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
		})
	}
	return gql__input_ActionByAddressRequest
}

// graphql__resolver_ActionService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_ActionService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_ActionService creates pointer of service struct
func new_graphql_resolver_ActionService(conn *grpc.ClientConn) *graphql__resolver_ActionService {
	return &graphql__resolver_ActionService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_ActionService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_ActionService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"ActionByVoter": &graphql.Field{
			Type: Gql__type_ActionResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ActionByVoter")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.ActionByVoter(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ActionByVoter")
				}
				return resp, nil
			},
		},
		"GetXrc20ByAddress": &graphql.Field{
			Type: Gql__type_ActionResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for GetXrc20ByAddress")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.GetXrc20ByAddress(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetXrc20ByAddress")
				}
				return resp, nil
			},
		},
		"ActionByDates": &graphql.Field{
			Type: Gql__type_ActionByDatesResponse(),
			Args: graphql.FieldConfigArgument{
				"startDate": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"endDate": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionByDatesRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ActionByDates")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.ActionByDates(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ActionByDates")
				}
				return resp, nil
			},
		},
		"ActionByHash": &graphql.Field{
			Type: Gql__type_ActionByHashResponse(),
			Args: graphql.FieldConfigArgument{
				"actHash": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionByHashRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ActionByHash")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.ActionByHash(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ActionByHash")
				}
				return resp, nil
			},
		},
		"ActionByAddress": &graphql.Field{
			Type: Gql__type_ActionByAddressResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionByAddressRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ActionByAddress")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.ActionByAddress(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ActionByAddress")
				}
				return resp, nil
			},
		},
		"ActionByType": &graphql.Field{
			Type: Gql__type_ActionByTypeResponse(),
			Args: graphql.FieldConfigArgument{
				"type": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ActionByTypeRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ActionByType")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.ActionByType(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ActionByType")
				}
				return resp, nil
			},
		},
		"EvmTransfersByAddress": &graphql.Field{
			Type: Gql__type_EvmTransfersByAddressResponse(),
			Args: graphql.FieldConfigArgument{
				"address": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"pagination": &graphql.ArgumentConfig{
					Type: pagination.Gql__input_Pagination(),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req EvmTransfersByAddressRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for EvmTransfersByAddress")
				}
				client := NewActionServiceClient(conn)
				resp, err := client.EvmTransfersByAddress(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC EvmTransfersByAddress")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_ActionService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterActionServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterActionServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterActionServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service ActionService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterActionServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_ActionService(conn))
}
