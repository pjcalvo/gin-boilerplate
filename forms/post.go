package forms

//PostForm ...
type PostForm struct {
	Title   string `form:"title" json:"title" binding:"required,max=100"`
	Content string `form:"content" json:"content" binding:"required,max=1000"`
}

//CommentForm ...
type CommentForm struct {
	Name    string `form:"name" json:"name" binding:"required,max=100"`
	Comment string `form:"comment" json:"comment" binding:"required,max=1000"`
}
