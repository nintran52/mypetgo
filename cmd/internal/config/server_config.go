package config

import "github.com/nintran52/mypetgo/cmd/internal/util"

type EchoServer struct {
	ListenAddress string
}

type Server struct {
	Echo EchoServer
}

func DefaultServerConfigFromEnv() Server {
	return Server{
		Echo: EchoServer{
			ListenAddress: util.GetEnv("SERVER_ECHO_LISTEN_ADDRESS", ":8080"),
		},
	}
}
