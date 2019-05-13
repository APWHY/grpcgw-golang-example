package main

// adapted from https://github.com/grpc/grpc-go/blob/master/examples/route_guide/client/client.go
import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"time"

	"github.com/APWHY/grpcgw-golang-example/proto"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:8081", "The server address in the format of host:port")
)

func sendAPet(client proto.PetServiceClient, pet *proto.Pet) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.PostPet(ctx, &proto.PostPetRequest{Pet: pet})
	if err != nil {
		log.Fatalf("%v.PostPet(_) = _, %v", client, err)
	}

	log.Print("received response after posting pet: ", res)
}

func getPets(client proto.PetServiceClient) {
	ctx := context.Background()

	res, err := client.GetPets(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("%v.GetPets(_) = _, %v", client, err)
	}

	log.Print("received response after getting all pets: ", res)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewPetServiceClient(conn)

	in := &proto.Pet{
		Name:    "Foghorn",
		Owner:   "Leo",
		Species: "bullfrog",
		Sex:     "m",
		Birth:   "2012-07-03",
	}

	sendAPet(client, in)
	in = &proto.Pet{
		Name:    "Cherry",
		Owner:   "Dan",
		Species: "guinea pig",
		Sex:     "f",
		Birth:   "2008-11-01",
	}

	sendAPet(client, in)
	getPets(client)
}
