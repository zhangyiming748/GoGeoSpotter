package regeo

import (
	"GoGeoSpotter/model"
	"GoGeoSpotter/mysql"
	"os"
	"testing"
)

func TestConv(t *testing.T) {
	n, e := conv("390848", "1171118")
	t.Logf("n = %v\te = %v\n", n, e)
}
func TestGetAddrFromGEO(t *testing.T) {
	mysql.SetMysql()
	mysql.GetMysql().Sync(model.Geo{})
	key := os.Getenv("KEY")
	location := "395712,1162438"
	extensions := "base"
	geo, err := GetAddrFromGEO(key, location, extensions)
	if err != nil {
		return
	}
	t.Log(geo)
}
