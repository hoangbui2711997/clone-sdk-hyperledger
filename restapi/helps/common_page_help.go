package helps

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPosCountPage(c *gin.Context) (pos, count int64) {
	posStr, ok := c.GetQuery("pos")
	if !ok {
		pos = 0
	} else {
		pos, _ = strconv.ParseInt(posStr, 10, 64)
	}

	countStr, ok := c.GetQuery("count")
	if !ok {
		count = 10
	} else {
		count, _ = strconv.ParseInt(countStr, 10, 64)
	}
	return
}

func GetPathInt(c *gin.Context, name string) (int, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.Atoi(val)
}
func GetPathInt64(c *gin.Context, name string) (int64, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.ParseInt(val, 10, 64)
}
func GetPathBool(c *gin.Context, name string) (bool, error) {
	val := c.Params.ByName(name)
	if val == "" {
		return false, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.ParseBool(val)
}
func GetQueryInt(c *gin.Context, name string) (int, error) {
	val := c.Query(name)
	if val == "" {
		return 0, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.Atoi(val)
}
func GetQueryBool(c *gin.Context, name string) (bool, error) {
	val := c.Query(name)
	if val == "" {
		return false, errors.New(name + " path parameter value is empty or not specified")
	}
	return strconv.ParseBool(val)
}
func GetQueryInt64(c *gin.Context, name string) (int64, error) {
	val := c.Query(name)
	if val == "" {
		return 0, errors.New(name + " query parameter value is empty or not specified")
	}
	return strconv.ParseInt(val, 10, 64)
}
