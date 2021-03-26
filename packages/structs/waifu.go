package structs

type AnimeWaifu struct {
	Id string `json:"id`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Title string `json:"title"`
	HairColor string `json:"hair_color"`
	Age int `json:"age" validate:"required"`
}

func NewAnimeWaifu() *AnimeWaifu {
	return &AnimeWaifu{}
}
