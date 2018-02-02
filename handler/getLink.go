package handler

import (
	"city6/au/pool"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
)

// GetLinkUser is
func GetLinkUser(ctx iris.Context) {
	// fms := ctx.FormValues()
	// fmt.Println(fms)

	skey := ctx.FormValue("skey")

	if skey == "" {
		ResponseBad(ctx, "skey where is go?")
		return
	}

	c := pool.RedisClient.Get()
	defer c.Close()

	secret, err := redis.String(c.Do("GET", skey))

	if err != nil {
		fmt.Println("redis get failed:", err)
		ResponseJSONError(ctx, err.Error())
		return
	}

	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(secret), &dat); err == nil {
		fmt.Println(dat)
		ResponseJSON(ctx, dat)
		return
	}

	ResponseJSON(ctx, secret)
}
