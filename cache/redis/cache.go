package redis

import (
	"errors"
	"reflect"
	"time"

	"github.com/go-redis/redis"
	"github.com/sebasdeweert/go-kit/cache"
	"github.com/sebasdeweert/go-kit/encoding"
	"github.com/sebasdeweert/go-kit/encoding/gob"
)

// Cache combines the Cache and redis.Cmdable interfaces.
type Cache interface {
	cache.Cache
	redis.Cmdable
}

type rcache struct {
	redis.Cmdable
	encoding.Encoder

	Expiration time.Duration
}

// NewCache returns a new Redis rcache.
func NewCache(cfg *Config) (Cache, error) {
	if cfg.Expiration < time.Second {
		return nil, errors.New(ErrNoExpirationSet)
	}

	return &rcache{
		Cmdable:    redis.NewRing(cfg.GetRingOptions()),
		Encoder:    gob.NewEncoder(),
		Expiration: cfg.Expiration,
	}, nil
}

// GetOne retrieves and decodes an object.
func (c *rcache) GetOne(index string, obj interface{}) error {
	encoded, err := c.Cmdable.Get(index).Result()

	if err != nil {
		// Cache miss.
		if err.Error() == ErrNilResponse {
			return cache.ErrCacheMiss
		}

		return err
	}

	return c.Encoder.Decode(encoded, obj)
}

// SetOne encodes and sets an object in rcache.
func (c *rcache) SetOne(index string, obj interface{}) (err error) {
	encoded, err := c.Encoder.Encode(obj)

	if err != nil {
		return err
	}

	return c.Cmdable.Set(index, *encoded, c.Expiration).Err()
}

// MultiSet encodes and sets multiple objects in rcache.
func (c *rcache) MultiSet(objs map[string]interface{}) error {
	var values []interface{}

	for index, obj := range objs {
		encoded, err := c.Encoder.Encode(obj)

		if err != nil {
			return err
		}

		values = append(values, index, *encoded)
	}

	pipe := c.Cmdable.Pipeline()

	pipe.MSet(values...)

	for index := range objs {
		pipe.Expire(index, c.Expiration)
	}

	_, err := pipe.Exec()

	return err
}

// MultiSet retrieves and decodes multiple objects.
func (c *rcache) MultiGet(objs map[string]interface{}) (map[string]error, error) {
	var indexes []string

	for index := range objs {
		indexes = append(indexes, index)
	}

	result, err := c.Cmdable.MGet(indexes...).Result()

	if err != nil {
		return nil, err
	}

	misses := map[string]error{}

	for i, encoded := range result {
		// Cache miss.
		if encoded == nil {
			misses[indexes[i]] = cache.ErrCacheMiss

			continue
		}

		if err := c.Encoder.Decode(encoded.(string), objs[indexes[i]]); err != nil {
			return nil, err
		}

		misses[indexes[i]] = nil
	}

	return misses, nil
}

// CheckHealth checks the health of the redis client.
func (c *rcache) CheckHealth() error {
	result, err := c.Cmdable.Ping().Result()

	if err != nil {
		return err
	}

	if result != PongMessage {
		return errors.New(ErrNotPong)
	}

	return nil
}

// Delete removes the given indexes.
func (c *rcache) Delete(indexes ...string) error {
	return c.Cmdable.Del(indexes...).Err()
}

// GetSetMembers returns members of the set with the given index.
func (c *rcache) GetSetMembers(index string) ([]string, error) {
	return c.Cmdable.SMembers(index).Result()
}

// GetSetMembersByScore returns members of the set with a given index sorted by score.
func (c *rcache) GetSetMembersByScore(index string, min string, max string, offset int64, count int64, direction cache.Direction) ([]string, error) {
	rangeOptions := redis.ZRangeBy{
		Min:    min,
		Max:    max,
		Offset: offset,
		Count:  count,
	}

	switch direction {
	case cache.Ascending:
		return c.Cmdable.ZRangeByScore(index, rangeOptions).Result()

	case cache.Descending:
		return c.Cmdable.ZRevRangeByScore(index, rangeOptions).Result()

	default:
		return nil, errors.New(cache.ErrInvalidDirection)
	}
}

// GetSetMembersByRank returns members of the set with a given index sorted by ranking of score.
func (c *rcache) GetSetMembersByRank(index string, offset int64, count int64, direction cache.Direction) ([]string, error) {
	stop := offset + count

	switch direction {
	case cache.Ascending:
		return c.Cmdable.ZRange(index, offset, stop).Result()

	case cache.Descending:
		return c.Cmdable.ZRevRange(index, offset, stop).Result()

	default:
		return nil, errors.New(cache.ErrInvalidDirection)
	}
}

// AddToSet adds members to the set with the given index.
func (c *rcache) AddToSet(index string, members []interface{}) error {
	if err := c.Cmdable.SAdd(index, members...).Err(); err != nil {
		return err
	}

	return c.Cmdable.Expire(index, c.Expiration).Err()
}

// GetCacheKeys returns cache keys for specified pattern.
func (c *rcache) GetCacheKeys(pattern string) ([]string, error) {
	return c.Cmdable.Keys(pattern).Result()
}

// PushOneLeft pushes an object to the list from the left.
func (c *rcache) PushOneLeft(index string, obj interface{}) error {
	encoded, err := c.Encoder.Encode(obj)

	if err != nil {
		return err
	}

	if err := c.LPush(index, *encoded).Err(); err != nil {
		return err
	}

	return c.Cmdable.Expire(index, c.Expiration).Err()
}

// PushOneRight pushes an object to the list stored at index key from the right.
func (c *rcache) PushOneRight(index string, obj interface{}) error {
	encoded, err := c.Encoder.Encode(obj)

	if err != nil {
		return err
	}

	if err := c.RPush(index, *encoded).Err(); err != nil {
		return err
	}

	return c.Cmdable.Expire(index, c.Expiration).Err()
}

// GetListOne gets one element of the list stored at index key at offset position into dest.
func (c *rcache) GetListOne(index string, offset int64, dest interface{}) error {
	encoded, err := c.LIndex(index, offset).Result()

	if err != nil {
		if err.Error() == ErrNilResponse {
			return cache.ErrCacheMiss
		}

		return err
	}

	return c.Encoder.Decode(encoded, dest)
}

// GetListRange gets the specified elements of the list stored at index key between start and stop positions into dest.
func (c *rcache) GetListRange(index string, start int64, stop int64, dest interface{}) error {
	v := reflect.ValueOf(dest)

	if !(v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Slice) {
		return cache.ErrDestNotPtrSlice
	}

	v = v.Elem()

	results, err := c.LRange(index, start, stop).Result()

	if err != nil {
		if err.Error() == ErrNilResponse {
			return cache.ErrCacheMiss
		}

		return err
	}

	v.Set(reflect.MakeSlice(v.Type(), len(results), len(results)))

	for i, s := range results {
		obj := reflect.New(v.Index(i).Type()).Interface()

		if err := c.Encoder.Decode(s, obj); err != nil {
			return err
		}

		v.Index(i).Set(reflect.ValueOf(obj).Elem())
	}

	return nil
}

// TrimList trims an existing list stored at index key.
func (c *rcache) TrimList(index string, start int64, stop int64) error {
	return c.LTrim(index, start, stop).Err()
}
