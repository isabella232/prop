package backend

import (
	"encoding/json"

	"github.com/xo/dburl"
)

type Backend interface {
	BackendExport() (PropertyCollection, bool, error)
	BackendImport(clear bool) (PropertyCollection, bool, error)
	BackendClear() (bool, error)
	Del(key string) (bool, error)
	Exists(key string) (bool, error)
	NamespaceExists(namespace string) (bool, error)
	NamespaceClear(namespace string) (bool, error)
	Get(key string, defaultValue string) (string, error)
	GetAll() (map[string]string, error)
	GetAllByPrefix(prefix string) (map[string]string, error)
	Set(key string, value string) (bool, error)
	Lindex(key string, index int) (string, error)
	Lismember(key string, element string) (bool, error)
	Llen(key string) (int, error)
	Lrange(key string) ([]string, error)
	Lrangefrom(key string, start int) ([]string, error)
	Lrangefromto(key string, start int, stop int) ([]string, error)
	Lrem(key string, countToRemove int, element string) (int, error)
	Rpush(key string, newElements ...string) (int, error)
	Sadd(key string, newMembers ...string) (int, error)
	Sismember(key string, member string) (bool, error)
	Smembers(key string) (map[string]bool, error)
	Srem(key string, membersToRemove ...string) (int, error)
}

func ConstructBackend(url string, namespace string) (Backend, error) {
	u, err := dburl.Parse(url)
	if err != nil {
		return NewUnimplementedBackend()
	}

	if namespace == "" {
		namespace = u.Query().Get("namespace")
	}

	if u.Scheme == "file" {
		return NewFileBackend(namespace, u)
	}

	return NewUnimplementedBackend()
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
