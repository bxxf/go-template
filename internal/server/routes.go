package server

import (
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/bxxf/go-template/gen/api/auth/v1/authconnect"
)

func (s *Server) defineRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(authconnect.NewAuthServiceHandler(s.authRouter))

	// Add reflection for development
	if s.config.Env == "development" {

		reflector := grpcreflect.NewStaticReflector(
			"auth.v1.AuthService",
		)

		mux.Handle(grpcreflect.NewHandlerV1(reflector))
		mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
	}

	return mux
}
