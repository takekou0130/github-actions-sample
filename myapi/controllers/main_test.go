package controllers_test

import (
	"testing"

	"github.com/takekou0130/github-actions-sample/myapi/controllers"
	"github.com/takekou0130/github-actions-sample/myapi/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
