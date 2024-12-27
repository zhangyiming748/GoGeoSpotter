package main

import (
	"GoGeoSpotter/model"
	"GoGeoSpotter/mysql"
	"GoGeoSpotter/ocr"
	"GoGeoSpotter/regeo"
	"GoGeoSpotter/util"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	mysql.SetMysql()
}

func main() {
	mysql.GetMysql().Sync(model.Geo{})
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatalln("需要${KEY}")
	}
	extensions := os.Getenv("EXTENSIONS")
	if extensions == "" {
		extensions = "base"
	}
	txt := "/data/coordinates.txt"
	if runtime.GOOS == "windows" {
		txt = "coordinates.txt"
	}
	md := "/data/coordinates.md"
	if runtime.GOOS == "windows" {
		md = "coordinates.md"
	}
	results := []string{
		"|纬度|纬度经度|地址|",
		"|:---:|:---:|:---:|",
	}
	dcim := "/data/dcim"
	if runtime.GOOS == "windows" {
		dcim = "dcim"
	}
	var locations []string
	if isExistDir(dcim) {
		// todo 图片流程
		pics := ocr.Pictures(dcim)
		log.Printf("符合条件的图片文件:%v\n", pics)
		host := "http://127.0.0.1:8080/file"
		for _, pic := range pics {
			form, err := ocr.PostForm(host, pic)
			if err != nil {
				log.Printf("form err:%v\n", err)
			}
			nums := ocr.GetNums(form)
			locations = append(locations, strings.Join([]string{nums[0], nums[1]}, ","))
		}
	} else {
		locations = util.ReadByLine(txt)
	}

	log.Printf("locations:%v\n", locations)

	for _, location := range locations {
		geo, err := regeo.GetAddrFromGEO(key, location, extensions)
		if err != nil {
			log.Printf(err.Error())
		}
		result := strings.Join([]string{"|", geo.Latitude, "|", geo.Longitude, "|", geo.Address, "|"}, "")
		results = append(results, result)
		log.Println(geo)
	}
	util.WriteByLine(md, results)
}

func isExistDir(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err == nil {
		fmt.Printf("%v文件夹存在\n", dirPath)
		return true
	} else if os.IsNotExist(err) {
		fmt.Printf("%v文件夹不存在\n", dirPath)
		return false
	} else {
		fmt.Printf("%v发生错误:%v", dirPath, err)
		return false
	}
}
