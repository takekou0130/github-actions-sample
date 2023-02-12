package repositories

import (
	"database/sql"

	"github.com/takekou0130/github-actions-sample/myapi/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values (?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	var newComment models.Comment
	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	newComment.CommentID = int(id)
	newComment.ArticleID = comment.ArticleID
	newComment.Message = comment.Message
	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdAt sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdAt)
		if err != nil {
			return nil, err
		}

		if createdAt.Valid {
			comment.CreatedAt = createdAt.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
