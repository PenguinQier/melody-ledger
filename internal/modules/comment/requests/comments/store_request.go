package comments

type StoreRequest struct {
	Content   string `json:"content" binding:"required,min=1,max=1000"`
	ArticleID uint   `json:"article_id" binding:"required"`
}
