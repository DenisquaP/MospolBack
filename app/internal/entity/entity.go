package entity

type CreateAtricleRequest struct {
	Author  int    `json:"article"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateAuthorRequest struct {
	Name        string `json:"author_name"`
	IsModerator bool   `json:"is_moderator"`
}

type ErrorResponse struct {
	Error error `json:"error"`
}
