package views

//Response struct created
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

//PostRequest struct created to decode incoming data from user.
type PostRequest struct {
	Name string `json:"name"`
	ToDo string `json:"todo"`
}
