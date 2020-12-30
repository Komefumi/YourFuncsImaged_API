package util

import "github.com/revel/revel"

func GetBodyString(c *revel.Controller) string {
	return string(c.Params.JSON)
}