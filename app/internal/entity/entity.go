package entity

type CreateAtricleRequest struct {
	Author  int    `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateAuthorRequest struct {
	Name        string `json:"author_name"`
	IsModerator bool   `json:"is_moderator"`
}

type CreateCommentRequest struct {
	Article     int    `json:"article_id"`
	Commentator int    `json:"commentator_id"`
	Comment     string `json:"comment"`
}

type GetAtricleResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
