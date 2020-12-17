package web

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wangtuanjie/ip17mon"

	"github.com/laojianzi/mdclubgo/conf"
)

// IP
func IP(ctx echo.Context) string {
	headerKeys := []string{"HTTP_X_FORWARDED_FOR", "HTTP_CLIENT_IP", "REMOTE_ADDR"}
	for _, key := range headerKeys {
		v := ctx.Request().Header.Get(key)
		if v != "" {
			return v
		}
	}

	return "0.0.0.0"
}

// Location
func Location(ctx echo.Context, ip string) string {
	if ip == "" {
		ip = IP(ctx)
	}

	loc, err := ip17mon.Find(ip)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(strings.Join([]string{loc.Country, loc.Region, loc.City}, " "))
}

// IPSign
func IPSign(ctx echo.Context) string {
	return strings.Join(regexp.MustCompilePOSIX("[a-z0-9A-Z]*?").FindAllString(IP(ctx), -1), ".")
}

// HostPath
func HostPath(ctx echo.Context) string {
	return fmt.Sprintf("%s://%s", ctx.Scheme(), ctx.Request().Host)
}

// StaticPath
func StaticPath(ctx echo.Context, defaultStaticPath string) string {
	url := conf.Server.SiteStaticURL
	if url == "" {
		url = fmt.Sprintf("%s/%s", HostPath(ctx), defaultStaticPath)
	}

	if url != "" && url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	return url
}

// SupportWebp
func SupportWebp(ctx echo.Context) bool {
	accept := ctx.Request().Header.Get("HTTP_ACCEPT")
	if accept == "" {
		return false
	}

	return strings.Contains(accept, "image/webp")
}
