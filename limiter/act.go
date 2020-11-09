package limiter

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/log"
)

// ActLimit obtaining the remaining executable times
func ActLimit(ctx echo.Context, id, action string, maxCount, period int) int {
	reqTimeSec, err := strconv.ParseInt(ctx.Request().Header.Get("REQUEST_TIME"), 10, 64)
	reqTime := time.Now()
	if err == nil {
		reqTime = time.Unix(reqTimeSec, 0)
	}

	reqTimeUnix := int(reqTime.Unix())
	ttl := reqTimeUnix/period*period + period - reqTimeUnix
	key := fmt.Sprintf("mdclubgo_throttle_%s_%s", action, id)
	currentCount, err := strconv.Atoi(cache.Get(key, "0"))
	currentCount++
	if err != nil || currentCount > maxCount {
		if err != nil {
			log.Error(fmt.Errorf("act limit get cache: %w", err).Error())
		}

		return 0
	}

	err = cache.Set(key, fmt.Sprintf("%d", currentCount), ttl)
	if err != nil {
		log.Error(fmt.Errorf("act limit set cache: %w", err).Error())
	}

	return maxCount - currentCount + 1
}
