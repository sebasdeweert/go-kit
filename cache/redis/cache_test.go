package redis

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/sebasdeweert/go-kit/cache"
	"github.com/sebasdeweert/go-kit/encoding/gob"
	"github.com/sebasdeweert/go-kit/mocks"
	"github.com/sebasdeweert/go-kit/test"
	"github.com/sebasdeweert/go-kit/types"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_NewCache(t *testing.T) {
	Convey("NewCache()", t, func() {
		Convey("Returns an error when no expiration is set", func() {
			c, err := NewCache(&Config{
				DB: 1,
			})

			So(err.Error(), ShouldEqual, ErrNoExpirationSet)
			So(c, ShouldBeNil)
		})

		Convey("Returns an error when the expiration is set to something silly", func() {
			c, err := NewCache(&Config{
				Expiration: time.Millisecond * 500,
				DB:         1,
			})

			So(err.Error(), ShouldEqual, ErrNoExpirationSet)
			So(c, ShouldBeNil)
		})

		Convey("Returns a redis cache with the given options and default configuration", func() {
			c, err := NewCache(&Config{
				Expiration: time.Hour,
				DB:         1,
			})

			So(err, ShouldBeNil)
			So(c.(*rcache).Cmdable.(*redis.Ring).Options().DB, ShouldEqual, 1)
			So(c.(*rcache).Encoder, ShouldResemble, gob.NewEncoder())
			So(c.(*rcache).Expiration, ShouldEqual, time.Hour)
		})
	})
}

func Test_rcache_Get(t *testing.T) {
	Convey("*rcache.GetOne()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
			Encoder: encoderMock,
		}

		Convey("Returns ErrCacheMiss when Cmdable.Get() returns the redis: nil error", func() {
			redisMock.EXPECT().
				Get("foo").
				Return(redis.NewStringResult(
					"bar",
					errors.New(ErrNilResponse),
				))

			err := c.GetOne("foo", nil)

			So(err, ShouldResemble, cache.ErrCacheMiss)
		})

		Convey("Returns an unhandled error returned by Cmdable.Get()", func() {
			redisMock.EXPECT().
				Get("foo").
				Return(redis.NewStringResult(
					"bar",
					errors.New("biz"),
				))

			err := c.GetOne("foo", nil)

			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns the error returned by Encoder.Decode()", func() {
			redisMock.EXPECT().
				Get("foo").
				Return(redis.NewStringResult(
					"bar",
					nil,
				))

			encoderMock.EXPECT().
				Decode("bar", "biz").
				Return(errors.New("baz"))

			err := c.GetOne("foo", "biz")

			So(err.Error(), ShouldEqual, "baz")
		})
	})
}

func Test_rcache_SetOne(t *testing.T) {
	Convey("*rcache.SetOne()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the error returned by Encoder.Encode()", func() {
			encoderMock.EXPECT().
				Encode("foo").
				Return(
					nil,
					errors.New("biz"),
				)

			err := c.SetOne("buz", "foo")

			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns the error returned by Cmdable.Set().Error()", func() {
			encoderMock.EXPECT().
				Encode("foo").
				Return(types.String("bar"), nil)

			redisMock.EXPECT().
				Set("biz", "bar", time.Minute).
				Return(redis.NewStatusResult(
					"baz",
					errors.New("buz"),
				))

			err := c.SetOne("biz", "foo")

			So(err.Error(), ShouldEqual, "buz")
		})
	})
}

