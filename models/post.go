package models

import (
	"errors"

	"github.com/pjcalvo/gin-boilerplate/db"
	"github.com/pjcalvo/gin-boilerplate/forms"
)

//Post ...
type Post struct {
	ID        int64    `db:"id, primarykey, autoincrement" json:"id"`
	UserID    int64    `db:"user_id" json:"-"`
	Title     string   `db:"title" json:"title"`
	Content   string   `db:"content" json:"content"`
	UpdatedAt int64    `db:"updated_at" json:"updated_at"`
	CreatedAt int64    `db:"created_at" json:"created_at"`
	User      *JSONRaw `db:"user" json:"user"`
}

//PostModel ...
type PostModel struct{}

//Create ...
func (m PostModel) Create(userID int64, form forms.PostForm) (postID int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.post(user_id, title, content) VALUES($1, $2, $3) RETURNING id", userID, form.Title, form.Content).Scan(&postID)
	return postID, err
}

//One ...
func (m PostModel) One(userID, id int64) (post Post, err error) {
	err = db.GetDB().SelectOne(&post, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.post a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
	return post, err
}

//All ...
func (m PostModel) All(userID int64) (posts []DataList, err error) {
	_, err = db.GetDB().Select(&posts, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(a.id) AS total FROM public.post AS a WHERE a.user_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.post a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC) d", userID)
	return posts, err
}

//Update ...
func (m PostModel) Update(userID int64, id int64, form forms.PostForm) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Post not found")
	}

	_, err = db.GetDB().Exec("UPDATE public.post SET title=$2, content=$3 WHERE id=$1", id, form.Title, form.Content)

	return err
}

//Delete ...
func (m PostModel) Delete(userID, id int64) (err error) {
	_, err = m.One(userID, id)

	if err != nil {
		return errors.New("Post not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM public.post WHERE id=$1", id)

	return err
}
