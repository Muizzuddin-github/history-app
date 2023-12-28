package responsebody

type Msg struct {
	Message string `json:"message"`
}

type Err struct {
	Errors []string `json:"errors"`
}

type Data struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type InsertDocument struct {
	Message    string `json:"message"`
	InsertedID string `json:"insertedID"`
}