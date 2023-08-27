package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) //memanggil parameter findbyid
	if arguments.Get(0) == nil {            //Get(0) = data pertama dalam array
		return nil
	} else {
		category := arguments.Get(0).(entity.Category) //di conversikan ke (entity.Category)
		return &category
	}

}
