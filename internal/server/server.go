package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"github.com/skinnykaen/rpa_clone/graph"
	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/graphql/directives"
	resolvers "github.com/skinnykaen/rpa_clone/internal/transports/graphql"
	http2 "github.com/skinnykaen/rpa_clone/internal/transports/http"
	"github.com/skinnykaen/rpa_clone/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func NewServer(
	m consts.Mode,
	lifecycle fx.Lifecycle,
	loggers logger.Loggers,
	resolver resolvers.Resolver,
	handlers http2.Handlers,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) (err error) {
				serverHost := viper.GetString("server_host")
				port := viper.GetString("graphql_server_port")
				router := gin.Default()
				router.Use(
					gin.Recovery(),
					gin.Logger(),
					GinContextToContextMiddleware(),
					AuthMiddleware(loggers.Err),
				)

				c := graph.Config{Resolvers: &resolver}
				c.Directives.HasRole = directives.HasRole(loggers.Err)
				srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))
				switch m {
				case consts.Production:
					router.POST("/query", gin.WrapH(srv))
					handlers.AuthHandler.SetupAuthRoutes(router)
					handlers.ProjectHandler.SetupProjectRoutes(router)
					handlers.AvatarHandler.SetupAvatarRoutes(router)
					handlers.ApplicationHandler.SetupApplicationRoutes(router)
				case consts.Development:
					router.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
					router.POST("/query", gin.WrapH(srv))
					handlers.AuthHandler.SetupAuthRoutes(router)
					handlers.ProjectHandler.SetupProjectRoutes(router)
					handlers.AvatarHandler.SetupAvatarRoutes(router)
					handlers.ApplicationHandler.SetupApplicationRoutes(router)
				}

				server := &http.Server{
					Addr: serverHost + ":" + port,
					Handler: cors.New(
						cors.Options{
							AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
							AllowCredentials: viper.GetBool("cors.allow_credentials"),
							AllowedMethods:   viper.GetStringSlice("cors.allowed_methods"),
							AllowedHeaders:   viper.GetStringSlice("cors.allowed_headers"),
						},
					).Handler(router),
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1 << 20,
				}

				loggers.Info.Printf(
					"Connect to %s:%s/ for GraphQL playground",
					serverHost,
					port,
				)
				loggers.Info.Printf(
					"The app is running in %s mode",
					m,
				)
				go func() {
					if err = server.ListenAndServe(); err != nil {
						loggers.Err.Fatal("Failed to listen and serve: %v", err)
					}
				}()
				return
			},
			OnStop: func(context.Context) error {
				return nil
			},
		})
}
