package data_base

import (
	"github.com/Crampustallin/redisProj/internal/model"
	"github.com/redis/go-redis/v9"
)

type DataBase struct {
	r *redis.Client
}

func NewDataBase() *DataBase {
	return nil
}

func (db *DataBase) SetUser(u *model.User) error {

}
