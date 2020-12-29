package controllers

// JSONResponse to make a JSON response
type JSONResponse struct {
	Message string `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}


// DummyRegisterResponse is a dummy
type DummyRegisterResponse struct {
	Email string ` json:"email" xml:"email" `
  Password string ` json:"password" xml:"password" `
}

// RegisterResponse gives response for registration
type RegisterResponse struct {
	ID uint `json:"id"`
}

// LoginSuccessResponse to respond with Auth data
type LoginSuccessResponse struct {
	Token string `json:"token"`
}

// CreateJSONResponse takes an error, message, and data, and returns a map that can be RenderJSON'd
func CreateJSONResponse(providedError error, message string, data interface{}) map[string]interface{} {
	returnable := make(map[string]interface{})
	returnable["error"] = error(nil)
	returnable["data"] = data
	returnable["message"] = message
	return returnable
}
