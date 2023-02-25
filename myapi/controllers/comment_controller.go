package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/takekou0130/github-actions-sample/myapi/apperrors"
	"github.com/takekou0130/github-actions-sample/myapi/controllers/services"
	"github.com/takekou0130/github-actions-sample/myapi/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "fail to decode req body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
