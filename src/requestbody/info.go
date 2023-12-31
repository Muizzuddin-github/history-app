package requestbody

type Info struct {
	Title       string `form:"title" validate:"required" json:"title"`
	Description string `form:"description" validate:"required" json:"description"`
	Image       string `form:"image"`
	Created_at  string
}

type InfoUpdateNoImage struct {
	Title       string `bson:"title" form:"title" validate:"required"`
	Description string `bson:"description" form:"description" validate:"required"`
}

type InfoUpdateWithImage struct {
	Title       string `bson:"title" form:"title" validate:"required"`
	Description string `bson:"description" form:"description" validate:"required"`
	ImageUrl    string `bson:"imageUrl"`
}