package connectionPool

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type ConnectionPool struct {
	mutex sync.RWMutex
	connections chan net.Conn
	maxSize uint64
}

var TotalConnections uint64
func CreateConnections() (net.Conn, error) {
	atomic.AddUint64(&TotalConnections, 1)
	return net.Dial("tcp", "localhost:8081")
}
func CreateConnectionPool(initialSize int, maxinumSize uint64) (*ConnectionPool, error) {
	pool := &ConnectionPool{
		connections: make(chan net.Conn, maxinumSize),
		maxSize:     maxinumSize,
	}
	TotalConnections = 0
	for iterator := 0; iterator < initialSize; iterator++ {
		atomic.AddUint64(&TotalConnections, 1)
		singleConnection, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			log.Fatal("error in creating initial connections:", err.Error())
		}
		pool.connections <- singleConnection
	}
	return pool, nil
}
func(pool *ConnectionPool) Length() int {
	allConnections := pool.GetPool()
	return len(allConnections)
}
func (pool *ConnectionPool)  GetPool()chan net.Conn {
	return pool.connections
}
func (pool *ConnectionPool)ClosePool()  {
	pool.mutex.Lock()
	pool.connections = nil
	pool.mutex.Unlock()

	TotalConnections = 0
}
func (pool *ConnectionPool) GetOneConnection() (net.Conn, error) {
	allConnections := pool.GetPool()
	if allConnections == nil {
		return nil, errors.New("channel is empty")
	}

	if atomic.LoadUint64(&TotalConnections) >= pool.maxSize {
		singleConnection := <-allConnections
		return singleConnection, nil
	} else {
		select {
		case singleConnection := <-allConnections:
			if singleConnection == nil {
				return nil, errors.New("returned a nil connection.")
			}
			return singleConnection, nil
		default:
			atomic.AddUint64(&TotalConnections, 1)
			return net.Dial("tcp", "localhost:8081")
		}
	}
}

func (pool *ConnectionPool) PutOneConnection(singleConnection net.Conn) error {
	if singleConnection == nil {
		return nil
	}
	if pool.connections == nil {
		singleConnection.Close()
		return errors.New("Pool was already closed.")
	}
	select {
	case pool.connections <-singleConnection:
	default:
		fmt.Println("Pool is full! connection closed.")
		singleConnection.Close()
	}
	return nil
}