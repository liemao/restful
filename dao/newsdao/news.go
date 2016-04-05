package newsdao

import (
    "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
	"restful/model"
    "restful/dao"
	"encoding/json"
	"fmt"
	"log"
)

const (
    NEWS_CACHE_KEY = "get:news:by:id:"
)

type NewsDao struct {
	db            *gorm.DB
    redisClients  *redis.Pool

}

func NewNewsDao() *NewsDao {
	return &NewsDao{
		db:          dao.DB,
        redisClients: dao.RedisClients,
	}
}

func (dao *NewsDao) GetNewsById(newsId int64) model.News {
    
    news := model.News{}
    redisClient := dao.redisClients.Get()
    defer redisClient.Close()
    
    key := NEWS_CACHE_KEY + fmt.Sprintf("%d", newsId)
    data, _ := redis.String(redisClient.Do("GET", key))
    if data != "" {
        if err := json.Unmarshal([]byte(data), &news); err != nil {
            log.Fatalln("news json to struct err", err)
        }
    } else {
        dao.db.First(&news, newsId)
        newsString, err := json.Marshal(news)

        if (err != nil) {
            log.Fatalln("news struct to json err", err)
        }
        
        if (news.NewsId > 0) {
            redisClient.Do("SET", key, newsString)
        }
    }
    
    return news
}

func (dao *NewsDao) GetNewsCount() int {
    var count int
    dao.db.Table("news").Count(&count)
    return count
}

func (dao *NewsDao) GetNewsByPage(title string, offset int, pageSize int) []model.News{
    news := []model.News{}
    
    if title != "" {
        dao.db.Where("title = ?", title);
    }
    
    dao.db.Offset(offset).Limit(pageSize).Find(&news)
    return news
}

func (dao *NewsDao) InsertNews(news *model.News) {
    dao.db.Create(&news)
    if news.NewsId > 0 {
        redisClient := dao.redisClients.Get()
        defer redisClient.Close()
        
        key := NEWS_CACHE_KEY + fmt.Sprintf("%d", news.NewsId)
        newsString, _ := json.Marshal(news)
        redisClient.Do("SET", key, newsString);
    }
}

func (dao *NewsDao) UpdateNews(news *model.News) (int64) {
    count := dao.db.Save(&news).RowsAffected
    if count > 0 {
        redisClient := dao.redisClients.Get()
        defer redisClient.Close()
        key := NEWS_CACHE_KEY + fmt.Sprintf("%d", news.NewsId)
        newsString, _ := json.Marshal(news)
        redisClient.Do("SET", key, newsString);
    }
    
    return count
}

func (dao *NewsDao) DeleteNews(newsId int64) (int64) {
    rows := dao.db.Delete(model.News{}, "news_id = ?", newsId).RowsAffected
    if rows> 0 {
        redisClient := dao.redisClients.Get()
        defer redisClient.Close()
        key := NEWS_CACHE_KEY + fmt.Sprintf("%d", newsId)
        redisClient.Do("DEL", key);
    }
    return rows
}