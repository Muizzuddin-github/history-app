package requestbody

type Info struct {
	Title       string `form:"title" validate:"required" json:"title"`
	Description string `form:"description" validate:"required" json:"description"`
	Image       string `form:"image"`
	Created_at  string
}

type InfoUpdateNoImage struct {
	Title       string `bson:"title" json:"title" validate:"required"`
	Description string `bson:"description" json:"description" validate:"required"`
}

type InfoImage struct {
	Image string `form:"image"`
}