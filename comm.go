package main

import (
    "math"
    "strings"
    "log"
    "strconv"
)

/**
 * 设置页码
 * @page int
 * @page_size int
 * @count int
 * @return int
 */
func setPage(page int, pageSize int, count int) int {
    
    if page < 1 {
        page = 1
    }
    
    lastPage := int(math.Ceil(float64(count)/float64(pageSize)))
    
    if page > lastPage {
        page = lastPage
    }
    
    return page    
}

/**
 * 版本号转换为数字
 * @version
 * @return int
 */
func versionToNumber(version string) int {
    versions := strings.Split(version, ".")
    if len(versions) != 3 {
        return 0
    }
    version_number := 0
    for key, value := range versions {
        i, _ := strconv.Atoi(value)
        version_number += (int(math.Pow10((2-key) * 3)) * i) 
    }

    log.Println(version_number)
    return version_number
}

/**
 * 返回统一数据格式
 * @err 错误码
 * @message 信息
 * @data 数据
 * @return interface{}
 */
func showJson(err int, message string, data interface{}) interface{}{
    rs := make(map[string]interface{})
    rs["err"] = err
    rs["data"] = data
    rs["message"] = message
    return rs
}