package cache

// Cache is an interface that allows cache handling.
type Cache interface {
	GetOne(index string, obj interface{}) error
	SetOne(index string, obj interface{}) error
	MultiGet(objs map[string]interface{}) (map[string]error, error)
	MultiSet(objs map[string]interface{}) error
	CheckHealth() error
	Delete(indexes ...string) error
	GetSetMembers(index string) ([]string, error)
	GetSetMembersByRank(index string, offset int64, count int64, direction Direction) ([]string, error)
	GetSetMembersByScore(index, min, max string, offset, count int64, direction Direction) ([]string, error)
	AddToSet(index string, members []interface{}) error
	GetCacheKeys(pattern string) ([]string, error)
	GetListRange(index string, start int64, stop int64, obj interface{}) error
	GetListOne(index string, offset int64, obj interface{}) error
	PushOneLeft(index string, obj interface{}) error
}
