package ke

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	rootPath string
	appPath string
	controllerSuffix string
)

func init()  {
	rootPath = GetConfig("root_path")
	appPath = GetConfig("app_path")
	controllerSuffix = GetConfig("controller_suffix")
}

// 解析文件
func ParseFile(path string)  {
	if !strings.HasSuffix(path, ".php") {
		return
	}
	data, err := ioutil.ReadFile(rootPath + appPath + "/" + path)
	if err != nil {
		fmt.Println(err)
		return
	}

	clams, err := regexp.Compile(`class(?s:(.*?))}\s*$`)
	if err != nil {
		println("class error", err)
		return
	}
	content := clams.FindString(string(data))


	reg := regexp.MustCompile(`/\*\*(?s:(.*?))}\s`)
	if reg == nil {
		println("match error")
		return
	}
	result := reg.FindAllStringSubmatch(content, -1)

	for _, text := range result {
		var route Route
		route.FileSrc = strings.ReplaceAll(path, "\\", "/")
		route.MethodName = getMethodName(text[1])
		setRouteInfo(text[1], &route)
		if route.RouteUri != "" {
			AddRoute(route)
		}
		// fmt.Println("text[1] = ", text[1], route.MethodName)
	}

}

// 获取路由
func setRouteInfo(str string, route *Route)  {
	reg := regexp.MustCompile(`@route\((.+?)\)`)
	if reg == nil {
		return
	}
	regs := reg.FindStringSubmatch(str)
	if len(regs) == 0 {
		return
	}
	// fmt.Println(regs[0])
	str = strings.ReplaceAll(regs[1], "'", "")
	arr := strings.Split(str, ",")
	for s := range arr {
		arr[s] = strings.TrimSpace(arr[s])
	}
	// println("s", arr)
	// fmt.Printf("%+v\r\n", arr)
	route.RouteUri = arr[0]
	if len(arr) == 2 {
		route.RouteMethod = arr[1]
	}

	path := strings.ReplaceAll(route.FileSrc, rootPath + appPath, "")
	temps := strings.Split(strings.ReplaceAll(path, "\\", "/"), "/")

	// fmt.Println("t", temps[1])

	var (
		module string
		controller string
	)

	if temps[1] == "controller" {
		module = ""
		controller = getControllerName(temps[len(temps) - 1])
		temps = temps[:len(temps)-1]
		temps = temps[2:]
		if len(temps) > 0 {
			controller = strings.Join(temps, ".") + "." + controller
		}
		route.RouteHandler = fmt.Sprintf("%s/%s", controller, route.MethodName)
	} else {
		module = temps[1]
		controller = getControllerName(temps[len(temps) - 1])
		temps = temps[:len(temps)-1]
		temps = temps[3:]
		if len(temps) > 0 {
			controller = strings.Join(temps, ".") + "." + controller
		}
		route.RouteHandler = fmt.Sprintf("%s/%s/%s", module, controller, route.MethodName)
	}


}

// 获取方法名
func getMethodName(str string) string {
	reg := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	if reg == nil {
		return ""
	}
	match := reg.FindStringSubmatch(str)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}

// 解出控制器名
func getControllerName(str string) string {
	reg := regexp.MustCompile(`(Controller)*\.php`)
	if reg == nil {
		return ""
	}
	return reg.ReplaceAllString(str, "")
}
