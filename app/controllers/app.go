package controllers

import (
	"fmt"

	models "github.com/Komefumi/YourFuncsImaged_API/app/models"
	"github.com/Komefumi/YourFuncsImaged_API/app/util"
	"github.com/revel/revel"
	"github.com/tidwall/gjson"
)

// EmptyResponse struct
type EmptyResponse struct {
	message string
}

// App Controller
type App struct {
	*revel.Controller
}

func (c App) getBodyString() string {
	return string(c.Params.JSON)
}

// Index handler
func (c App) Index() revel.Result {
	return c.Render()
}

// Register handler
func (c App) Register() revel.Result {
	body := c.getBodyString()
	fmt.Println(body)
	email := gjson.Get(body, "data.email").String()
	password := gjson.Get(body, "data.password").String()
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println("This has to work")
	user := models.User{ Email: email }
	user.SetPassword(password)
	db := util.DBAccessorFunc()
	db.Create(&user)
	response := CreateJSONResponse(error(nil), "Success!", RegisterResponse{ ID: user.ID })
	return c.RenderJSON(response)
}

// Login handler
func (c App) Login() revel.Result {
	body := c.getBodyString()
	fmt.Println(body)
	email := gjson.Get(body, "data.email").String()
	password := gjson.Get(body, "data.password").String()
	foundUser, err := models.FindAuthUser(email, password)
	if err != nil {
		response := CreateJSONResponse(error(nil), err.Error(), struct{}{})
		return c.RenderJSON(response)
	}
	// foundUser.PasswordHash = nil
	token, tokenErr := util.CreateToken(uint64(foundUser.ID))
	if tokenErr != nil {
		response := CreateJSONResponse(error(nil), tokenErr.Error(), struct{}{})
		return c.RenderJSON(response)
	}
	response := CreateJSONResponse(error(nil), "Successfully authenticated!", LoginSuccessResponse{ Token: token })
	// println(email, "\n", password)
	return c.RenderJSON(response)
}