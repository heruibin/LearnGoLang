package MongoByDriver

import (
	"log"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
)

func main1() {
	//client, err := mongo.NewClient("mongodb://localhost:27017")
	opts := clientopt.AppName("test")
	client, err := mongo.NewClientWithOptions("mongodb://localhost:27017", opts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("music_user").Collection("abc")
	res, err := collection.InsertOne(context.Background(), map[string]string{"a": "a1", "b": "b1", "c": "c1"})

	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Println(id)
	client.Disconnect(context.Background())
}
