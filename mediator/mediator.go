package mediator

import (
	"fmt"
	"sync"
)

type train interface {
	requestArrival()
	departure()
	permitArrival()
}


type mediator interface {
	canLand(train2 train) bool
	notifyFree()
}
type passengerTrain struct {
	mediator mediator
}
func (t *passengerTrain) permitArrival()  {
	fmt.Println("PassengerTrain arrival permitted.")
}
func (t *passengerTrain) departure()  {
	fmt.Println("PassengerTrain leaving")
	t.mediator.notifyFree()
}
func (t *passengerTrain) requestArrival()  {
	if t.mediator.canLand(t) {
		fmt.Println("PassengerTrain landing")
	}else {
		fmt.Println("passengerTrain waiting")
	}
}

type GoodsTrain struct {
	mediator mediator
}

func (t *GoodsTrain) permitArrival()  {
	fmt.Println("GoodsTrain arrival permitted.")
}
func (t *GoodsTrain) departure()  {
	fmt.Println("GoodsTrain leaving")
	t.mediator.notifyFree()
}
func (t *GoodsTrain) requestArrival()  {
	if t.mediator.canLand(t) {
		fmt.Println("GoodsTrain landing")
	}else {
		fmt.Println("GoodsTrain waiting")
	}
}

type stationManager struct {
	isPlatformFree bool
	lock * sync.Mutex
	trainQueue []train
}
func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}
func (s *stationManager)notifyFree()  {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTainInQueue.permitArrival()
	}
}
func TestMediator(){
	stationManager := newStationManager()
	passengerTrain := &passengerTrain{mediator: stationManager}
	goodsTrain := &GoodsTrain{mediator: stationManager}

	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}