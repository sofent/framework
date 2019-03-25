package cache

import (
    "fmt"
    "time"
)

type RateLimit struct {
    cache cacher
}

const RATE_LIMITE_CACHE_KEY = "TOTOVAL_RATE_LIMIT_%s"
const RATE_LIMITE_TIMER_CACHE_KEY = "TOTOVAL_RATE_LIMIT_%s_TIMER"

func rateLimitCacheKey(key string) string {
    return fmt.Sprintf(RATE_LIMITE_CACHE_KEY, key)
}
func rateLimitTimerCacheKey(key string) string {
    return fmt.Sprintf(RATE_LIMITE_TIMER_CACHE_KEY, key)
}

func (rl *RateLimit) TooManyAttempts(key string, maxAttempts int64) bool {
    if rl.Attempts(key) >= maxAttempts {
        if rl.cache.Has(rateLimitTimerCacheKey(key)) {
            return true
        }

        rl.ResetAttempts(key)
    }
    return false
}
func (rl *RateLimit) Hit(key string, decayMinutes int) (incremented int64) {
    expiredAt := time.Now().Add(time.Duration(decayMinutes) * time.Minute)
    rl.cache.Add(rateLimitTimerCacheKey(key), expiredAt.Unix(), expiredAt)

    added := rl.cache.Add(rateLimitCacheKey(key), int64(0), expiredAt)
    incremented, success := rl.cache.Increment(rateLimitCacheKey(key), 1)

    if !added && (success && incremented == 1) {
        rl.cache.Put(rateLimitCacheKey(key), int64(1), expiredAt)
    }

    return incremented
}
func (rl *RateLimit) Attempts(key string) int64 {
    return rl.cache.Get(rateLimitCacheKey(key), int64(0)).(int64)
}
func (rl *RateLimit) ResetAttempts(key string) bool {
    return rl.cache.Forget(rateLimitCacheKey(key))
}
func (rl *RateLimit) RetriesLeft(key string, maxAttempts int64) int64 {
    attempts := rl.Attempts(key)
    return maxAttempts - attempts
}
func (rl *RateLimit) Clear(key string) {
    rl.ResetAttempts(key)
    rl.cache.Forget(rateLimitTimerCacheKey(key))
}
func (rl *RateLimit) AvailableIn(key string) time.Duration {
    expiredAtUnix := rl.cache.Get(rateLimitTimerCacheKey(key)).(int64)
    expiredAt := time.Unix(expiredAtUnix, 0)
    return expiredAt.Sub(time.Now())
}