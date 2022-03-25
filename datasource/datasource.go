package datasource

import (
	"github.com/Keshav-Agrawal/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IDataSource interface {
	InsertOneTask(work model.Homework) error
	UpdateOneTask(workId string) (string, error)
	DeleteOneTask(workId string) (string, error)
	DeleteAllTask() (int64, error)
	GetAllTask() ([]primitive.M, error)
}
