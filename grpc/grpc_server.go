package grpc

import (
	"net"

	pb "github.com/bogatyr285/test_task/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPC struct {
	addr    string
	grpcSrv *grpc.Server
	srvCtrl pb.FilesAPIServer
}

func NewGRPCServer(
	addr string,
	srvCtrl pb.FilesAPIServer,
) (*GRPC, error) {
	server := &GRPC{
		addr:    addr,
		srvCtrl: srvCtrl,
	}

	return server, nil
}

func (s *GRPC) Start() error {
	netListener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}

	s.grpcSrv = grpc.NewServer(opts...)

	pb.RegisterFilesAPIServer(s.grpcSrv, s.srvCtrl)
	// Register reflection service on gRPC server.
	reflection.Register(s.grpcSrv)

	go func() {
		log.Info().Msgf("GRPC server listening: %s", s.addr)
		log.Info().Msgf("GRPC server terminated with: %v", s.grpcSrv.Serve(netListener))
	}()

	return nil
}

// Stop - gracefully stop server & listeners
func (s *GRPC) Stop() {
	s.grpcSrv.GracefulStop()
	log.Info().Msg("GRPC gracefully stopped")
}
