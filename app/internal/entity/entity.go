package entity

type CreateAtricleRequest struct {
	Author  int    `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateAuthorRequest struct {
	Email       string `json:"email"`
	Name        string `json:"author_name"`
	Password    string `json:"password"`
	IsModerator bool   `json:"is_moderator"`
}

type CreateCommentRequest struct {
	Article     int    `json:"article_id"`
	Commentator int    `json:"commentator_id"`
	Comment     string `json:"comment"`
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User        string `json:"user" db:"author_name"`
	IsModerator bool   `json:"is_moderator" db:"is_moderator"`
}

type GetAtricleResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type GetArticlesResponse struct {
	Articles []GetAtricleResponse `json:"articles"`
	LastPage int                  `json:"last_page"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type OkResponse struct {
	Message string `json:"message"`
}
