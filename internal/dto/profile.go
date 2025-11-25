package dto
type ProfileDto struct {
	Username string
	Links	*[]LinkDto
	Theme	*ThemeDto
	Media    *[]string
	Bio      *string
	BgSong   *string
}

