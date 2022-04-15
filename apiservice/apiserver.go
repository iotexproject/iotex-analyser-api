package apiservice

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"text/template"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/config"
	graphqlruntime "github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var DocsHTML embed.FS

func registerAPIService(ctx context.Context, grpcServer *grpc.Server) {
	api.RegisterAccountServiceServer(grpcServer, &AccountService{})
	api.RegisterStakingServiceServer(grpcServer, &StakingService{})
	api.RegisterActionsServiceServer(grpcServer, &ActionsService{})
	api.RegisterDelegateServiceServer(grpcServer, &DelegateService{})
	api.RegisterChainServiceServer(grpcServer, &ChainService{})
	api.RegisterActionServiceServer(grpcServer, &ActionService{})
}

func registerProxyAPIService(ctx context.Context, mux *runtime.ServeMux) error {
	if err := api.RegisterAccountServiceHandlerServer(ctx, mux, &AccountService{}); err != nil {
		return err
	}
	if err := api.RegisterStakingServiceHandlerServer(ctx, mux, &StakingService{}); err != nil {
		return err
	}
	if err := api.RegisterActionsServiceHandlerServer(ctx, mux, &ActionsService{}); err != nil {
		return err
	}
	if err := api.RegisterDelegateServiceHandlerServer(ctx, mux, &DelegateService{}); err != nil {
		return err
	}
	if err := api.RegisterChainServiceHandlerServer(ctx, mux, &ChainService{}); err != nil {
		return err
	}
	if err := api.RegisterActionServiceHandlerServer(ctx, mux, &ActionService{}); err != nil {
		return err
	}
	return nil
}

func registerGraphQLAPIService(ctx context.Context, mux *graphqlruntime.ServeMux) error {
	addr := fmt.Sprintf("127.0.0.1:%d", config.Default.Server.GrpcAPIPort)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	if err := api.RegisterActionsServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterStakingServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterAccountServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterDelegateServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterChainServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterActionServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	return nil
}

func StartGRPCService(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Default.Server.GrpcAPIPort))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	registerAPIService(ctx, grpcServer)
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}

func StartGRPCProxyService(templates embed.FS) error {
	gwmux := runtime.NewServeMux()
	ctx := context.Background()
	if err := registerProxyAPIService(ctx, gwmux); err != nil {
		return err
	}

	playgroundMiddleware := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (context.Context, error) {
		if r.Method == "GET" {
			tmpl, err := template.ParseFS(templates, "templates/graphql-playground.html")
			if err != nil {
				return ctx, err
			}
			if err := tmpl.Execute(w, nil); err != nil {
				return ctx, err
			}
		}
		return ctx, nil
	}
	graphqlMux := graphqlruntime.NewServeMux(playgroundMiddleware)
	if err := registerGraphQLAPIService(ctx, graphqlMux); err != nil {
		return err
	}

	http.Handle("/graphql", graphqlMux)
	fsys, err := fs.Sub(DocsHTML, "docs-html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.FS(fsys))))
	http.Handle("/", gwmux)

	port := fmt.Sprintf(":%d", config.Default.Server.HTTPAPIPort)
	return http.ListenAndServe(port, nil)
}
