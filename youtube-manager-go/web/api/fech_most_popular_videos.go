package api

import(
	"github.com/labstack/echo"
	"github.com/valuala/fasthttp"
)

func FetchMostPopularVideos()
echo.HandlerFunc{
	return func(c echo.context)
error{
	return c.JSON(fasthttp.StatusOK, "Most Popular")
}
}