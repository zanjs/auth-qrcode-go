package pool

import (
	"city6/au/models"

	"github.com/garyburd/redigo/redis"
)

var (
	rdspool = newPool()
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 120, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":"+models.Config.Redis.Port)
			if err != nil {
				panic(err.Error())
			}
			_, err = c.Do("AUTH", models.Config.Redis.Password)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}

// GetPool is
func GetPool() redis.Conn {
	c := rdspool.Get()
	// defer c.Close()
	return c
}