func Test_rcache_MultiSet(t *testing.T) {
	Convey("*rcache.MultiSet()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)
		pipelineMock := mocks.NewMockPipeliner(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns an error returned by Encoder.Encode()", func() {
			encoderMock.EXPECT().
				Encode("foo").
				Return(
					nil,
					errors.New("bar"),
				)

			err := c.MultiSet(map[string]interface{}{
				"biz": "foo",
			})

			So(err.Error(), ShouldEqual, "bar")
		})

		Convey("Returns the error returned by pipe.Exec()", func() {
			encoderMock.EXPECT().
				Encode("foo").
				Return(
					types.String("bar"),
					nil,
				)

			encoderMock.EXPECT().
				Encode("biz").
				Return(
					types.String("baz"),
					nil,
				)

			redisMock.EXPECT().
				Pipeline().
				Return(pipelineMock)

			pipelineMock.EXPECT().
				MSet(test.WithElements([]interface{}{"qux", "bar", "qix", "baz"}))

			pipelineMock.EXPECT().
				Expire("qux", time.Minute)

			pipelineMock.EXPECT().
				Expire("qix", time.Minute)

			pipelineMock.EXPECT().
				Exec().
				Return(
					nil,
					errors.New("buz"),
				)

			err := c.MultiSet(map[string]interface{}{
				"qux": "foo",
				"qix": "biz",
			})

			So(err.Error(), ShouldEqual, "buz")
		})
	})
}

func Test_rcache_MultiGet(t *testing.T) {
	Convey("*rcache.MultiGet()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
			Encoder: encoderMock,
		}

		Convey("Returns the error returned by redis.MGet()", func() {
			redisMock.EXPECT().
				MGet("foo").
				Return(redis.NewSliceResult(
					nil,
					errors.New("bar"),
				))

			misses, err := c.MultiGet(map[string]interface{}{
				"foo": "biz",
			})

			So(misses, ShouldBeNil)
			So(err.Error(), ShouldEqual, "bar")
		})

		Convey("Returns nil results returned from redis.MGet() on the misses return value", func() {
			redisMock.EXPECT().
				MGet("foo").
				Return(redis.NewSliceResult(
					[]interface{}{nil},
					nil,
				))

			misses, err := c.MultiGet(map[string]interface{}{
				"foo": nil,
			})

			So(misses, ShouldResemble, map[string]error{"foo": cache.ErrCacheMiss})
			So(err, ShouldBeNil)
		})

		Convey("Returns an error returned by Encoder.Decode()", func() {
			redisMock.EXPECT().
				MGet("foo").
				Return(redis.NewSliceResult(
					[]interface{}{"bar"},
					nil,
				))

			encoderMock.EXPECT().
				Decode("bar", nil).
				Return(errors.New("biz"))

			misses, err := c.MultiGet(map[string]interface{}{
				"foo": nil,
			})

			So(misses, ShouldBeNil)
			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns no cache misses and nil", func() {
			redisMock.EXPECT().
				MGet(
					test.WithElements([]string{"foo", "bar"}),
				).
				Return(redis.NewSliceResult(
					[]interface{}{"biz", "baz"},
					nil,
				))

			encoderMock.EXPECT().
				Decode("biz", nil).
				Return(nil)

			encoderMock.EXPECT().
				Decode("baz", nil).
				Return(nil)

			misses, err := c.MultiGet(map[string]interface{}{
				"foo": nil,
				"bar": nil,
			})

			So(err, ShouldBeNil)
			So(misses, ShouldResemble, map[string]error{
				"foo": nil,
				"bar": nil,
			})
		})
	})
}

func Test_rcache_HealthCheck(t *testing.T) {
	Convey("*rcache.HealthCheck()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
		}

		Convey("Returns the error returned by Cmdable.Ping()", func() {
			redisMock.EXPECT().
				Ping().
				Return(redis.NewStatusResult(
					"",
					errors.New("foo"),
				))

			err := c.CheckHealth()

			So(err.Error(), ShouldEqual, "foo")
		})

		Convey("Returns the unexpected value returned by Cmdable.Ping()", func() {
			redisMock.EXPECT().
				Ping().
				Return(redis.NewStatusResult(
					ErrNotPong,
					nil,
				))

			err := c.CheckHealth()

			So(err.Error(), ShouldEqual, "unexpected return value")
		})

		Convey("Returns nil returned by Cmdable.Ping()", func() {
			redisMock.EXPECT().
				Ping().
				Return(redis.NewStatusResult(
					"PONG",
					nil,
				))

			err := c.CheckHealth()

			So(err, ShouldBeNil)
		})
	})
}

