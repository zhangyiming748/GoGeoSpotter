package regeo

import (
	"GoGeoSpotter/model"
	"GoGeoSpotter/mysql"
	"GoGeoSpotter/util"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var HOST string = "https://restapi.amap.com/v3/geocode/regeo"

type Rep struct {
	Status    string `json:"status"`
	Regeocode struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"regeocode"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
}
type One struct {
	Latitude  string //纬度
	Longitude string //经度
	Address   string
}

func GetAddrFromGEO(key, location, extensions string) (*One, error) {
	log.Printf("请求1")
	g := new(model.Geo)
	o := new(One)
	headers := map[string]string{
		"Accept": "application/json",
	}
	n, e := conv(strings.Split(location, ",")[0], strings.Split(location, ",")[1])
	g.Longitude = e
	g.Latitude = n
	if mysql.UseMysql() {
		has, _ := g.FindByCoordinate()
		if has {
			o.Address = g.Address
			o.Latitude = g.Latitude
			o.Longitude = g.Longitude
			log.Printf("从缓存中找到:%+v\n", o)
			return o, nil
		}
	}
	location = strings.Join([]string{e, n}, ",")
	//fmt.Printf("location:%s\n", location)
	location = strings.Replace(location, ".,", "", -1)
	radius := os.Getenv("RADIUS")
	if radius == "" {
		radius = "100"
	}
	params := map[string]string{
		"key":        key,        // 用户在高德地图官网 申请 Web 服务 API 类型 Key
		"location":   location,   // 传入内容规则：经度在前，纬度在后，经纬度间以“,”分割，经纬度小数点后不要超过 6 位。
		"extensions": extensions, // extensions 参数默认取值是 base，也就是返回基本地址信息；extensions 参数取值为 all 时会返回基本地址信息、附近 POI 内容、道路信息以及道路交叉口信息。
		"output":     "JSON",     // 可选输入内容包括：JSON，XML。设置 JSON 返回结果数据将会以 JSON 结构构成；如果设置 XML 返回结果数据将以 XML 结构构成。
		"radius":     radius,     // 搜索半径 radius 取值范围：0~3000，默认值：1000。单位：米
	}
	b, err := util.HttpGet(headers, params, HOST)
	log.Printf("请求2:%v\n", string(b))
	if err != nil {
		return nil, err
	}
	f, _ := os.OpenFile("example.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	f.Write(b)
	var p Rep
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}
	log.Println(p.Regeocode.FormattedAddress)
	g.Address = p.Regeocode.FormattedAddress
	o.Latitude = n
	o.Longitude = e
	o.Address = p.Regeocode.FormattedAddress

	one, err := g.InsertOne()
	if err != nil {
		log.Printf("插入失败:%v\n", err)
		return nil, err
	} else {
		log.Printf("插入成功:%v\n", one)
	}

	return o, nil
}

/*
E 代表东经（East），是用来表示经度的；而 N 代表北纬（North），用于表示纬度 。经度用来标识地球表面东西方向的位置，纬度则标识南北方向的位置。与之对应的，西经用 W（West）表示，南纬用 S（South）表示。
*/
func conv(N, E string) (string, string) {
	//390848,1171118
	n := hms2num(N)
	e := hms2num(E)
	return fmt.Sprintf("%.4f", n), fmt.Sprintf("%.4f", e)
}
func hms2num(hms string) float64 {
	// 提取最后两位数字保存为s
	s := hms[len(hms)-2:]
	// 删除最后两位数字
	hms = hms[:len(hms)-2]
	// 再次提取最后两位数字保存为m
	m := hms[len(hms)-2:]
	// 删除这两位数
	hms = hms[:len(hms)-2]
	//fmt.Printf("h: %v, s: %v, m: %v\n", hms, s, m)
	hh, _ := strconv.ParseFloat(hms, 64)
	mm, _ := strconv.ParseFloat(m, 64)
	ss, _ := strconv.ParseFloat(s, 64)
	num := hh + mm/60 + ss/3600
	return num
}
