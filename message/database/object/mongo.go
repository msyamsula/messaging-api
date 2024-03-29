package object

import (
	"context"
	"fmt"
	"time"

	"github.com/msyamsula/messaging-api/message/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func createCollection(p1 int64, p2 int64) (string, error) {

	var start, end int64
	if p1 == p2 {
		return "", database.ErrCreatingCollection
	}

	if p1 < p2 {
		start = p1
		end = p2
	} else {
		start = p2
		end = p1
	}

	return fmt.Sprintf("%vw%v", start, end), nil
}

func (mg *Mongo) InsertMessage(ctx context.Context, m database.MessageToInsert) error {

	var err error
	var cName string
	cName, err = createCollection(m.SenderID, m.ReceiverID)
	if err != nil {
		return database.ErrCreatingCollection
	}
	c := mg.Client.Database(database.DatabaseName).Collection(cName)

	_, err = c.InsertOne(ctx, m)
	return err
}

func (mg *Mongo) GetConversation(ctx context.Context, person1 int64, person2 int64) ([]database.Message, error) {

	var err error
	var res []database.Message
	threeMonths := time.Now().AddDate(0, -3, 0).Unix()
	filter := bson.D{{Key: "unix_created_at", Value: bson.D{{Key: "$gte", Value: threeMonths}}}}
	opts := options.Find().SetSort(bson.D{{Key: "unix_created_at", Value: 1}})

	var cName string
	cName, err = createCollection(person1, person2)
	if err != nil {
		return res, database.ErrCreatingCollection
	}
	c := mg.Client.Database(database.DatabaseName).Collection(cName)

	var cur *mongo.Cursor
	cur, err = c.Find(ctx, filter, opts)
	if err != nil {
		return res, err
	}

	err = cur.All(ctx, &res)
	return res, err

}

func (mg *Mongo) ReadMessage(ctx context.Context, sendeID int64, receiverID int64) error {

	var err error

	filter := bson.D{{Key: "sender_id", Value: sendeID}, {Key: "receiver_id", Value: receiverID}, {Key: "is_read", Value: false}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_read", Value: true}}}}

	var cName string
	cName, err = createCollection(sendeID, receiverID)
	c := mg.Client.Database(database.DatabaseName).Collection(cName)
	c.UpdateMany(ctx, filter, update)

	return err
}

func New(uri string) (database.Database, error) {
	db := &Mongo{}

	var err error
	var client *mongo.Client

	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err != nil {
		return db, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return db, err
	}

	db.Client = client
	return db, nil
}

func (mg *Mongo) CountUnread(ctx context.Context, senderID int64, receiverID int64) (int64, error) {
	var unread int64
	var err error

	var cname string
	cname, err = createCollection(senderID, receiverID)
	if err != nil {
		return unread, err
	}

	c := mg.Client.Database(database.DatabaseName).Collection(cname)

	filter := bson.D{{Key: "sender_id", Value: senderID}, {Key: "receiver_id", Value: receiverID}, {Key: "is_read", Value: false}}
	return c.CountDocuments(ctx, filter)

}
