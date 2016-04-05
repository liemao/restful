package controller

import (
	"restful/service"
	"net/http"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/coreos/go-semver/semver"
    "log"
    "math"
    "strings"
    "strconv"
)

var (
	newsService *service.NewsService
)

type SemVerMiddleware struct {
	MinVersion string
	MaxVersion string
}

func (mw *SemVerMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

    minVersion, err := semver.NewVersion(mw.MinVersion)
    if err != nil {
        panic(err)
    }

    maxVersion, err := semver.NewVersion(mw.MaxVersion)
    if err != nil {
        panic(err)
    }

    return func(w rest.ResponseWriter, r *rest.Request) {

        version, err := semver.NewVersion(r.PathParam("version"))
        if err != nil {
            rs := ShowJson(http.StatusBadRequest, "Invalid version: "+err.Error(), "")
            w.WriteJson(&rs)
            return
        }

        if version.LessThan(*minVersion) {
            rs := ShowJson(http.StatusBadRequest, "Min supported version is "+minVersion.String(), "")
            w.WriteJson(&rs)
            return
        }

        if maxVersion.LessThan(*version) {
            rs := ShowJson(http.StatusBadRequest, "Max supported version is "+maxVersion.String(), "")
            w.WriteJson(&rs)
            return
        }

        r.Env["VERSION"] = version
        handler(w, r)
    }
}

func Init() {
	service.Init()

	newsService = service.NewNewsService()
	//这里可以继续初始化各个业务的service
    
    
    
    svmw := SemVerMiddleware{
        MinVersion: "1.0.0",
        MaxVersion: "3.0.0",
    }
    
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Post("/#version/news/add", svmw.MiddlewareFunc(AddNews)),
        rest.Put("/#version/news/delete/:id", svmw.MiddlewareFunc(DeleteNews)),
        rest.Get("/#version/news/list/:page/:pageSize", svmw.MiddlewareFunc(newsList)),
        rest.Get("/#version/news/:id", svmw.MiddlewareFunc(NewsDetail)),
        rest.Put("/#version/news/update", svmw.MiddlewareFunc(UpdateNews)),
        
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}



/**
 * 版本号转换为数字
 * @version
 * @return int
 */
func VersionToNumber(version string) int {
    versions := strings.Split(version, ".")
    if len(versions) != 3 {
        return 0
    }
    versionNumber := 0
    for key, value := range versions {
        i, _ := strconv.Atoi(value)
        versionNumber += (int(math.Pow10((2-key) * 3)) * i) 
    }

    return versionNumber
}

/**
 * 返回统一数据格式
 * @err 错误码
 * @message 信息
 * @data 数据
 * @return interface{}
 */
func ShowJson(err int, message string, data interface{}) interface{}{
    rs := make(map[string]interface{})
    rs["err"] = err
    rs["data"] = data
    rs["message"] = message
    return rs
}