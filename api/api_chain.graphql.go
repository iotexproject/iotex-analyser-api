// Code generated by proroc-gen-graphql, DO NOT EDIT.
package api

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_VotingResultMeta                *graphql.Object      // message VotingResultMeta in api_chain.proto
	gql__type_TotalTransferredTokensResponse  *graphql.Object      // message TotalTransferredTokensResponse in api_chain.proto
	gql__type_TotalTransferredTokensRequest   *graphql.Object      // message TotalTransferredTokensRequest in api_chain.proto
	gql__type_NumberOfActionsResponse         *graphql.Object      // message NumberOfActionsResponse in api_chain.proto
	gql__type_NumberOfActionsRequest          *graphql.Object      // message NumberOfActionsRequest in api_chain.proto
	gql__type_MostRecentTPSResponse           *graphql.Object      // message MostRecentTPSResponse in api_chain.proto
	gql__type_MostRecentTPSRequest            *graphql.Object      // message MostRecentTPSRequest in api_chain.proto
	gql__type_ChartSyncResponse_State         *graphql.Object      // message ChartSyncResponse.State in api_chain.proto
	gql__type_ChartSyncResponse               *graphql.Object      // message ChartSyncResponse in api_chain.proto
	gql__type_ChartSyncRequest                *graphql.Object      // message ChartSyncRequest in api_chain.proto
	gql__type_ChainResponse_Rewards           *graphql.Object      // message ChainResponse.Rewards in api_chain.proto
	gql__type_ChainResponse                   *graphql.Object      // message ChainResponse in api_chain.proto
	gql__input_VotingResultMeta               *graphql.InputObject // message VotingResultMeta in api_chain.proto
	gql__input_TotalTransferredTokensResponse *graphql.InputObject // message TotalTransferredTokensResponse in api_chain.proto
	gql__input_TotalTransferredTokensRequest  *graphql.InputObject // message TotalTransferredTokensRequest in api_chain.proto
	gql__input_NumberOfActionsResponse        *graphql.InputObject // message NumberOfActionsResponse in api_chain.proto
	gql__input_NumberOfActionsRequest         *graphql.InputObject // message NumberOfActionsRequest in api_chain.proto
	gql__input_MostRecentTPSResponse          *graphql.InputObject // message MostRecentTPSResponse in api_chain.proto
	gql__input_MostRecentTPSRequest           *graphql.InputObject // message MostRecentTPSRequest in api_chain.proto
	gql__input_ChartSyncResponse_State        *graphql.InputObject // message ChartSyncResponse.State in api_chain.proto
	gql__input_ChartSyncResponse              *graphql.InputObject // message ChartSyncResponse in api_chain.proto
	gql__input_ChartSyncRequest               *graphql.InputObject // message ChartSyncRequest in api_chain.proto
	gql__input_ChainResponse_Rewards          *graphql.InputObject // message ChainResponse.Rewards in api_chain.proto
	gql__input_ChainResponse                  *graphql.InputObject // message ChainResponse in api_chain.proto
)

