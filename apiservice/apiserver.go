package apiservice

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"text/template"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/auth"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	graphqlruntime "github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// DocsHTML embed the docs HTML
var DocsHTML embed.FS
var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

const (
	// MaxRecvMsgSize is the max recv size
	MaxRecvMsgSize = 1024 * 1024 * 40 // 40 MB
	// MaxSendMsgSize is the max send size
	MaxSendMsgSize = 1024 * 1024 * 40 // 40 MB
)

func registerAPIService(ctx context.Context, grpcServer *grpc.Server) {
	api.RegisterAccountServiceServer(grpcServer, &AccountService{})
	api.RegisterStakingServiceServer(grpcServer, &StakingService{})
	api.RegisterActionsServiceServer(grpcServer, &ActionsService{})
	api.RegisterDelegateServiceServer(grpcServer, &DelegateService{})
	api.RegisterChainServiceServer(grpcServer, &ChainService{})
	api.RegisterActionServiceServer(grpcServer, &ActionService{})
	api.RegisterVotingServiceServer(grpcServer, &VotingService{})
	api.RegisterXRC20ServiceServer(grpcServer, &XRC20Service{})
	api.RegisterXRC721ServiceServer(grpcServer, &XRC721Service{})
	api.RegisterHermesServiceServer(grpcServer, &HermesService{})
	api.RegisterStreamServiceServer(grpcServer, &StreamService{})
	api.RegisterApprovalServiceServer(grpcServer, &ApprovalService{})
	api.RegisterExitQueueServiceServer(grpcServer, &ExitQueueService{})
	api.RegisterIotexscanServiceServer(grpcServer, &IotexscanService{})
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
	if err := api.RegisterVotingServiceHandlerServer(ctx, mux, &VotingService{}); err != nil {
		return err
	}
	if err := api.RegisterXRC20ServiceHandlerServer(ctx, mux, &XRC20Service{}); err != nil {
		return err
	}
	if err := api.RegisterXRC721ServiceHandlerServer(ctx, mux, &XRC721Service{}); err != nil {
		return err
	}
	if err := api.RegisterHermesServiceHandlerServer(ctx, mux, &HermesService{}); err != nil {
		return err
	}
	if err := api.RegisterApprovalServiceHandlerServer(ctx, mux, &ApprovalService{}); err != nil {
		return err
	}
	if err := api.RegisterExitQueueServiceHandlerServer(ctx, mux, &ExitQueueService{}); err != nil {
		return err
	}
	if err := api.RegisterIotexscanServiceHandlerServer(ctx, mux, &IotexscanService{}); err != nil {
		return err
	}
	return nil
}

func registerGraphQLAPIService(ctx context.Context, mux *graphqlruntime.ServeMux) error {
	addr := fmt.Sprintf("127.0.0.1:%d", config.Default.Server.GrpcAPIPort)
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxRecvMsgSize)))
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
	if err := api.RegisterVotingServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterXRC20ServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterXRC721ServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterHermesServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterApprovalServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	if err := api.RegisterExitQueueServiceGraphqlHandler(mux, conn); err != nil {
		return err
	}
	return nil
}

// StartGRPCService starts the GRPC service
func StartGRPCService(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Default.Server.GrpcAPIPort))
	if err != nil {
		return err
	}
	customFunc = func(p interface{}) (err error) {
		log.Println("panic :", p, string(debug.Stack()))
		return status.Errorf(codes.InvalidArgument, "Panic triggered")
	}
	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(MaxRecvMsgSize),
		grpc.MaxSendMsgSize(MaxSendMsgSize),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc)),
		)),
	}

	grpcServer := grpc.NewServer(options...)
	//serviceName: grpc.health.v1.Health
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	registerAPIService(ctx, grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}

// StartGRPCProxyService starts the GRPC API Proxy service
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
	graphqlMux := graphqlruntime.NewServeMux(playgroundMiddleware, auth.JWTTokenValid())
	if err := registerGraphQLAPIService(ctx, graphqlMux); err != nil {
		return err
	}

	http.Handle("/graphql", instrument("graphql", graphqlMux))
	fsys, err := fs.Sub(DocsHTML, "docs-html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/docs/", instrument("docs", http.StripPrefix("/docs/", http.FileServer(http.FS(fsys)))))
	http.HandleFunc("/healthz", healthzHandler)
	http.Handle("/", instrument("api", auth.JWTTokenMiddleware(auth.CheckWhiteListMiddleware(gwmux))))
	http.Handle("/metrics", promhttp.Handler())

	port := fmt.Sprintf(":%d", config.Default.Server.HTTPAPIPort)
	return http.ListenAndServe(port, nil)
}

// healthzHandler is an unauthenticated LB probe: pings the DB and returns 200/503.
// No Prometheus instrumentation so high-frequency probes don't pollute /metrics.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	gdb := db.DB()
	if gdb == nil {
		http.Error(w, "db not initialized", http.StatusServiceUnavailable)
		return
	}
	sqlDB, err := gdb.DB()
	if err != nil {
		http.Error(w, "db unavailable", http.StatusServiceUnavailable)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		http.Error(w, "db ping failed", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
