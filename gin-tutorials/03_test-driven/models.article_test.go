package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()
	if len(alist) != len(articleList) {
		t.Fail()
	}
	for i, v := range alist {
		if v.ID != articleList[i].ID || v.Title != articleList[i].Title || v.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}

}

func TestCreateArticle(t *testing.T) {
	length := len(getAllArticles())
	a, err := createNewArticle("newTitle", "newContent")

	newLength := len(getAllArticles())
	if err != nil || length != newLength-1 || a.Title != "newTitle" || a.Content != "newContent" {
		t.Fail()
	}
}

func TestArticleCreationAuthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()

	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	r.POST("/article/create", createArticle)

	articlePayload := getArticlePOSTPayload()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Submission Successful</title>") < 0 {
		t.Fail()
	}
	restoreLists()
}

func getArticlePOSTPayload() string {
	params := url.Values{}
	params.Add("title", "Test Article Title")
	params.Add("content", "Test Article Content")

	return params.Encode()
}
