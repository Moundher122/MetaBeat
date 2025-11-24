package dto

type RegisterDto struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password"   validate:"required,min=8"`
	Bio      string `json:"bio"`
	BgSong   string `json:"bg_song"`
}

type UserDto struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	BgSong   string `json:"bg_song"`
}
