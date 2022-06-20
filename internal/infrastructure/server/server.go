package server

import (
	"backend2/internal/usecases/app/repos/envrepo"
	"backend2/internal/usecases/app/repos/userrepo"
	"context"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
	un  *userrepo.Users
	en  *envrepo.Envs
	ern *envrepo.EnvsRel
}

func NewServer(addr string, h http.Handler) *Server {
	s := &Server{}

	s.srv = http.Server{
		Addr:              addr,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	err := s.srv.Shutdown(ctx)
	cancel()
	if err != nil {
		panic(err)
	}
}

func (s *Server) Start(un *userrepo.Users, en *envrepo.Envs, ern *envrepo.EnvsRel) {
	s.un = un
	s.en = en
	s.ern = ern

	go func() {
		_ = s.srv.ListenAndServe()
	}()
}
