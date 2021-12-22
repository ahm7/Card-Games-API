package DataBase

import (
	"main/Model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
     "main/ApiHelper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert_Deck(Deck Model.Deck) {

	client, ctx, cancel, err := connect(mongoUrl)
	_ = client
	_ = ctx
	_ = cancel
	_ = err

	usersCollection := client.Database("API").Collection("Cards_Api")
	result, err := usersCollection.InsertOne(context.TODO(), Deck)
	_ = result
	_ = err

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
	if err != nil {
		panic(err)
	}
}

func Get_Deck(w http.ResponseWriter, id string) []byte {

	client, ctx, cancel, err := connect(mongoUrl)
	_ = client
	_ = ctx
	_ = cancel
	_ = err

	usersCollection := client.Database("API").Collection("Cards_Api")

	var result bson.M
	err = usersCollection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", id)
		ApiHelper.ContentNotFOund(w)
		return []byte(err.Error())
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	return jsonData

}
func Delete_cards(w http.ResponseWriter, count int, id string) []Model.Card {

	client, ctx, cancel, err := connect(mongoUrl)
	_ = client
	_ = ctx
	_ = cancel
	_ = err

	var deck Model.Deck
	json.Unmarshal([]byte(Get_Deck(w, id)), &deck) // get the dick with this id
	arr := deck.Cards[0:count]
	usersCollection := client.Database("API").Collection("Cards_Api")
	filter := bson.M{"id": id}
	statement := bson.M{"$pull": bson.M{"cards": bson.M{"$in": arr}}}
	result, err := usersCollection.UpdateMany(ctx, filter, statement)

	result2, err2 := usersCollection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"remaining", deck.Remaining - count}}},
		},
	)

	_ = result2
	_ = err2
	_= result
	return arr

}
