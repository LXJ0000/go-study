package main

import (
	"stream_server/pb"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{"Hello!", "Howdy!", "Hi there!", "Greetings!"}

	for _, word := range words {
		data := &pb.HelloResponse{Reply: word + in.GetName()}
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

func main() {

}
