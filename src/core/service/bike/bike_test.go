package bike

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports/mocks"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

var (
	repository           *mocks.BikeRepository
	service              ports.BikeService
	caloi10, absoluteMia *domain.Bike
	errRepository        error
)

func TestMain(m *testing.M) {
	var err error

	repository = &mocks.BikeRepository{}

	service, err = New(repository)
	if err != nil {
		log.Fatal(err)
	}

	caloi10 = &domain.Bike{
		ID:    uuid.New(),
		Brand: "Caloi",
		Model: "10",
	}

	absoluteMia = &domain.Bike{
		ID:    uuid.New(),
		Brand: "Absolute",
		Model: "Mia",
	}

	errRepository = errors.New("mock repository error")

	os.Exit(m.Run())
}

func TestList_should_get_a_bike_list_with_success(t *testing.T) {
	filter := value.BikeListFilter{}

	expected := &domain.Bikes{
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
		expected    *domain.Bikes
		errExpected = errRepository
	)

	filter := value.BikeListFilter{
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
		expected    *domain.Bike
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

	var expected *domain.Bike

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

	repository.On("Get", expected.ID).Return(expected, nil)
	repository.On("Delete", *expected).Return(nil)

	_, err := service.Delete(*expected)

	if err != nil {
		t.Errorf("\nError should be nil, received: %v", err)
	}
}

func TestDelete_should_not_delete_a_bike_when_the_repository_fails(t *testing.T) {
	var (
		bike        = absoluteMia
		errExpected = errRepository
	)

	repository.On("Get", bike.ID).Return(bike, nil)
	repository.On("Delete", *bike).Return(errExpected)

	_, err := service.Delete(*bike)

	if errExpected != err {
		t.Errorf("\nExpected: %v \nGot: %v\n", errExpected, err)
	}
}
