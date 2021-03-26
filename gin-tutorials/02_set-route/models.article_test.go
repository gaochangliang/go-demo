package main

import "testing"

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
