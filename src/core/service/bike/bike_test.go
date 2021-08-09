package bike

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	"github.com/rickcorilaco/api-bike-v3/src/core/values"
	bikeRepositoryMocks "github.com/rickcorilaco/api-bike-v3/src/repository/bike/mocks"
)

var (
	repository           *bikeRepositoryMocks.Repository
	service              *BikeService
	caloi10, absoluteMia *entity.Bike
	errRepository        error
)

func TestMain(m *testing.M) {
	var err error

	repository = &bikeRepositoryMocks.Repository{}

	service, err = New(repository)
	if err != nil {
		log.Fatal(err)
	}

	caloi10 = &entity.Bike{
		ID:    uuid.New(),
		Brand: "Caloi",
		Model: "10",
	}

	absoluteMia = &entity.Bike{
		ID:    uuid.New(),
		Brand: "Absolute",
		Model: "Mia",
	}

	errRepository = errors.New("mock repository error")

	os.Exit(m.Run())
}

func TestList_should_get_a_bike_list_with_success(t *testing.T) {
	filter := values.BikeListFilter{}

	expected := []entity.Bike{
		*caloi10,
		*absoluteMia,
	}

	repository.On("List", filter).Return(expected, nil)

	received, err := service.List(filter)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if err != nil {
		t.Errorf("\nError should be nil, received: %v", err)
	}
}

func TestList_should_not_get_a_bike_list_when_the_repository_fails(t *testing.T) {
	var (
		expected    []entity.Bike
		errExpected = errRepository
	)

	filter := values.BikeListFilter{
		Brand: &caloi10.Brand,
	}

	repository.On("List", filter).Return(expected, errExpected)

	received, err := service.List(filter)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if !errors.Is(errExpected, err) {
		t.Errorf("\nExpected: %v \nGot: %v\n", errExpected, err)
	}
}
func TestGet_should_get_a_bike_with_success(t *testing.T) {
	expected := caloi10

	repository.On("Get", expected.ID).Return(expected, nil)

	received, err := service.Get(expected.ID)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if err != nil {
		t.Errorf("\nError should be nil, received: %v", err)
	}
}

func TestGet_should_not_get_a_bike_when_the_repository_fails(t *testing.T) {
	var (
		bikeID      = uuid.New()
		expected    *entity.Bike
		errExpected = errors.New("repository error")
	)

	repository.On("Get", bikeID).Return(expected, errExpected)

	received, err := service.Get(bikeID)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if !errors.Is(errExpected, err) {
		t.Errorf("\nExpected: %v \nGot: %v\n", errExpected, err)
	}
}

func TestCreate_should_create_a_bike_with_success(t *testing.T) {
	expected := absoluteMia

	repository.On("Create", *expected).Return(expected, nil)

	received, err := service.Create(*expected)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if err != nil {
		t.Errorf("\nError should be nil, received: %v", err)
	}
}

func TestCreate_should_not_create_a_bike_when_the_repository_fails(t *testing.T) {
	bike := caloi10

	var expected *entity.Bike

	errExpected := errors.New("repository error")

	repository.On("Create", *bike).Return(expected, errExpected)

	received, err := service.Create(*bike)

	if !reflect.DeepEqual(expected, received) {
		t.Errorf("\nExpected: %v \nGot: %v\n", expected, received)
	}

	if errExpected != err {
		t.Errorf("\nExpected: %v \nGot: %v\n", errExpected, err)
	}
}

func TestDelete_should_delete_a_bike_with_success(t *testing.T) {
	expected := caloi10

	repository.On("Delete", *expected).Return(nil)

	err := service.Delete(*expected)

	if err != nil {
		t.Errorf("\nError should be nil, received: %v", err)
	}
}

func TestDelete_should_not_delete_a_bike_when_the_repository_fails(t *testing.T) {
	var (
		bike        = absoluteMia
		errExpected = errRepository
	)

	repository.On("Delete", *bike).Return(errExpected)

	err := service.Delete(*bike)

	if errExpected != err {
		t.Errorf("\nExpected: %v \nGot: %v\n", errExpected, err)
	}
}