func Test_rcache_Delete(t *testing.T) {
	Convey("*rcache.Delete()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
		}

		Convey("Returns the error returned by Cmdable.Del()", func() {
			redisMock.EXPECT().
				Del("foo", "bar").
				Return(redis.NewIntResult(
					1,
					errors.New("biz"),
				))

			err := c.Delete("foo", "bar")

			So(err.Error(), ShouldEqual, "biz")
		})
	})
}

func Test_rcache_GetSetMembers(t *testing.T) {
	Convey("*rcache.GetSetMembers()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
		}

		Convey("Returns the result returned by Cmdable.SMembers()", func() {
			redisMock.EXPECT().
				SMembers("foo").
				Return(redis.NewStringSliceResult(
					[]string{"bar"},
					errors.New("biz"),
				))

			res, err := c.GetSetMembers("foo")

			So(res, ShouldResemble, []string{"bar"})
			So(err.Error(), ShouldEqual, "biz")
		})
	})
}

func Test_rcache_GetSetMembersByScore(t *testing.T) {
	Convey("*rcache.GetSetMembersByScore()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the result returned by Cmdable.ZRangeByScore() when ascending order is provided", func() {
			redisMock.EXPECT().
				ZRangeByScore(
					"foo",
					redis.ZRangeBy{
						Min:    "min",
						Max:    "max",
						Offset: int64(1),
						Count:  int64(2),
					},
				).
				Return(redis.NewStringSliceResult(
					[]string{"bar"},
					errors.New("biz"),
				))

			res, err := c.GetSetMembersByScore("foo", "min", "max", int64(1), int64(2), cache.Ascending)

			So(res, ShouldResemble, []string{"bar"})
			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns the result returned by Cmdable.ZRevRangeByScore() when descending order is provided", func() {
			redisMock.EXPECT().
				ZRevRangeByScore(
					"foo",
					redis.ZRangeBy{
						Min:    "min",
						Max:    "max",
						Offset: int64(1),
						Count:  int64(2),
					},
				).
				Return(redis.NewStringSliceResult(
					[]string{"bar"},
					errors.New("biz"),
				))

			res, err := c.GetSetMembersByScore("foo", "min", "max", int64(1), int64(2), cache.Descending)

			So(res, ShouldResemble, []string{"bar"})
			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns an error when an unknown direction is provided", func() {
			res, err := c.GetSetMembersByScore("", "", "", int64(0), int64(0), 3)

			So(res, ShouldBeNil)
			So(err.Error(), ShouldEqual, cache.ErrInvalidDirection)
		})
	})
}

func Test_rcache_GetSetMembersByRank(t *testing.T) {
	Convey("*rcache.GetSetMembersByRank()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the result returned by Cmdable.ZRangeByScore() when ascending order is provided", func() {
			redisMock.EXPECT().
				ZRange("foo", int64(1), int64(3)).
				Return(redis.NewStringSliceResult(
					[]string{"bar"},
					errors.New("biz"),
				))

			res, err := c.GetSetMembersByRank("foo", int64(1), int64(2), cache.Ascending)

			So(res, ShouldResemble, []string{"bar"})
			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns the result returned by Cmdable.ZRevRangeByRank() when descending order is provided", func() {
			redisMock.EXPECT().
				ZRevRange("foo", int64(1), int64(3)).
				Return(redis.NewStringSliceResult(
					[]string{"bar"},
					errors.New("biz"),
				))

			res, err := c.GetSetMembersByRank("foo", int64(1), int64(2), cache.Descending)

			So(res, ShouldResemble, []string{"bar"})
			So(err.Error(), ShouldEqual, "biz")
		})

		Convey("Returns an error when an unknown direction is provided", func() {
			res, err := c.GetSetMembersByRank("", int64(0), int64(0), 3)

			So(res, ShouldBeNil)
			So(err.Error(), ShouldEqual, cache.ErrInvalidDirection)
		})
	})
}

