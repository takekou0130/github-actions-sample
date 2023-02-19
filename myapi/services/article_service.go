package services

import (
	"github.com/takekou0130/github-actions-sample/myapi/models"
	"github.com/takekou0130/github-actions-sample/myapi/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArtilceDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	comments, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, comments...)
	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articles, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	article.NiceNum += 1
	return article, nil
}
