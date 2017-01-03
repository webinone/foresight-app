package libs

import (
	REDIS "gopkg.in/redis.v5"
	"time"
)

var RedisClient *REDIS.Client

func LoadRedisClient () {

	RedisClient = REDIS.NewClient(&REDIS.Options{
		Addr	    :         Config.REDIS.Url,
		Password    : Config.REDIS.Password,
		DB	    : Config.REDIS.Db,  // use default DB
		DialTimeout :  10 * time.Second,
		ReadTimeout :  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize    :     10,
		PoolTimeout :  30 * time.Second,
	})
	//RedisClient.FlushDb()
}