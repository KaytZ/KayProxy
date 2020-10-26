package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"golang.org/x/text/width"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func UnGzipV2(gzipData io.Reader) (io.Reader, error) {
	r, err := gzip.NewReader(gzipData)
	if err != nil {
		log.Println("UnGzipV2 error:", err)
		return gzipData, err
	}
	//defer r.Close()
	return r, nil
}
func UnGzip(gzipData []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(gzipData))
	if err != nil {
		log.Println("UnGzip error:", err, ", will return original data")
		return gzipData, err
	}
	defer r.Close()
	var decryptECBBytes = gzipData
	decryptECBBytes, err = ioutil.ReadAll(r)
	if err != nil {
		log.Println("UnGzip")
		return gzipData, err
	}
	return decryptECBBytes, nil
}
func LogInterface(i interface{}) string {
	return fmt.Sprintf("%+v", i)
}
func ReplaceAll(str string, expr string, replaceStr string) string {
	reg := regexp.MustCompile(expr)
	str = reg.ReplaceAllString(str, replaceStr)
	return str
}
func ParseJson(data []byte) map[string]interface{} {
	var result map[string]interface{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	d.Decode(&result)
	return result
}
func ParseJsonV2(reader io.Reader) map[string]interface{} {
	var result map[string]interface{}
	d := json.NewDecoder(reader)
	d.UseNumber()
	d.Decode(&result)
	return result
}
func ParseJsonV3(data []byte, dest interface{}) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	return d.Decode(dest)
}
func ParseJsonV4(reader io.Reader, dest interface{}) error {
	d := json.NewDecoder(reader)
	d.UseNumber()
	return d.Decode(dest)
}
func PanicWrapper(f func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover panic : ", r)
		}
	}()
	f()
}

func ToJson(object interface{}) string {
	jsonObj, err := json.Marshal(object)
	if err != nil {
		log.Println("ToJson Error：", err)
		return "{}"
	}
	return string(jsonObj)
}
func ToJsonByte(object interface{}) []byte {
	jsonObj, err := json.Marshal(object)
	if err != nil {
		log.Println("ToJson Error：", err)
	}
	return jsonObj
}
func ToSimpleJson(object interface{}) *simplejson.Json {
	jsonObj, err := json.Marshal(object)
	if err != nil {
		log.Println("ToJson Error：", err)
	}
	sJson, err := simplejson.NewJson(jsonObj)
	if err != nil {
		log.Println("ToJson Error：", err)
	}
	return sJson
}
func Exists(keys []string, h map[string]interface{}) bool {
	for _, key := range keys {
		if !Exist(key, h) {
			return false
		}
	}
	return true
}
func Exist(key string, h map[string]interface{}) bool {
	_, ok := h[key]
	return ok
}
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".\n`)
	}
	return path[0 : i+1], nil
}
func MD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
func GenRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

var leftPairedSymbols = width.Narrow.String("(（《<[{「【『\"'")
var rightPairedSymbols = width.Narrow.String(")）》>]}」】』\"'")

func parsePairedSymbols(data string, sub string, substr []string, keyword map[string]int) string {
	data = strings.TrimSpace(data)
	leftIndex := strings.Index(leftPairedSymbols, sub)
	rightIndex := strings.Index(rightPairedSymbols, sub)
	subIndex := 0
	for index, key := range substr {
		if key == sub {
			subIndex = index
			break
		}
	}
	index := -1
	if leftIndex != -1 {
		index = leftIndex
	} else if rightIndex != -1 {
		index = rightIndex
	}
	if index != -1 {
		leftSymbol := leftPairedSymbols[index : index+len(sub)]
		rightSymbol := rightPairedSymbols[index : index+len(sub)]
		leftCount := strings.Count(data, leftSymbol)
		rightCount := strings.Count(data, rightSymbol)
		if leftCount == rightCount && leftCount > 0 {
			for i := 0; i < leftCount; i++ {
				lastLeftIndex := strings.LastIndex(data, leftSymbol)
				matchedRightIndex := strings.Index(data[lastLeftIndex:], rightSymbol)
				if matchedRightIndex == -1 {
					continue
				}
				key := strings.TrimSpace(data[lastLeftIndex+len(leftSymbol) : lastLeftIndex+matchedRightIndex])
				data = data[:lastLeftIndex] + " " + data[lastLeftIndex+matchedRightIndex+len(rightSymbol):]
				substr2 := substr[subIndex+1:]
				parseKeyWord(key, substr2, keyword)

			}
		}
	}
	return data
}
func parseKeyWord(data string, substr []string, keyword map[string]int) {
	if len(data) == 0 {
		return
	}
	data = strings.TrimSpace(data)
	for _, sub := range substr {
		if strings.Contains(data, sub) {
			if strings.Contains(leftPairedSymbols, sub) || strings.Contains(rightPairedSymbols, sub) {
				data = parsePairedSymbols(data, sub, substr, keyword)
			} else {
				splitData := strings.Split(data, sub)
				for _, key := range splitData {
					newKey := strings.TrimSpace(key)
					parseKeyWord(newKey, substr, keyword)
					data = strings.Replace(data, key, "", 1)

				}
				data = strings.ReplaceAll(data, sub, "")
			}

		}
	}
	data = strings.TrimSpace(data)
	if len(data) > 0 {
		if strings.EqualFold(data, "LIVE版") {
			data = "LIVE"
		}
		keyword[data] = 1
	}
}

type ByLenSort []string

func (a ByLenSort) Len() int {
	return len(a)
}

func (a ByLenSort) Less(i, j int) bool {
	return len(a[i]) > len(a[j])
}

func (a ByLenSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
