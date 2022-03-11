package mongo

import (
	"fmt"

	"github.com/Keshav-Agrawal/mongoapi/datasource"
	"github.com/Keshav-Agrawal/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockMongo struct {
	f func() (interface{}, error)
}

func (m mockMongo) InsertOneTask(work model.Homework) error {
	//TODO implement me
	panic("implement me")
}

func (m mockMongo) UpdateOneTask(workId string) {
	//TODO implement me
	panic("implement me")
}

func (m mockMongo) DeleteOneTask(workId string) {
	//TODO implement me
	panic("implement me")
}

func (m mockMongo) DeleteAllTask() int64 {
	//TODO implement me
	panic("implement me")
}

func (m mockMongo) GetAllTask() ([]primitive.M, error) {
	fmt.Println("Custom mock")
	resp, err := m.f()
	if err != nil {
		return nil, err
	}
	return resp.([]primitive.M), nil
}

func NewMock(f func() (interface{}, error)) datasource.IDataSource {
	return &mockMongo{
		f: f,
	}
}
