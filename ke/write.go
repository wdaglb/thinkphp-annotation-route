package ke

import (
	"fmt"
	"os"
	"strings"
)

var (
	file string
)

// 写入文件
func Write(sendfile string)  {
	file = sendfile

	var files []string
	files, err := GetAllFile(rootPath + appPath, files)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// println(ke.GetConfig("root_path") + ke.GetConfig("app_path"))
	for _, file := range files {
		file = strings.ReplaceAll(file, "\\", "/")
		ParseFile(strings.ReplaceAll(file, rootPath + appPath, ""))
		if file != "" {
			fmt.Println(file + " ok!")
		}
	}

	save()
}

// 更新更新的文件
func Update()  {
	files := GetActiveFile()

	// fmt.Printf("active: %v\r\n", files)
	// println(ke.GetConfig("root_path") + ke.GetConfig("app_path"))
	for _, file := range files {
		// println("fff", file)
		file = strings.ReplaceAll(file, "\\", "/")
		RmfileRoute(strings.ReplaceAll(file, rootPath + appPath, ""))
		// println("fff2", strings.ReplaceAll(file, rootPath + appPath, ""))
		ParseFile(strings.ReplaceAll(file, rootPath + appPath, ""))
		if file != "" {
			fmt.Println(file + " ok!")
		}
	}
	save()
}

// 保存文件
func save()  {
	var data []string

	data = append(data, "<?php\r\n/* build_route提示：本文件为自动生成，请不要编辑 */\r\n")
	for _, item := range GetAllRoute() {
		if item.RouteMethod == "" {
			item.RouteMethod = "any"
		}
		data = append(data, fmt.Sprintf("Route::%s('%s', '%s');",
			item.RouteMethod,
			item.RouteUri,
			item.RouteHandler))
	}
	fs, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fs.Close()
	n, err := fs.Write([]byte(strings.Join(data, "\r\n")))
	if err == nil && n < len(data) {
		fmt.Println(err.Error())
		return
	}
	// println("OK")
}