func Test_rcache_AddToSet(t *testing.T) {
	Convey("*rcache.AddToSet()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Expiration: time.Minute,
		}

		Convey("Returns the error returned by Cmdable.SAdd()", func() {
			redisMock.EXPECT().
				SAdd("foo", "bar", "biz").
				Return(redis.NewIntResult(
					1,
					errors.New("baz"),
				))

			err := c.AddToSet("foo", []interface{}{"bar", "biz"})

			So(err.Error(), ShouldEqual, "baz")
		})

		Convey("Returns the error returned by Cmdable.Expire()", func() {
			redisMock.EXPECT().
				SAdd("foo", "bar", "biz").
				Return(redis.NewIntResult(
					1,
					nil,
				))

			redisMock.EXPECT().
				Expire("foo", time.Minute).
				Return(redis.NewBoolResult(
					true,
					errors.New("baz"),
				))

			err := c.AddToSet("foo", []interface{}{"bar", "biz"})

			So(err.Error(), ShouldEqual, "baz")
		})
	})
}

func Test_rcache_GetCacheKeys(t *testing.T) {
	Convey("*rcache.GetCacheKeys()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable: redisMock,
		}

		Convey("Returns the result of Cmdable.Keys().Result()", func() {
			redisMock.EXPECT().
				Keys("foo:*").
				Return(redis.NewStringSliceResult(
					[]string{"foo:bar", "foo:biz"},
					errors.New("asdf"),
				))

			result, err := c.GetCacheKeys("foo:*")

			So(result, ShouldResemble, []string{"foo:bar", "foo:biz"})
			So(err.Error(), ShouldEqual, "asdf")
		})
	})
}

func Test_rcache_PushOneLeft(t *testing.T) {
	Convey("*rcache.PushOneLeft()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the error returned by Encoder.Encode()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(
					nil,
					errors.New("baz"),
				)

			err := c.PushOneLeft("foo", "bar")

			So(err.Error(), ShouldEqual, "baz")
		})

		Convey("Returns the error returned by Cmdable.LPush().Error()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(types.String("baz"), nil)

			redisMock.EXPECT().
				LPush("foo", "baz").
				Return(redis.NewIntResult(
					1,
					errors.New("qux"),
				))

			err := c.PushOneLeft("foo", "bar")

			So(err.Error(), ShouldEqual, "qux")
		})

		Convey("Returns the error returned by Cmdable.Expire()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(types.String("biz"), nil)

			redisMock.EXPECT().
				LPush("foo", "biz").
				Return(redis.NewIntResult(
					1,
					nil,
				))

			redisMock.EXPECT().
				Expire("foo", time.Minute).
				Return(redis.NewBoolResult(
					true,
					errors.New("quux"),
				))

			err := c.PushOneLeft("foo", "bar")

			So(err.Error(), ShouldEqual, "quux")
		})
	})
}

func Test_rcache_PushOneRight(t *testing.T) {
	Convey("*rcache.PushOneRight()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the error returned by Encoder.Encode()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(
					nil,
					errors.New("baz"),
				)

			err := c.PushOneRight("foo", "bar")

			So(err.Error(), ShouldEqual, "baz")
		})

		Convey("Returns the error returned by Cmdable.RPush().Error()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(types.String("baz"), nil)

			redisMock.EXPECT().
				RPush("foo", "baz").
				Return(redis.NewIntResult(
					1,
					errors.New("qux"),
				))

			err := c.PushOneRight("foo", "bar")

			So(err.Error(), ShouldEqual, "qux")
		})

		Convey("Returns the error returned by Cmdable.Expire()", func() {
			encoderMock.EXPECT().
				Encode("bar").
				Return(types.String("biz"), nil)

			redisMock.EXPECT().
				RPush("foo", "biz").
				Return(redis.NewIntResult(
					1,
					nil,
				))

			redisMock.EXPECT().
				Expire("foo", time.Minute).
				Return(redis.NewBoolResult(
					true,
					errors.New("quux"),
				))

			err := c.PushOneRight("foo", "bar")

			So(err.Error(), ShouldEqual, "quux")
		})
	})
}

