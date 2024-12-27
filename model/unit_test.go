package model

import (
	"GoGeoSpotter/mysql"
	"testing"
)

func TestSql(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(Geo{})
	one := new(Geo)
	one.Latitude = "39.9533"
	one.Longitude = "116.4106"
	sql, err := one.FindAddressBySql()
	if err != nil {
		return
	}
	t.Logf("%+v\n", sql)
}
func TestSame(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(Geo{})
	one := new(Geo)
	one.Latitude = "39.9533"
	one.Longitude = "116.4106"
	sql, err := one.FindByCoordinate()
	if err != nil {
		return
	}
	t.Logf("%+v\n", sql)
	t.Logf("%+v\n", one)
}
