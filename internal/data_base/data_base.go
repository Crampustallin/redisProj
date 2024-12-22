package data_base

import (
	"context"

	"github.com/Crampustallin/redisProj/internal/model"
	"github.com/redis/go-redis/v9"
)

type DataBase struct {
	r *redis.Client
}

func NewDataBase(address string) *DataBase {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return &DataBase{
		r: client,
	}
}

func (db *DataBase) SetUser(u *model.User) error {
	return db.r.HSet(context.Background(), u.Login, u).Err()
}

func (db *DataBase) GetUser(name string) (u *model.User, err error) {
	user := model.User{}
	if err := db.r.HGetAll(context.Background(), name).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DataBase) Close() (err error) {
	return db.r.Close()
}
