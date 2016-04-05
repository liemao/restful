package service

import (
	"restful/dao"
    "restful/dao/newsdao"
    "math"
)

var (
    newsDao *newsdao.NewsDao
)

func Init() {
    InitKafka()
	dao.Init()
	newsDao = newsdao.NewNewsDao()
	//这里可以初始化各个业务的dao
}

//设置页码
func SetPage(page int, pageSize int, count int) int {
    
    if page < 1 {
        page = 1
    }
    
    lastPage := int(math.Ceil(float64(count)/float64(pageSize)))
    
    if page > lastPage {
        page = lastPage
    }
    
    return page
}

//设置分页offset
func SetOffset(page int, pageSize int) (int) {
    offset := ((page - 1) * pageSize)
    return offset
}