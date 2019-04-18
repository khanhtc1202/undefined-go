package main

import (
	"context"

	"github.com/khanhtc1202/undefined-go/protoc-play/app/server"

	"github.com/izumin5210/grapi/pkg/grapiserver"
)

func run() error {
	// Application context
	ctx := context.Background()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(),
		),
	)
	return s.ServeContext(ctx)
}
