package db

import (
	"items_api/models"
	"sync"
)

var storage = NewStorage()

//carsWithStatus only for showcase
type Storage struct {
	cars           []models.Car
	carsWithStatus []models.CarWithStatus
	statuses       []models.Status
	mu             *sync.RWMutex
}

func NewStorage() *Storage {
	s := &Storage{
		cars:           make([]models.Car, 10),           //for simplicity made range of 10
		carsWithStatus: make([]models.CarWithStatus, 10), //for simplicity made range of 10
		statuses:       make([]models.Status, 10),        //for simplicity made range of 10
		mu:             &sync.RWMutex{},
	}

	//insert dummy data
	for _, car := range models.DummyCars() {
		s.Create(car)
	}

	//only for showcase
	for _, carWithS := range models.DummyCarsWithStatuses() {
		s.CreateWithS(carWithS)
	}

	for _, status := range models.DummyStatuses() {
		s.CreateS(status)
	}

	return s
}

func Set(car models.Car) models.Car {
	storage.Create(car)
	return storage.Read(car.Id)
}

//not most elegant way but for demo is ok
func GetAll() []models.Car {
	var cars []models.Car

	for _, car := range storage.cars {
		if car.Id != 0 {
			cars = append(cars, car)
		}
	}
	return cars
}

func Get(id uint32) models.Car {
	return storage.Read(id)
}

//only for showcase
func GetWithStatus(id uint32) models.CarWithStatus {
	return storage.ReadWithStatus(id)
}

//temporary return changed cars array
func Delete(id uint32) []models.Car {
	storage.Delete(id)

	return GetAll()
}

//----------------------------------------
//statuses

func SetS(status models.Status) models.Status {
	storage.CreateS(status)
	return storage.ReadS(status.Id)
}
func GetAllS() []models.Status {
	var statuses []models.Status

	for _, status := range storage.statuses {
		if status.Id != 0 {
			statuses = append(statuses, status)
		}
	}
	return statuses
}
func DeleteS(id uint32) []models.Status {
	storage.DeleteS(id)
	return GetAllS()
}

//----------------------------------------
//car.Id data type is uint32
func (s Storage) Read(key uint32) models.Car {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.cars[key]
}

//in real sql query i would join by id or take from cache
func (s Storage) ReadWithStatus(key uint32) models.CarWithStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.carsWithStatus[key]
}

//here key will be Id
//will be same as update for simplicity
func (s Storage) Create(car models.Car) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cars[car.Id] = car
}

//in reality i would create in same as default
func (s Storage) CreateWithS(car models.CarWithStatus) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.carsWithStatus[car.Id] = car
}

//just return changed array
func (s Storage) Delete(key uint32) []models.Car {
	cars := s.cars
	cars[key] = cars[len(cars)-1]

	return cars[:len(cars)-1]
}

//----------------------------------------
//statuses

func (s Storage) ReadS(key uint32) models.Status {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.statuses[key]
}
func (s Storage) CreateS(status models.Status) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.statuses[status.Id] = status
}
func (s Storage) DeleteS(key uint32) []models.Status {
	statuses := s.statuses
	statuses[key] = statuses[len(statuses)-1]

	return statuses[:len(statuses)-1]
}
