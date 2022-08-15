package redis_test

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func TestRedisOper(t *testing.T) {
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	fmt.Println("Redis连接成功")

	// Set & Get示例
	// 第三个参数代表key的过期时间，0代表不会过期。
	err = redisDb.Set("username", "tom", 0).Err()
	if err != nil {
		panic(err)
	}

	var val string
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err = redisDb.Get("username").Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("username: ", val)

	//批量设置key1对应的值为value1，key2对应的值为value2，key3对应的值为value3，
	err = redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}

	// MGet函数可以传入任意个key，一次性返回多个值。
	// 这里Result返回两个值，第一个值是一个数组，第二个值是错误信息
	vals, err := redisDb.MGet("key1", "key2", "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals) //[value1 value2 value3]

	// Hash Operation
	//根据key和field字段设置,field字段的值。 user_1 是hash key，username 是字段名, admin是字段值
	err = redisDb.HSet("user_1", "username", "admin").Err()
	if err != nil {
		panic(err)
	}
	//根据key和field字段,查询field字段的值。user_1 是hash key，username是字段名
	username, err := redisDb.HGet("user_1", "username").Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(username) //admin
	//继续往user_1中添加字段password
	_ = redisDb.HSet("user_1", "password", "abc123").Err()

	// HGetAll 一次性返回key=user_1的所有hash字段和值
	data, err := redisDb.HGetAll("user_1").Result()
	if err != nil {
		panic(err)
	}
	// data是一个map类型，这里使用使用循环迭代输出
	for field, val := range data {
		fmt.Println(field, val)
	}

	// 初始化hash数据的多个字段值
	batchData := make(map[string]interface{})
	batchData["username"] = "test"
	batchData["password"] = 123456
	// 一次性保存多个hash字段值
	err = redisDb.HMSet("user_2", batchData).Err()
	if err != nil {
		panic(err)
	}

	//如果email字段不存在，则设置hash字段值
	redisDb.HSetNX("user_2", "email", "ourlang@foxmail.com")
	// HMGet支持多个field字段名，意思是一次返回多个字段值
	values, err := redisDb.HMGet("user_2", "username", "password", "email").Result()
	if err != nil {
		panic(err)
	}
	// values是一个数组
	fmt.Println("user_2=", values) //user_2= [test 123456 ourlang@foxmail.com]
}
