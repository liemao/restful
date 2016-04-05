package model

import (
    "time"
)

type News struct {
    NewsId      int64       `gorm:"primary_key"`
    Title       string      `sql:"size:60"`
    Content     string      `sql:"size: 255"`
    AddTime     time.Time
    UpdateTime  time.Time
}