package core

import (
	"github.com/gomodule/redigo/redis"
	"github.com/soxft/busuanzi/config"
	"github.com/soxft/busuanzi/library/tool"
	"github.com/soxft/busuanzi/process/redisutil"
)

// Count
// @description return and count the number of users in the redis
func Count(host string, path string, userIdentity string) (int, int, int, int) {
	_redis := redisutil.Pool.Get()
	defer func(_redis redis.Conn) {
		_ = _redis.Close()
	}(_redis)

	// encode
	var pathUnique = tool.Md5(host + "&" + path)
	var siteUnique = tool.Md5(host)

	redisPrefix := config.Redis.Prefix
	siteUvKey := redisPrefix + ":site_uv:" + siteUnique
	pageUvKey := redisPrefix + ":page_uv:" + siteUnique + ":" + pathUnique

	sitePvKey := redisPrefix + ":site_pv:" + siteUnique
	pagePvKey := redisPrefix + ":page_pv:" + siteUnique

	// count sitePv ans pagePv
	sitePv, _ := redis.Int(_redis.Do("INCR", sitePvKey))
	pagePv, _ := redis.Int(_redis.Do("ZINCRBY", pagePvKey, 1, pathUnique))
	_, _ = _redis.Do("SADD", siteUvKey, userIdentity)
	_, _ = _redis.Do("SADD", pageUvKey, userIdentity)

	siteUv, _ := redis.Int(_redis.Do("SCARD", siteUvKey))
	pageUv, _ := redis.Int(_redis.Do("SCARD", pageUvKey))

	if config.Bsz.Expire > 0 {
		go setExpire(sitePvKey, siteUvKey, pagePvKey, pageUvKey)
	}

	return sitePv, siteUv, pagePv, pageUv
}

func setExpire(key ...string) {
	var _redis = redisutil.Pool.Get()
	defer func(_redis redis.Conn) {
		_ = _redis.Close()
	}(_redis)
	// multi-set expire
	_, _ = _redis.Do("MULTI")
	for _, k := range key {
		_, _ = _redis.Do("EXPIRE", k, config.Bsz.Expire)
	}
	_, _ = _redis.Do("EXEC")
}
