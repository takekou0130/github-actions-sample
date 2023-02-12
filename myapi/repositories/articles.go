package repositories

import (
	"database/sql"

	"github.com/takekou0130/github-actions-sample/myapi/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)

	if err != nil {
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	var newArticle models.Article
	newArticle.Title = article.Title
	newArticle.Contents = article.Contents
	newArticle.UserName = article.UserName
	newArticle.ID = int(id)
	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`
	const numPerPage = 5
	rows, err := db.Query(sqlStr, numPerPage, (page-1)*numPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return nil, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArtilceDetail(db *sql.DB, artilceID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	result := db.QueryRow(sqlStr, artilceID)
	if err := result.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdAt sql.NullTime
	err := result.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdAt)
	if err != nil {
		return models.Article{}, err
	}

	if createdAt.Valid {
		article.CreatedAt = createdAt.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `
		update articles
		set nice = ?
		where article_id = ?;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	result := tx.QueryRow(sqlGetNice, articleID)
	if err := result.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var niceNum int
	err = result.Scan(&niceNum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
