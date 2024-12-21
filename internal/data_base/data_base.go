package data_base

import (
	"context"

	"github.com/Crampustallin/redisProj/internal/model"
	"github.com/redis/go-redis/v9"
)

type DataBase struct {
	r *redis.Client
}

func NewDataBase(r *redis.Client) *DataBase {
	return &DataBase{
		r: r,
	}
}

func (db *DataBase) SetUser(u *model.User) error {
	return db.r.HSet(context.Background(), u.Name, u).Err()
}

func (db *DataBase) GetUser(name string) (u *model.User, err error) {
	if err = db.r.Get(context.Background(), name).Scan(u); err != nil {
		return nil, err
	}
	return u, nil
}
