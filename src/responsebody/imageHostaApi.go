package responsebody

type Chain struct {
	Image string `json:"Image"`
}
type resource struct {
	Chain Chain `json:"chain"`
}

type file struct {
	Resource resource `json:"resource"`
}
type image struct {
	File file `json:"file"`
}
type ImageHostAPI struct {
	Status_code int16 `json:"status_code"`
	Image       image `json:"image"`
}
