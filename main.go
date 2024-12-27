package main

import (
	"GoGeoSpotter/model"
	"GoGeoSpotter/mysql"
	"GoGeoSpotter/regeo"
	"GoGeoSpotter/util"
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
	mysql.SetMysql()
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
		"|经度|纬度|地址|",
		"|:---:|:---:|:---:|",
	}

	locations := util.ReadByLine(txt)
	for _, location := range locations {
		geo, err := regeo.GetAddrFromGEO(key, location, extensions)
		if err != nil {
			return
		}
		result := strings.Join([]string{"|", geo.Latitude, "|", geo.Longitude, "|", geo.Address, "|"}, "")
		results = append(results, result)
		log.Println(geo)
	}
	util.WriteByLine(md, results)
}
