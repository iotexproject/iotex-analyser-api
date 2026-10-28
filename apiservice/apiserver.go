package apiservice

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"math"
	"mime"
	"net"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
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
	api.RegisterVoterRewardServiceServer(grpcServer, &VoterRewardService{})
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
	if err := api.RegisterVoterRewardServiceHandlerServer(ctx, mux, &VoterRewardService{}); err != nil {
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

	graphqlHandler, err := newGraphQLHandler(ctx, templates)
	if err != nil {
		return err
	}

	http.Handle("/graphql", instrument("graphql", graphqlHandler))
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

func newGraphQLHandler(ctx context.Context, templates fs.FS) (http.Handler, error) {
	playground, err := template.ParseFS(templates, "templates/graphql-playground.html")
	if err != nil {
		return nil, err
	}
	graphqlMux := graphqlruntime.NewServeMux(auth.JWTTokenValid())
	if err := registerGraphQLAPIService(ctx, graphqlMux); err != nil {
		return nil, err
	}
	// The Playground is served outside graphqlMux because the mux always writes
	// its own result after its middlewares.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Add("Vary", "Accept")
			if wantsPlayground(r) {
				playground.Execute(w, nil)
				return
			}
		}
		graphqlMux.ServeHTTP(w, r)
	}), nil
}

// wantsPlayground reports whether a GET gets the Playground page rather than a
// GraphQL result. Browsers rank text/html above application/json, also when
// opening a Playground link with a pre-filled ?query=. Clients without a
// preference get the page only if they send no query and accept HTML.
func wantsPlayground(r *http.Request) bool {
	accept := strings.Join(r.Header.Values("Accept"), ",")
	html, json := acceptWeight(accept, "text/html"), acceptWeight(accept, "application/json")
	if html != json {
		return html > json
	}
	return html > 0 && r.URL.Query().Get("query") == ""
}

// acceptWeight returns the weight accept gives mediaType, taken from the most
// specific matching range (RFC 9110, section 12.5.1). Ranges with an invalid
// weight are ignored, and an empty header accepts everything.
func acceptWeight(accept, mediaType string) float64 {
	if strings.TrimSpace(accept) == "" {
		return 1
	}
	typ, _, _ := strings.Cut(mediaType, "/")
	weight, specificity := 0.0, 0
	for _, mediaRange := range strings.Split(accept, ",") {
		name, params, err := mime.ParseMediaType(mediaRange)
		if err != nil {
			continue
		}
		var s int
		switch name {
		case mediaType:
			s = 3
		case typ + "/*":
			s = 2
		case "*/*":
			s = 1
		}
		if s <= specificity {
			continue
		}
		q := 1.0
		if v, ok := params["q"]; ok {
			if q, err = strconv.ParseFloat(v, 64); err != nil || math.IsNaN(q) || q < 0 || q > 1 {
				continue
			}
		}
		weight, specificity = q, s
	}
	return weight
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
