package ke

import (
	"fmt"
	"io/ioutil"
	"strings"
)


var (
	files []string // 已操作文件
	acts []string // 待操作文件
)

// 添加文件
func AddFile(src string)  {
	if strings.HasSuffix(src, ".php") {
		src = strings.ReplaceAll(src, "\\", "/")
		// println("ssss", src)
		for _, f := range acts {
			if f == src {
				return
			}
		}
		acts = append(acts, src)
	}
}

// 删除文件
func DelFile(src string)  {
	if strings.HasSuffix(src, ".php") {
		src = strings.ReplaceAll(src, "\\", "/")
		for i, v := range acts {
			if v == src {
				acts = append(acts[:i], acts[i+1:]...)
				break
			}
		}
	}
}

// 获取所有文件
func GetAllFile(pathname string, s []string) ([]string, error) {
	//if len(files) > 0 {
	//	arr := make([]string, len(files) + len(acts))
	//	for _, v := range files {
	//		arr = append(arr, v)
	//	}
	//	for _, v := range acts {
	//		arr = append(arr, v)
	//	}
	//	acts = acts[0:0]
	//	return arr, nil
	//}
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	files = s
	return s, nil
}

// 获取活动文件并在内存中删除
func GetActiveFile() []string {
	arr := make([]string, len(acts))
	for _, f := range acts {
		arr = append(arr, f)
	}
	acts = acts[0:0]
	return arr
}
