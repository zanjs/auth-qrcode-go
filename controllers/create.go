package controllers

import (
	"fmt"
	"strings"

	"city6/au/models"
	"city6/au/pool"

	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
)

// CreateAppSecret is
func CreateAppSecret(ctx iris.Context) {

	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	v4 := strings.Replace(u2.String(), "-", "", -1)
	fmt.Println(v4)
	key := v4

	c := pool.GetPool()

	_, err = c.Do("SET", key, "", "EX", "120")

	if err != nil {
		fmt.Println("redis set failed:", err)
		ResponseJSONError(ctx, err.Error())
		return
	}
	ResponseJSON(ctx, key)
}

// LinkAppSecret is
func LinkAppSecret(ctx iris.Context) {
	c := &models.Secret{}
	if err := ctx.ReadJSON(c); err != nil {
		ResponseBad(ctx, err.Error())
		return
	}
	fmt.Println(c)
	rq := models.Secret{}

	rq.Key = "c.Key"
	ResponseJSON(ctx, rq)
}

// JoinAppSecret is
func joinAppSecret(c redis.Conn, key string) (string, error) {
	secret, err := redis.String(c.Do("GET", key))

	if err != nil {
		fmt.Println("redis get failed:", err)
		return "", err
	}
	fmt.Printf("Get key: %v \n", secret)
	return secret, nil
}
