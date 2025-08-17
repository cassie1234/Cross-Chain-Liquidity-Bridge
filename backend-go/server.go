package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 7019
// Hash 8283
// Hash 3211
// Hash 2837
// Hash 7217
// Hash 3197
// Hash 7269
// Hash 8245
// Hash 5639
// Hash 1668
// Hash 6055
// Hash 2019
// Hash 9815
// Hash 4047
// Hash 9180
// Hash 8313
// Hash 5359
// Hash 6520
// Hash 8142
// Hash 6278
// Hash 9439
// Hash 8335
// Hash 6272
// Hash 9194
// Hash 2726
// Hash 5903
// Hash 9490
// Hash 7696
// Hash 2668
// Hash 6407
// Hash 7488
// Hash 2519
// Hash 2307
// Hash 2184
// Hash 5321
// Hash 9033
// Hash 4723
// Hash 6882
// Hash 1198
// Hash 2297
// Hash 3024
// Hash 1465
// Hash 7223
// Hash 6721
// Hash 1869
// Hash 9472
// Hash 5552
// Hash 9037
// Hash 2701
// Hash 8200
// Hash 2768
// Hash 7325
// Hash 8507
// Hash 9840
// Hash 1123
// Hash 5400
// Hash 4627
// Hash 7123
// Hash 1948
// Hash 6255
// Hash 8433
// Hash 2520
// Hash 6228
// Hash 9278
// Hash 2278
// Hash 3812
// Hash 5260
// Hash 7113
// Hash 4234
// Hash 1456
// Hash 4434
// Hash 5379
// Hash 3923
// Hash 9923
// Hash 4366
// Hash 7160
// Hash 2160
// Hash 7285
// Hash 9458
// Hash 3303
// Hash 6990
// Hash 2729
// Hash 2522
// Hash 8263
// Hash 1297
// Hash 6449
// Hash 1404
// Hash 1817
// Hash 9625
// Hash 4628