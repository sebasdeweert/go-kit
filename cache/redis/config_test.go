package redis

import (
	"testing"

	"github.com/go-redis/redis"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Config_GetRingOptions(t *testing.T) {
	Convey("*Config.GetRingOptions()", t, func() {
		Convey("Returns a redis ring options instance based on the configuration fields", func() {
			c := &Config{
				Addresses: []string{"foo", "bar"},
				Password:  "biz",
				DB:        1,
			}

			So(c.GetRingOptions(), ShouldResemble, &redis.RingOptions{
				DB: 1,
				Addrs: map[string]string{
					"server1": "foo",
					"server2": "bar",
				},
				Password: "biz",
			})
		})
	})
}
