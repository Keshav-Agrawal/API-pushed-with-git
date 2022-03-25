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
	_, err := m.f()
	return err
}

func (m mockMongo) UpdateOneTask(workId string) (string, error) {
	//TODO implement me
	panic("implement me")
	_, err := m.f()
	return (workId), err
}

func (m mockMongo) DeleteOneTask(workId string) (string, error) {
	//TODO implement me
	panic("implement me")
	_, err := m.f()
	return (workId), err
}

func (m mockMongo) DeleteAllTask() (int64, error) {
	//TODO implement me
	panic("implement me")
	resp, err := m.f()
	if err != nil {
		return 0, err
	}
	return resp.(int64), nil
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
