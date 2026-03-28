package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	// Mock data
	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
		Phone: "999-999-9999",
	}, nil
}

func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {
	log.Printf("Received request for StudentList")

	return &pb.StudentListResponse{
		Student: []*pb.StudentResponse{
			{
				Id:    1,
				Name:  "Ali Baba",
				Major: "Computer Science",
				Email: "Ali@university.com",
				Phone: "123-112-1111",
			},
			{
				Id:    2,
				Name:  "Goo Gle",
				Major: "Computer Science",
				Email: "Goo@university.com",
				Phone: "321-321-2221",
			},
			{
				Id:    3,
				Name:  "Ama Zon",
				Major: "Computer Science",
				Email: "Ama@university.com",
				Phone: "444-444-3333",
			},
		},
	}, nil
}


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}