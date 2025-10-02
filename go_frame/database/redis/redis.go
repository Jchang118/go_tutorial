package distributed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// value是简单的string
func StringValue(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "大乔乔"
	defer client.Del(ctx, key) //函数结束时删除redis上的key,不影响下次运行演示

	err := client.Set(ctx, key, value, 1*time.Second).Err() //1秒后失效.0表示永不失效
	checkError(err)

	client.Expire(ctx, key, 3*time.Second) //通过Expire设置3秒后失效.该方法对任意类型的redis value都适用
	time.Sleep(2 * time.Second)

	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)

	err = client.Set(ctx, "age", 18, 1*time.Second).Err() //int写入redis后会转成string
	checkError(err)
	v3, err := client.Get(ctx, "age").Int()
	checkError(err)
	fmt.Printf("age=%d\n", v3)
}

type Student struct {
	Id   int
	Name string
}

func WriteStudent2Redis(client *redis.Client, stu *Student) error {
	if stu == nil {
		return nil
	}
	key := "STU_" + strconv.Itoa(stu.Id) //避免各种id冲突,用前缀区分
	v, err := json.Marshal(stu)
	if err != nil {
		return err
	}
	err = client.Set(context.Background(), key, string(v), 5*time.Minute).Err()
	return err
}

func GetStudentFromRedis(client *redis.Client, sid int) *Student {
	key := "STU_" + strconv.Itoa(sid) //避免各种id冲突,用前缀区分
	v, err := client.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.Nil { //如果key不存在,会返回redis.Nil
			log.Println(err)
		}
		return nil
	}
	var stu Student
	err = json.Unmarshal([]byte(v), &stu)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &stu
}

func DeleteKey(ctx context.Context, client *redis.Client) {
	n, err := client.Del(ctx, "not_exists").Result()
	if err == nil {
		fmt.Printf("删除%d个key\n", n)
	}
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil { //读redis发生error,大部分情况是因为key不存在
			fmt.Println("key不存在")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
