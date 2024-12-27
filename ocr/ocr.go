package ocr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	FORM = "http://127.0.0.1:8080/file"
)

type Rep struct {
	Result  string `json:"result"`
	Version string `json:"version"`
}

func PostForm(url, fp string) (string, error) {
	log.Println("发请求")
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(fp)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(fp))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return "", errFile1
	}
	err := writer.Close()
	if err != nil {
		return "", err

	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var p Rep
	if err = json.Unmarshal(body, &p); err != nil {
		return "", err
	}
	fmt.Println(p)
	return p.Result, nil
}

func GetNums(str string) (locations []string) {
	// 定义正则表达式
	re := regexp.MustCompile(`\d{1,3}°\d{1,2}'\d{1,2}"`)

	// 查找所有匹配的子字符串
	matches := re.FindAllString(str, -1)

	// 打印匹配结果
	for _, match := range matches {
		//fmt.Printf("%d:%s\n", i, match)
		//395446,1161142
		location, _ := PurgeNum(match)
		locations = append(locations, location)
	}
	return locations
}

func PurgeNum(str string) (string, error) {

	// 定义正则表达式，匹配度分秒中的数字部分
	re := regexp.MustCompile(`(\d+)[°'](\d+)['](\d+)["]`)
	matches := re.FindStringSubmatch(str)

	if len(matches) == 4 {
		degree := matches[1]
		minute := matches[2]
		second := matches[3]

		// 拼接数字部分
		result := degree + minute + second
		_, err := strconv.Atoi(result)
		if err != nil {
			return "", err
		}
		//fmt.Println(number)
		return result, nil
	}
	return "", nil
}
