package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "51HW/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewWeatherServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    req := &pb.WeatherRequest{Location: "Tashkent"}
    stream, err := client.GetWeatherUpdates(ctx, req)
    if err != nil {
        log.Fatalf("could not get weather updates: %v", err)
    }

    for {
        update, err := stream.Recv()
        if err != nil {
            log.Fatalf("error while receiving: %v", err)
        }
        log.Printf("Weather Update: %s - %s, %.2f°C, %.2f%% humidity at %s", update.GetLocation(), update.GetDescription(), update.GetTemperature(), update.GetHumidity(), update.GetTimestamp())
    }
}
