package service

import (
	"restful/model"
    "restful/dao/newsdao"
)

type NewsService struct {
	newsDao *newsdao.NewsDao
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

//返回ID新闻信息
func (service *NewsService) GetNewsById(newsId int64) (news model.News) {
	return newsDao.GetNewsById(newsId)
}

//返回新闻总数
func (service *NewsService) GetNewsCount() int {
    return newsDao.GetNewsCount()
}

//返回新闻列表
func (service *NewsService) GetNewsByPage(title string, page int, pageSize int) interface{} {
    count := newsDao.GetNewsCount()
    page = SetPage(page, pageSize, count)
    offset := SetOffset(page, pageSize)
    newsList := newsDao.GetNewsByPage("", offset, pageSize)
    
    data := make(map[string]interface{}, 2)
    data["count"] = count
    data["news_list"] = newsList
    data["page"] = page
    return data
}

//添加新闻
func (service *NewsService) InsertNews(news *model.News) {
    newsDao.InsertNews(news)
}

//更新新闻
func (service *NewsService) UpdateNews(news *model.News) (int64) {    
    return newsDao.UpdateNews(news)
}

//删除新闻
func (service *NewsService) DeleteNews(newsId int64) (int64) {
    return newsDao.DeleteNews(newsId)
}