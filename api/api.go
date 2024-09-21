package api

// type TodoParams struct {
// 	ID string
// }

type TodoResponse struct {
	// success code
	Code string `json:"code"`

	Data interface{} `json:"data"`
}

type Error struct {
	// error code
	Code int

	Message string
}
