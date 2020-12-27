package controllers

import (
	req "github.com/Komefumi/YourFuncsImaged_API/app/request"

	"github.com/revel/revel"
)

type EmptyResponse struct {
	message string
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Register() revel.Result {
	var jsonData map[string]req.JSONRequest
	response := JSONResponse{}
	err := error(nil)
	response.Success = err == nil
	response.Data = EmptyResponse{ message: "Yeah, yeah" }

	return c.RenderJSON(response)
}