func Gql__type_VotingResultMeta() *graphql.Object {
	if gql__type_VotingResultMeta == nil {
		gql__type_VotingResultMeta = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_VotingResultMeta",
			Fields: graphql.Fields{
				"totalCandidates": &graphql.Field{
					Type: graphql.Int,
				},
				"totalWeightedVotes": &graphql.Field{
					Type: graphql.String,
				},
				"votedTokens": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_VotingResultMeta
}

func Gql__type_TotalTransferredTokensResponse() *graphql.Object {
	if gql__type_TotalTransferredTokensResponse == nil {
		gql__type_TotalTransferredTokensResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_TotalTransferredTokensResponse",
			Fields: graphql.Fields{
				"totalTransferredTokens": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_TotalTransferredTokensResponse
}

func Gql__type_TotalTransferredTokensRequest() *graphql.Object {
	if gql__type_TotalTransferredTokensRequest == nil {
		gql__type_TotalTransferredTokensRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_TotalTransferredTokensRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_TotalTransferredTokensRequest
}

func Gql__type_NumberOfActionsResponse() *graphql.Object {
	if gql__type_NumberOfActionsResponse == nil {
		gql__type_NumberOfActionsResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_NumberOfActionsResponse",
			Fields: graphql.Fields{
				"exist": &graphql.Field{
					Type: graphql.Boolean,
				},
				"count": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_NumberOfActionsResponse
}

func Gql__type_NumberOfActionsRequest() *graphql.Object {
	if gql__type_NumberOfActionsRequest == nil {
		gql__type_NumberOfActionsRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_NumberOfActionsRequest",
			Fields: graphql.Fields{
				"startEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"epochCount": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_NumberOfActionsRequest
}

func Gql__type_MostRecentTPSResponse() *graphql.Object {
	if gql__type_MostRecentTPSResponse == nil {
		gql__type_MostRecentTPSResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_MostRecentTPSResponse",
			Fields: graphql.Fields{
				"mostRecentTPS": &graphql.Field{
					Type: graphql.Float,
				},
			},
		})
	}
	return gql__type_MostRecentTPSResponse
}

func Gql__type_MostRecentTPSRequest() *graphql.Object {
	if gql__type_MostRecentTPSRequest == nil {
		gql__type_MostRecentTPSRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_MostRecentTPSRequest",
			Fields: graphql.Fields{
				"blockWindow": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_MostRecentTPSRequest
}

func Gql__type_ChartSyncResponse_State() *graphql.Object {
	if gql__type_ChartSyncResponse_State == nil {
		gql__type_ChartSyncResponse_State = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ChartSyncResponse_State",
			Fields: graphql.Fields{
				"time": &graphql.Field{
					Type: graphql.String,
				},
				"size": &graphql.Field{
					Type: graphql.String,
				},
				"serverVersion": &graphql.Field{
					Type: graphql.String,
				},
				"blockNumber": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__type_ChartSyncResponse_State
}

func Gql__type_ChartSyncResponse() *graphql.Object {
	if gql__type_ChartSyncResponse == nil {
		gql__type_ChartSyncResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ChartSyncResponse",
			Fields: graphql.Fields{
				"states": &graphql.Field{
					Type: graphql.NewList(Gql__type_ChartSyncResponse_State()),
				},
			},
		})
	}
	return gql__type_ChartSyncResponse
}

func Gql__type_ChartSyncRequest() *graphql.Object {
	if gql__type_ChartSyncRequest == nil {
		gql__type_ChartSyncRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ChartSyncRequest",
			Fields: graphql.Fields{
				"archive": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_ChartSyncRequest
}

func Gql__type_ChainResponse_Rewards() *graphql.Object {
	if gql__type_ChainResponse_Rewards == nil {
		gql__type_ChainResponse_Rewards = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ChainResponse_Rewards",
			Fields: graphql.Fields{
				"totalBalance": &graphql.Field{
					Type: graphql.String,
				},
				"totalUnclaimed": &graphql.Field{
					Type: graphql.String,
				},
				"totalAvailable": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_ChainResponse_Rewards
}

func Gql__type_ChainResponse() *graphql.Object {
	if gql__type_ChainResponse == nil {
		gql__type_ChainResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Api_Type_ChainResponse",
			Fields: graphql.Fields{
				"mostRecentEpoch": &graphql.Field{
					Type: graphql.Int,
				},
				"mostRecentBlockHeight": &graphql.Field{
					Type: graphql.Int,
				},
				"totalSupply": &graphql.Field{
					Type: graphql.String,
				},
				"totalCirculatingSupply": &graphql.Field{
					Type: graphql.String,
				},
				"totalCirculatingSupplyNoRewardPool": &graphql.Field{
					Type: graphql.String,
				},
				"votingResultMeta": &graphql.Field{
					Type: Gql__type_VotingResultMeta(),
				},
				"exactCirculatingSupply": &graphql.Field{
					Type: graphql.String,
				},
				"rewards": &graphql.Field{
					Type: Gql__type_ChainResponse_Rewards(),
				},
			},
		})
	}
	return gql__type_ChainResponse
}

func Gql__input_VotingResultMeta() *graphql.InputObject {
	if gql__input_VotingResultMeta == nil {
		gql__input_VotingResultMeta = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_VotingResultMeta",
			Fields: graphql.InputObjectConfigFieldMap{
				"totalCandidates": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"totalWeightedVotes": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"votedTokens": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_VotingResultMeta
}

func Gql__input_TotalTransferredTokensResponse() *graphql.InputObject {
	if gql__input_TotalTransferredTokensResponse == nil {
		gql__input_TotalTransferredTokensResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_TotalTransferredTokensResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"totalTransferredTokens": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_TotalTransferredTokensResponse
}

func Gql__input_TotalTransferredTokensRequest() *graphql.InputObject {
	if gql__input_TotalTransferredTokensRequest == nil {
		gql__input_TotalTransferredTokensRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_TotalTransferredTokensRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_TotalTransferredTokensRequest
}

func Gql__input_NumberOfActionsResponse() *graphql.InputObject {
	if gql__input_NumberOfActionsResponse == nil {
		gql__input_NumberOfActionsResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_NumberOfActionsResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"exist": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
				"count": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_NumberOfActionsResponse
}

func Gql__input_NumberOfActionsRequest() *graphql.InputObject {
	if gql__input_NumberOfActionsRequest == nil {
		gql__input_NumberOfActionsRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_NumberOfActionsRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"startEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_NumberOfActionsRequest
}

func Gql__input_MostRecentTPSResponse() *graphql.InputObject {
	if gql__input_MostRecentTPSResponse == nil {
		gql__input_MostRecentTPSResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_MostRecentTPSResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"mostRecentTPS": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
				},
			},
		})
	}
	return gql__input_MostRecentTPSResponse
}

func Gql__input_MostRecentTPSRequest() *graphql.InputObject {
	if gql__input_MostRecentTPSRequest == nil {
		gql__input_MostRecentTPSRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_MostRecentTPSRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"blockWindow": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_MostRecentTPSRequest
}

func Gql__input_ChartSyncResponse_State() *graphql.InputObject {
	if gql__input_ChartSyncResponse_State == nil {
		gql__input_ChartSyncResponse_State = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ChartSyncResponse_State",
			Fields: graphql.InputObjectConfigFieldMap{
				"time": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"size": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"serverVersion": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"blockNumber": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_ChartSyncResponse_State
}

func Gql__input_ChartSyncResponse() *graphql.InputObject {
	if gql__input_ChartSyncResponse == nil {
		gql__input_ChartSyncResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ChartSyncResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"states": &graphql.InputObjectFieldConfig{
					Type: graphql.NewList(Gql__input_ChartSyncResponse_State()),
				},
			},
		})
	}
	return gql__input_ChartSyncResponse
}

func Gql__input_ChartSyncRequest() *graphql.InputObject {
	if gql__input_ChartSyncRequest == nil {
		gql__input_ChartSyncRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ChartSyncRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"archive": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_ChartSyncRequest
}

func Gql__input_ChainResponse_Rewards() *graphql.InputObject {
	if gql__input_ChainResponse_Rewards == nil {
		gql__input_ChainResponse_Rewards = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ChainResponse_Rewards",
			Fields: graphql.InputObjectConfigFieldMap{
				"totalBalance": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"totalUnclaimed": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"totalAvailable": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_ChainResponse_Rewards
}

func Gql__input_ChainResponse() *graphql.InputObject {
	if gql__input_ChainResponse == nil {
		gql__input_ChainResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Api_Input_ChainResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"mostRecentEpoch": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"mostRecentBlockHeight": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"totalSupply": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"totalCirculatingSupply": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"totalCirculatingSupplyNoRewardPool": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"votingResultMeta": &graphql.InputObjectFieldConfig{
					Type: Gql__input_VotingResultMeta(),
				},
				"exactCirculatingSupply": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"rewards": &graphql.InputObjectFieldConfig{
					Type: Gql__input_ChainResponse_Rewards(),
				},
			},
		})
	}
	return gql__input_ChainResponse
}

// graphql__resolver_ChainService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_ChainService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_ChainService creates pointer of service struct
func new_graphql_resolver_ChainService(conn *grpc.ClientConn) *graphql__resolver_ChainService {
	return &graphql__resolver_ChainService{
		conn:        conn,
		host:        "localhost:50051",
		dialOptions: []grpc.DialOption{},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_ChainService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
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
func (x *graphql__resolver_ChainService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"Chain": &graphql.Field{
			Type: Gql__type_ChainResponse(),
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ChainRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for Chain")
				}
				client := NewChainServiceClient(conn)
				resp, err := client.Chain(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC Chain")
				}
				return resp, nil
			},
		},
		"MostRecentTPS": &graphql.Field{
			Type: Gql__type_MostRecentTPSResponse(),
			Args: graphql.FieldConfigArgument{
				"blockWindow": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req MostRecentTPSRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for MostRecentTPS")
				}
				client := NewChainServiceClient(conn)
				resp, err := client.MostRecentTPS(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC MostRecentTPS")
				}
				return resp, nil
			},
		},
		"NumberOfActions": &graphql.Field{
			Type: Gql__type_NumberOfActionsResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req NumberOfActionsRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for NumberOfActions")
				}
				client := NewChainServiceClient(conn)
				resp, err := client.NumberOfActions(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC NumberOfActions")
				}
				return resp, nil
			},
		},
		"TotalTransferredTokens": &graphql.Field{
			Type: Gql__type_TotalTransferredTokensResponse(),
			Args: graphql.FieldConfigArgument{
				"startEpoch": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"epochCount": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req TotalTransferredTokensRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for TotalTransferredTokens")
				}
				client := NewChainServiceClient(conn)
				resp, err := client.TotalTransferredTokens(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC TotalTransferredTokens")
				}
				return resp, nil
			},
		},
		"ChartSync": &graphql.Field{
			Type: Gql__type_ChartSyncResponse(),
			Args: graphql.FieldConfigArgument{
				"archive": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ChartSyncRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for ChartSync")
				}
				client := NewChainServiceClient(conn)
				resp, err := client.ChartSync(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ChartSync")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_ChainService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterChainServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterChainServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterChainServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
// service ChainService {
//    option (graphql.service) = {
//        host: "host:port"
//        insecure: true or false
//    };
//
//    ...with RPC definitions
// }
func RegisterChainServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_ChainService(conn))
}
