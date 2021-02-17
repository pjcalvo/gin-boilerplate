package models

import (
	"errors"

	"github.com/pjcalvo/gin-boilerplate/db"
	"github.com/pjcalvo/gin-boilerplate/forms"
)

//Comment ...
type Comment struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	PostID    string `db:"post_id" json:"post_id"`
	Name      string `db:"name" json:"name"`
	Comment   string `db:"comment" json:"comment"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
}

//CommentModel ...
type CommentModel struct{}

//Create ...
func (m CommentModel) Create(postID int64, form forms.CommentForm) (commendID int64, err error) {
	err = db.GetDB().QueryRow("INSERT INTO public.comment(post_id, name, comment) VALUES($1, $2, $3) RETURNING id", postID, form.Name, form.Comment).Scan(&commendID)
	return commendID, err
}

//One ...
func (m CommentModel) One(postID, id int64) (comment Comment, err error) {
	err = db.GetDB().SelectOne(&comment, "SELECT id, post_id, name, comment, updated_at, created_at FROM public.comment a WHERE post_id=$1 AND id=$2 LIMIT 1", postID, id)
	return comment, err
}

//All ...
func (m CommentModel) All(userID int64) (comments []DataList, err error) {
	_, err = db.GetDB().Select(&comments, "SELECT COALESCE(array_to_json(array_agg(row_to_json(d))), '[]') AS data, (SELECT row_to_json(n) FROM ( SELECT count(id) AS total FROM public.comment AS a WHERE post_id=$1 LIMIT 1 ) n ) AS meta FROM ( SELECT id, name, comment, updated_at, created_at from public.comment WHERE post_id=$1 ORDER BY id DESC) d", userID)
	return comments, err
}

//Update ...
func (m CommentModel) Update(postID int64, id int64, form forms.CommentForm) (err error) {
	_, err = m.One(postID, id)

	if err != nil {
		return errors.New("Comment not found")
	}

	_, err = db.GetDB().Exec("UPDATE public.comment SET name=$2, comment=$3 WHERE id=$1", id, form.Name, form.Comment)

	return err
}

//Delete ...
func (m CommentModel) Delete(postID, id int64) (err error) {
	_, err = m.One(postID, id)

	if err != nil {
		return errors.New("Comment not found")
	}

	_, err = db.GetDB().Exec("DELETE FROM public.comment WHERE id=$1", id)

	return err
}
