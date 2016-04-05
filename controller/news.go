package controller

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/coreos/go-semver/semver"
    "log"
    "strconv"
    "time"
	"restful/model"
)

func newsList(w rest.ResponseWriter, r *rest.Request) {
    version := r.Env["VERSION"].(*semver.Version)
    versionNumber := VersionToNumber(version.String())
    log.Println(versionNumber)
    //count := newsService.GetNewsCount()
    page, _:= strconv.Atoi(r.PathParam("page"))
    pageSize, _ := strconv.Atoi(r.PathParam("pageSize"))
    data := newsService.GetNewsByPage("", page, pageSize)
    rs := ShowJson(0, "", data)
    w.WriteJson(&rs)
}

func AddNews(w rest.ResponseWriter, r *rest.Request) {
    news := model.News{}
    r.DecodeJsonPayload(&news)
    newsService.InsertNews(&news)
    rs := ShowJson(0, "", news)
    w.WriteJson(&rs)
}

func UpdateNews(w rest.ResponseWriter, r *rest.Request) {
    news := model.News{}
    newsInfo := model.News{}
    r.DecodeJsonPayload(&news)
    
    newsInfo = newsService.GetNewsById(news.NewsId)
    log.Println(newsInfo)
    if newsInfo.NewsId == 0 {
        rs := ShowJson(1, "新闻不存在", "")
        w.WriteJson(&rs)
        return 
    }
    
    news.UpdateTime = time.Now()
    newsService.UpdateNews(&news)
    rs := ShowJson(0, "更新成功", news)
    w.WriteJson(&rs)
}

func NewsDetail(w rest.ResponseWriter, r *rest.Request) {
    newsId, _ := strconv.ParseInt(r.PathParam("id"), 10, 64)
    news := newsService.GetNewsById(newsId)
    rs := ShowJson(0, "", news)
    w.WriteJson(&rs)
}

func DeleteNews(w rest.ResponseWriter, r *rest.Request) {
    newsId, _ := strconv.ParseInt(r.PathParam("id"), 10, 64)
    newsInfo := newsService.GetNewsById(newsId)
    
    if newsInfo.NewsId == 0 {
        rs := ShowJson(1, "新闻不存在", "")
        w.WriteJson(&rs)
        return 
    }
    
    rows := newsService.DeleteNews(newsId)
    rs := ShowJson(0, "删除成功", rows)
    w.WriteJson(&rs)
}