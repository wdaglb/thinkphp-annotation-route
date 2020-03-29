package ke


type Route struct {
	MethodName string
	RouteUri string
	RouteMethod string
	RouteHandler string
	FileSrc string
}

var routes []Route

// 添加路由
func AddRoute(route Route)  {
	routes = append(routes, route)
}

// 删除文件的路由
func RmfileRoute(file string)  {
	if file == "" {
		return
	}
	var arr []Route
	// println("删除文件路由：", file)
	for _, f := range routes {
		// println("文件路由：", f.FileSrc, file)
		if f.FileSrc != file {
			arr = append(arr, f)
		}
	}
	routes = arr
}

func GetAllRoute() []Route {
	return routes
}