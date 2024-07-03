package main

import (
	"GoBank/api"
	db "GoBank/db/sqlc"
	"GoBank/gapi"
	"GoBank/pb"
	"GoBank/util"
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // 数据库
	_ "github.com/golang-migrate/migrate/v4/source/file"       // 来自本地文件的迁移
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"

	// 通过导入改包来运行其init函数，
	_ "GoBank/doc/statik"
)

func main() {
	// 读取配置文件中的内容
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config:")
	}

	// 美化log输出，不直接显示json格式     生产环境中打开
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	runDBMigration(config.MigrationURL, config.DbSource)

	store := db.NewStore(conn)
	//runGinServer(config, store)
	go runGatewayServer(config, store)
	runGRPCServer(config, store)
}

func runGRPCServer(config util.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}
	// grpc拦截器 log日志捕获和输出
	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterGoBankServer(grpcServer, server)
	//log.Println(grpcServer.GetServiceInfo())

	// 注册自我反射
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener:")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server:")
	}

}

func runGatewayServer(config util.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}
	// 使用蛇形命名
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions:   protojson.MarshalOptions{UseProtoNames: true},
		UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	// 生成一个上下文和一个取消它的函数
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterGoBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("无法注册服务器:")
	}

	// 创建http多路复用器
	mux := http.NewServeMux()

	mux.Handle("/", grpcMux)

	// 添加Swagger     如果添加了命名空间，需要使用NewWithNamespace()
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("无法解析静态文件:")
	}
	// 新建一个HTTP文件服务器运行swagger
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(staticFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create listener:")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start HTTP gateway server:")
	}

}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("无法创建数据迁移：")
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("数据迁移运行失败：")
	}
	log.Info().Msg("数据迁移完成")
}
