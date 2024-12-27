package ocr

import (
	"strings"
	"testing"
)

func TestOCR(t *testing.T) {
	fname := "微信图片_20241227161411.jpg"
	form, err := PostForm("http://127.0.0.1:8080/file", fname)
	if err != nil {
		return
	}
	result := GetNums(form)
	t.Log(strings.Join([]string{result[0], result[1]}, ","))
}
