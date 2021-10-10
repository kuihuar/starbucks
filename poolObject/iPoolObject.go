package poolObject

import (
	"fmt"
	"sync"
)

type IPoolObject interface {
	GetID() string
}

type pool struct {
	idle []IPoolObject
	active []IPoolObject
	capacity int
	mulock *sync.Mutex
}

func InitPool(poolObjects []IPoolObject) (*pool, error) {
	if len(poolObjects) ==0 {
		return nil, fmt.Errorf("Cannot create a pool of 0 length")
	}
	active := make([]IPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}
//拿一个
func (p *pool)Loan() (IPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free.Please request after sometime")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	return obj,nil
}

//归还
func (p *pool)Receive(target IPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	return nil
}
func(p *pool) remove (target IPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.GetID() == target.GetID() {
			p.active[currentActiveLength - 1], p.active[i] = p.active[i], p.active[currentActiveLength - 1]
			p.active = p.active[:currentActiveLength - 1]
			return nil
		}
	}
	return fmt.Errorf("Target pool object doesn't belong to the pool")
}