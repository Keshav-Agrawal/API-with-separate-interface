package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Keshav-Agrawal/mongoapi/datasource"
	"github.com/Keshav-Agrawal/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://keshav1:keshav1@cluster0.sjkrk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "homework"
const colName = "task"

type mongoDS struct {
	c *mongo.Collection
}

//dont use init
func NewDs() datasource.IDataSource {
	var err error
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	return &mongoDS{client.Database(dbName).Collection(colName)}
}

func (m mongoDS) InsertOneTask(work model.Homework) error {
	inserted, err := m.c.InsertOne(context.Background(), work)

	if err != nil {
		log.Println("InsertOneTask::error::", err)
		return errors.New("insertion failure")
	}
	fmt.Println("Inserted 1 task in db with id: ", inserted.InsertedID)
	return nil
}

func (m mongoDS) UpdateOneTask(workId string) {
	id, _ := primitive.ObjectIDFromHex(workId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"done": true}}

	result, err := m.c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func (m mongoDS) DeleteOneTask(workId string) {
	id, _ := primitive.ObjectIDFromHex(workId)
	filter := bson.M{"_id": id}
	deleteCount, err := m.c.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("task got delete with delete count: ", deleteCount)
}

func (m mongoDS) DeleteAllTask() int64 {

	deleteResult, err := m.c.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of task delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func (m mongoDS) GetAllTask() ([]primitive.M, error) {
	cur, err := m.c.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var worklist []primitive.M

	for cur.Next(context.Background()) {
		var work bson.M
		err := cur.Decode(&work)
		if err != nil {
			log.Fatal(err)
		}
		worklist = append(worklist, work)

	}

	defer cur.Close(context.Background())
	return worklist, nil
}
