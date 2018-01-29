package handler

import (
	"encoding/json"
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

	_, err = c.Do("SET", key, "", "EX", models.Config.Redis.Expire)

	if err != nil {
		fmt.Println("redis set failed:", err)
		ResponseJSONError(ctx, err.Error())
		return
	}
	ResponseJSON(ctx, key)
}

// LinkAppSecret is
func LinkAppSecret(ctx iris.Context) {
	data := &models.Secret{}
	if err := ctx.ReadJSON(data); err != nil {
		ResponseBad(ctx, err.Error())
		return
	}
	fmt.Println(data)

	if data.Token == "" || data.Name == "" {
		ResponseBad(ctx, "user where is go?")
		return
	}

	c := pool.GetPool()
	err := joinAppSecret(c, data.Key)
	if err != nil {
		ResponseBad(ctx, err.Error())
		return
	}

	user := models.User{}

	user = data.User

	jsons, errs := json.Marshal(user) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}

	_, err = c.Do("SET", data.Key, jsons, "EX", models.Config.Redis.Expire)

	if err != nil {
		ResponseBad(ctx, err.Error())
		return
	}

	ResponseJSON(ctx, data)
}

// JoinAppSecret is
func joinAppSecret(c redis.Conn, key string) error {
	secret, err := redis.String(c.Do("GET", key))

	if err != nil {
		fmt.Println("redis get failed:", err)
		return err
	}
	fmt.Printf("Get key: %v \n", secret)
	return nil
}
