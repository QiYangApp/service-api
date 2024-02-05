package helpers

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

//作者：一位不愿意透露姓名的杨先生
//链接：https://juejin.cn/post/6981348822353985566
//来源：稀土掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

type Limiters struct {
	limiters map[string]*Limiter
	lock     sync.Mutex
}

type Limiter struct {
	limiter *rate.Limiter
	lastGet time.Time //上一次获取token的时间
	key     string
}

var LimiterClient = &Limiters{
	limiters: make(map[string]*Limiter),
	lock:     sync.Mutex{},
}

func NewLimiter(r rate.Limit, b int, key string) *Limiter {
	LimiterClient.lock.Lock()

	defer LimiterClient.lock.Unlock()

	limiter, ok := LimiterClient.limiters[key]
	if ok {
		return limiter
	}

	l := &Limiter{
		limiter: rate.NewLimiter(r, b),
		lastGet: time.Now(),
		key:     key,
	}

	LimiterClient.limiters[key] = l

	return l
}

func (l *Limiter) Allow() bool {

	l.lastGet = time.Now()

	return l.limiter.Allow()

}

func (ls *Limiters) GetLimiter(key string) *Limiter {
	limiter, ok := LimiterClient.limiters[key]
	if ok {
		return limiter

		panic("get limit not exist, name: " + key)
	}

	return limiter
}

// 清除过期的限流器
func (ls *Limiters) clearLimiter() {

	for {

		time.Sleep(1 * time.Minute)
		ls.lock.Lock()
		for i, i2 := range ls.limiters {

			//超过1分钟
			if time.Now().Unix()-i2.lastGet.Unix() > 60 {
				delete(ls.limiters, i)
			}

		}
		ls.lock.Unlock()
	}

}
