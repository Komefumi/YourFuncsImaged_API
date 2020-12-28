package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/tidwall/gjson"
)

type EmptyResponse struct {
	message string
}

type App struct {
	*revel.Controller
}

func (c App) getBodyString() string {
	return string(c.Params.JSON)
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Register() revel.Result {
	body := c.getBodyString()
	fmt.Println(body)
	email := gjson.Get(body, "data.email").String()
	password := gjson.Get(body, "data.password").String()
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println("This has to work")
	response := CreateJSONResponse(error(nil), "Dummy response, check!", DummyRegisterResponse{ Email: email, Password: password })
	return c.RenderJSON(response)
}