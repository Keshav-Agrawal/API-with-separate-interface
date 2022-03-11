package datasource

import (
	"github.com/Keshav-Agrawal/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDataSource interface {
	InsertOneTask(work model.Homework) error
	UpdateOneTask(workId string)
	DeleteOneTask(workId string)
	DeleteAllTask() int64
	GetAllTask() ([]primitive.M, error)
}
