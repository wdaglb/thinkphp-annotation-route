package ke

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	data = map[string]string{}
)

func SetConfig(key string, value string)  {
	data[key] = value
}

func GetConfig(key string) string {
	return data[key]
}

func GetAllConfig() map[string]string {
	return data
}

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err.Error())
	}

	rootPath := flag.String("root", dir + "/", "root_path")
	appPath := flag.String("app", "application", "app_path")
	routeFile := flag.String("route", "./route/build_route.php", "route")
	isWatch := flag.Bool("watch", false, "watch")

	flag.Parse()
	// println(*rootPath, *appPath, *isWatch)
	// 设置根目录
	SetConfig("root_path", strings.ReplaceAll(*rootPath, "\\", "/"))
	// 设置应用目录名
	SetConfig("app_path", *appPath)
	// 设置控制器后缀
	SetConfig("controller_suffix", "Controller")
	// 保存路由文件
	SetConfig("route_file", *routeFile)
	// 是否监听文件变动
	if *isWatch {
		SetConfig("is_watch", "1")
	} else {
		SetConfig("is_watch", "0")
	}
}