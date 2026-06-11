package storage

import (
	"context"
	"hash/fnv"
	// "net/url"
	// "strings"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStore struct{
	Client *redis.Client
}

func NewRedisStore() *RedisStore {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})

	return &RedisStore{Client: rdb};
}

func (s *RedisStore) SaveURL(originalURL string) (string, error){
	h := fnv.New64a()
	h.Write([]byte(originalURL))
	//fmt.Print("The hashing using sum64 ",h.Sum64(), "\n")
	hashURLtoString := fmt.Sprintf("%x" , h.Sum64()); // this is using convert to Hex string, Conver to decimal using '%d' 

	err := s.Client.Set(ctx, originalURL,hashURLtoString,0).Err()
	if err != nil {
		return "" , err
	}

	err = s.Client.Set(ctx,hashURLtoString , originalURL,0).Err() // SET (context, key, val ,experation) so 0 no experation
	if err != nil {
		return "" , err
	}
	
	return hashURLtoString, nil
}

