package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	pb "51HW/proto"

	"google.golang.org/grpc"
)

type weatherServer struct {
	pb.UnimplementedWeatherServiceServer
}

func (s *weatherServer) GetWeatherUpdates(req *pb.WeatherRequest, stream pb.WeatherService_GetWeatherUpdatesServer) error {
	location := req.GetLocation()
	for {
		update := &pb.WeatherUpdate{
			Location:    location,
			Description: randomWeatherDescription(),
			Temperature: 15.0 + rand.Float32()*(30.0-15.0),
			Humidity:    50.0 + rand.Float32()*(90.0-50.0),
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		if err := stream.Send(update); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

func randomWeatherDescription() string {
	descriptions := []string{"Sunny", "Cloudy", "Rainy", "Stormy", "Snowy"}
	return descriptions[rand.Intn(len(descriptions))]
}


func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &weatherServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
