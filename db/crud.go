package db

import (
	"items_api/models"
	"sync"
)

var storage = NewStorage()

type Storage struct {
	cars  []models.Car
	mu    *sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		cars:  make([]models.Car, 10), //for simplicity made range of 10
		mu:    &sync.RWMutex{},
	}
}

func Set(car models.Car) models.Car {
	storage.Create(car)
	return storage.Read(car.Id)
}

//not most elegant way but for demo is ok
func GetAll() []models.Car {
	var cars []models.Car

	for _, car := range storage.cars {
		if car.Id!= 0 {
			cars = append(cars, car)
		}
	}
	return cars
}

func Get(id uint32) models.Car {
	return storage.Read(id)
}

//temporary return changed cars array
func Delete(id uint32) []models.Car {
	storage.Delete(id)

	return GetAll()
}
//----------------------------------------
func (s Storage) ReadAll() []models.Car {
	return s.cars
}

//car.Id data type is uint32
func (s Storage) Read(key uint32) models.Car{
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.cars[key]
}

//here key will be Id
//will be same as update for simplicity
func (s Storage) Create(car models.Car) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cars[car.Id] = car
}
//just return changed array
func (s Storage) Delete (key uint32) []models.Car {
	cars := s.cars
	cars[key] = cars[len(cars)-1]

	return cars[:len(cars)-1]
}