package pool

import (
	"github.com/garyburd/redigo/redis"
)

var (
	pool = newPool()
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 120, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			_, err = c.Do("AUTH", "root")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

// GetPool is
func GetPool() redis.Conn {
	c := pool.Get()
	// defer c.Close()
	return c
}
