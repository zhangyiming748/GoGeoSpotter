package ocr

import (
	"fmt"
	"github.com/h2non/filetype"
	"os"
	"path/filepath"
)

func Pictures(root string) (files []string) {
	// 替换为你想要遍历的根目录
	//var files []string

	// Walk 函数会遍历目录及子目录下的所有文件
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 只添加文件，不添加目录
		if !info.IsDir() {
			// Open a file descriptor
			file, _ := os.Open(path)
			// We only have to pass the file header = first 261 bytes
			head := make([]byte, 261)
			file.Read(head)
			if filetype.IsImage(head) {
				fmt.Println("File is an image")
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录时出错:", err)
		return
	}

	// 打印文件列表
	//for _, file := range files {
	//	fmt.Println(file)
	//}
	return files
}
