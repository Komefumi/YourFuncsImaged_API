package request

// JSONRequest will be the container type that holds a data field, for the particulars of any Request's main contents
type JSONRequest struct {
	data interface{}
}

// RegistrationRequest will holds the payload for registration
type RegistrationRequest struct {
	email string
	password string
}