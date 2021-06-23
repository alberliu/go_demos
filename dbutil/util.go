package dbutil

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	ErrNotFound = errors.New("not found error")
)

type DBUtil struct {
	db    *gorm.DB
	redis redis.Client
}

func NewDBUtil(db *gorm.DB, redis redis.Client) *DBUtil {
	return &DBUtil{
		db:    db,
		redis: redis,
	}
}

type RedisKey interface {
	RedisKey() string
	Expiration() time.Duration
}

// SetToCache 将指定值设置到redis中，使用json的序列化方式
func (u *DBUtil) SetToCache(value RedisKey) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = u.redis.Set(value.RedisKey(), bytes, value.Expiration()).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetFromCache 从redis中读取指定值，使用json的反序列化方式
func (u *DBUtil) GetFromCache(value RedisKey) error {
	bytes, err := u.redis.Get(value.RedisKey()).Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}

// Get 先从redis获取，如果查询不到就从数据库获取，再设置缓存
func (u *DBUtil) Get(value RedisKey, where ...interface{}) error {
	err := u.GetFromCache(value)
	if err != nil && err != redis.Nil {
		return err
	}
	if err == nil {
		return nil
	}

	err = u.db.First(value, where).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	err = u.SetToCache(value)
	if err != nil {
		return err
	}
	return nil
}

// Update 先从redis获取，如果查询不到就从数据库获取，再设置缓存
func (u *DBUtil) Update(value RedisKey, where ...interface{}) error {
	err := u.GetFromCache(value)
	if err != nil && err != redis.Nil {
		return err
	}
	if err == nil {
		return nil
	}

	err = u.db.First(value, where).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}

	err = u.SetToCache(value)
	if err != nil {
		return err
	}
	return nil
}
