package model

import (
	"GoGeoSpotter/mysql"
	"fmt"
	"log"
	"strings"
	"time"
)

type Geo struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Latitude  string    `xorm:"comment('(北N)纬度') VARCHAR(255)"`
	Longitude string    `xorm:"comment('(东E)经度') VARCHAR(255)"`
	Address   string    `xorm:"comment('格式化位置') VARCHAR(255)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (g *Geo) InsertOne() (int64, error) {
	return mysql.GetMysql().InsertOne(g)
}
func (g *Geo) FindByLatitude() (bool, error) {
	return mysql.GetMysql().Where("latitude = ?", g.Latitude).Get(g)
}
func (g *Geo) FindByLongitude() (bool, error) {
	return mysql.GetMysql().Where("longitude = ?", g.Longitude).Get(g)
}
func (g *Geo) FindByCoordinate() (bool, error) {
	return mysql.GetMysql().Where("longitude = ?", g.Longitude).And("latitude = ?", g.Latitude).Get(g)
}
func (g *Geo) FindAddressBySql() (address string, notfound error) {
	sql := strings.Join([]string{"SELECT * FROM geo WHERE latitude = '", g.Latitude, "' AND longitude = '", g.Longitude, "' ORDER BY created_at DESC LIMIT 1"}, "")
	log.Printf("查询语句%s\n", sql)
	results, err := mysql.GetMysql().QueryString(sql)
	if err != nil {
		panic(err)
	}
	log.Println(results)
	if len(results) == 0 {
		return "", fmt.Errorf("未找到记录")
	}
	for _, row := range results {
		log.Printf("row = %+v\n", row)
		return row["address"], nil

	}
	return "", err
}
