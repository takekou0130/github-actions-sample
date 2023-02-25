package services

import (
	"github.com/takekou0130/github-actions-sample/myapi/apperrors"
	"github.com/takekou0130/github-actions-sample/myapi/models"
	"github.com/takekou0130/github-actions-sample/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
