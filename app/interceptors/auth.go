package interceptors

import (
	"github.com/Komefumi/YourFuncsImaged_API/app/controllers"
	util "github.com/Komefumi/YourFuncsImaged_API/app/util"
	"github.com/revel/revel"
	"github.com/tidwall/gjson"
)

// AuthInterceptor to intercept authentication routes
func AuthInterceptor(c *revel.Controller) revel.Result {
	body := util.GetBodyString(c)
  if error := util.ValidateToken(gjson.Get(body, "data.token").String()); error != nil {
    return c.Redirect(controllers.App.Index)
	}
	return nil
}