package config

import (
    "github.com/jinzhu/configor"
)

var Config = struct {
    APPName string `default:"app name"`
    
    DB struct {
        Name     string `default:"test"`
        User     string `default:"root"`
        Password string `required:"true" default:"password"`
        Port     uint   `default:"3306"`
    }

    Redis struct {
        Host    string  `default:"127.0.0.1"`
        Port    uint    `default:"6379"`
    }
}{}

func Init() {
    configor.Load(&Config, "config/config.yml")
}