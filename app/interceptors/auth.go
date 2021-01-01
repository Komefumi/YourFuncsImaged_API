package interceptors

import (
	"fmt"

	util "github.com/Komefumi/YourFuncsImaged_API/app/util"
	"github.com/revel/revel"
)

// AuthInterceptor to intercept authentication routes
func AuthInterceptor(c *revel.Controller) revel.Result {
	tokenString := c.Request.Header.Get("Authorization")
	error := util.ValidateToken(tokenString)
	fmt.Println(error)
	if error != nil {
		fmt.Println("Currently error is ", error)
		// return c.Redirect("/")
		return c.RenderJSON(struct{}{})
	}
	return nil
}