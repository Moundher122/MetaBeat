package dto

type CreateLinkDto struct {
	Title string
	URL   string
}
type LinkDto struct {
	ID    uint
	Title string
	Icon  string
	URL   string
	Type  string
}