func Test_rcache_GetListOne(t *testing.T) {
	Convey("*rcache.GetListOne()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns the cache.ErrCacheMiss error returned by Cmdable.LIndex()", func() {
			redisMock.EXPECT().
				LIndex("foo", int64(1)).
				Return(redis.NewStringResult(
					"",
					errors.New(ErrNilResponse),
				))

			err := c.GetListOne("foo", 1, nil)

			So(err, ShouldResemble, cache.ErrCacheMiss)
		})

		Convey("Returns the error returned by Cmdable.LIndex().Error()", func() {
			redisMock.EXPECT().
				LIndex("foo", int64(1)).
				Return(redis.NewStringResult(
					"",
					errors.New("baz"),
				))

			err := c.GetListOne("foo", 1, nil)

			So(err, ShouldResemble, errors.New("baz"))
		})

		Convey("Returns the error returned by Encoder.Decode()", func() {
			redisMock.EXPECT().
				LIndex("foo", int64(1)).
				Return(redis.NewStringResult(
					"bar",
					nil,
				))

			encoderMock.EXPECT().
				Decode("bar", nil).
				Return(errors.New("baz"))

			err := c.GetListOne("foo", 1, nil)

			So(err, ShouldResemble, errors.New("baz"))
		})
	})
}

func Test_rcache_GetListRange(t *testing.T) {
	Convey("*rcache.GetListRange()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns error when dest type is not a slice", func() {
			So(c.GetListRange("foo", 1, 2, "foo"), ShouldResemble, cache.ErrDestNotPtrSlice)
		})

		Convey("Returns error when dest type is not a pointer", func() {
			So(c.GetListRange("foo", 1, 2, []string{}), ShouldResemble, cache.ErrDestNotPtrSlice)
		})

		Convey("Returns ErrNilResponse returned by c.LRange().Result()", func() {
			redisMock.EXPECT().
				LRange("foo", int64(1), int64(2)).
				Return(
					redis.NewStringSliceResult(
						[]string{},
						errors.New(ErrNilResponse),
					))

			So(c.GetListRange("foo", 1, 2, &[]*string{}), ShouldResemble, cache.ErrCacheMiss)
		})

		Convey("Returns error returned by c.LRange().Result()", func() {
			redisMock.EXPECT().
				LRange("foo", int64(1), int64(2)).
				Return(
					redis.NewStringSliceResult(
						[]string{},
						errors.New("foo"),
					))

			So(c.GetListRange("foo", 1, 2, &[]*string{}), ShouldResemble, errors.New("foo"))
		})

		Convey("Returns the error returned by Encoder.Decode()", func() {
			redisMock.EXPECT().
				LRange("foo", int64(1), int64(2)).
				Return(
					redis.NewStringSliceResult(
						[]string{"foo", "bar"},
						nil,
					))

			encoderMock.EXPECT().
				Decode("foo", gomock.Any()).
				Return(errors.New("foo"))

			So(c.GetListRange("foo", 1, 2, &[]*string{}), ShouldResemble, errors.New("foo"))
		})

		Convey("Decodes the list into dest", func() {
			redisMock.EXPECT().
				LRange("foo", int64(1), int64(2)).
				Return(
					redis.NewStringSliceResult(
						[]string{"foo", "bar"},
						nil,
					))

			encoderMock.EXPECT().
				Decode("foo", gomock.Any()).
				Return(nil)

			encoderMock.EXPECT().
				Decode("bar", gomock.Any()).
				Return(nil)

			dest := []*string{}
			err := c.GetListRange("foo", 1, 2, &dest)

			So(err, ShouldBeNil)
			So(dest, ShouldResemble, []*string{nil, nil})
		})
	})
}

func Test_rcache_TrimList(t *testing.T) {
	Convey("*rcache.TrimList()", t, func() {
		mockCtrl := gomock.NewController(t)
		redisMock := mocks.NewMockCmdable(mockCtrl)
		encoderMock := mocks.NewMockEncoder(mockCtrl)

		defer mockCtrl.Finish()

		c := &rcache{
			Cmdable:    redisMock,
			Encoder:    encoderMock,
			Expiration: time.Minute,
		}

		Convey("Returns error returned by Cmdable.LTrim().Err()", func() {
			redisMock.EXPECT().
				LTrim("foo", int64(1), int64(2)).
				Return(redis.NewStatusResult(
					"OK",
					errors.New("bar"),
				))

			So(c.TrimList("foo", 1, 2), ShouldResemble, errors.New("bar"))
		})
	})
}
