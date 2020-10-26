package jd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Kaytz/KayProxy/common"
	"github.com/Kaytz/KayProxy/network"
	"github.com/Kaytz/KayProxy/utils"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	path1  = "serverConfig"
	path2  = "wareBusiness"
	path3  = "basicConfig"
	floors common.SliceType
)

func AddPrice(httpInfo *common.HttpInfo) bool {
	modified := false

	jsonBody, err := simplejson.NewJson(utils.ToJsonByte(httpInfo.JsonBody))
	if err != nil {
		return false
	}
	if strings.Contains(httpInfo.EndPoint, path1) {
		jsonBody.Get("serverConfig").Del("httpdns")
		jsonBody.Get("serverConfig").Del("dnsvip")
		jsonBody.Get("serverConfig").Del("dnsvip_v6")
		modified = true
		m, err := jsonBody.Map()
		if err != nil {
			return false
		}
		httpInfo.JsonBody = m
	}
	if strings.Contains(httpInfo.EndPoint, path2) {
		floors, err = jsonBody.Get("floors").Array()
		if err != nil {
			return false
		}
		commodityInfo := floors[len(floors)-1].(map[string]interface{})
		shareUrl := utils.ToSimpleJson(commodityInfo).Get("data").Get("property").Get("shareUrl").MustString()
		//shareUrl := commodity_info["data"].(map[string]interface{})["property"].(map[string]interface{})["shareUrl"].(string)
		requestHistoryPrice(shareUrl, requestHistoryPriceCallback)
		modified = true
		jsonBody.Set("floors", floors)
		m, err := jsonBody.Map()
		if err != nil {
			return false
		}
		httpInfo.JsonBody = m
	}
	if strings.Contains(httpInfo.EndPoint, path3) {
		jsonBody.Get("data").Get("JDHttpToolKit").Del("httpdns")
		jsonBody.Get("data").Get("JDHttpToolKit").Del("dnsvipV6")
		modified = true
		m, err := jsonBody.Map()
		if err != nil {
			return false
		}
		httpInfo.JsonBody = m
	}
	return modified
}

func requestHistoryPriceCallback(data map[string]interface{}) {
	if data != nil {
		jsonData, err := simplejson.NewJson(utils.ToJsonByte(data))
		if err != nil {
			return
		}
		lowerWord := adwordObj()
		lowerWord.Get("data").Get("ad").Set("textColor", "#fe0000")
		bestIndex := 0
		for index := 0; index < len(floors); index++ {
			element := floors[index].(map[string]interface{})
			if element["mId"].(string) == lowerWord.Get("mId").MustString() {
				bestIndex = index + 1
				break
			} else {
				sortId, _ := element["sortId"].(json.Number).Float64()
				if sortId > lowerWord.Get("sortId").MustFloat64() {
					bestIndex = index
					break
				}
			}
		}
		if jsonData.Get("ok").MustInt() == 1 && jsonData.Get("single") != nil {
			lower := lowerMsgs(jsonData.Get("single"))
			detail := priceSummary(jsonData)
			tip := data["PriceRemark"].(map[string]interface{})["Tip"].(string) + "（仅供参考）"
			lowerWord.MustMap()["data"].(map[string]interface{})["ad"].(map[string]interface{})["adword"] = fmt.Sprintf(`%s %s\n%s`, lower, tip, detail)
			floors = insert(floors, bestIndex, lowerWord.MustMap())
		}
		if jsonData.Get("ok").MustInt() == 0 && len(jsonData.Get("msg").MustString()) > 0 {
			lowerWord.MustMap()["data"].(map[string]interface{})["ad"].(map[string]interface{})["adword"] = "⚠️ " + data["msg"].(string)
			floors = insert(floors, bestIndex, lowerWord.MustMap())
		}
	}
}

type Callback func(x interface{})

func requestHistoryPrice(shareUrl string, callback func(data map[string]interface{})) {
	header := make(http.Header, 4)
	header["Content-Type"] = append(header["Content-Type"], "application/x-www-form-urlencoded;charset=utf-8")
	header["User-Agent"] = append(header["User-Agent"], "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 - mmbWebBrowse - ios")
	clientRequest := network.ClientRequest{
		Method:    http.MethodPost,
		Host:      "apapia-history.manmanbuy.com",
		RemoteUrl: "https://apapia-history.manmanbuy.com/ChromeWidgetServices/WidgetServices.ashx",
		Header:    header,
		Body:      ioutil.NopCloser(bytes.NewBufferString("methodName=getHistoryTrend&p_url=" + url.QueryEscape(shareUrl))),
		Proxy:     true,
	}
	resp, err := network.Request(&clientRequest)
	if err != nil {
		callback(nil)
	}
	defer resp.Body.Close()
	body, err := network.GetResponseBody(resp)
	if err != nil {
		callback(nil)
	}
	jsonBody := utils.ParseJson(body)
	callback(jsonBody)

}

func adwordObj() *simplejson.Json {
	js, _ := simplejson.NewJson([]byte(`{
        "bId": "eCustom_flo_199",
        "cf": {
            "bgc": "#ffffff",
            "spl": "empty"
        },
        "data": {
            "ad": {
                "adword": "",
                "textColor": "#8C8C8C",
                "color": "#f23030",
                "newALContent": true,
                "hasFold": true,
                "class": "com.jd.app.server.warecoresoa.domain.AdWordInfo.AdWordInfo",
                "adLinkContent": "",
                "adLink": ""
            }
        },
        "mId": "bpAdword",
        "refId": "eAdword_0000000028",
        "sortId": 13
    }`))
	return js
}
func lowerMsgs(data *simplejson.Json) string {
	lower := data.Get("lowerPriceyh").MustString()
	lowerDate := dateFormat(data.Get("lowerDateyh").MustString())
	lowerMsg := fmt.Sprintf("〽️历史最低到手价：¥%s (%s) ", lower, lowerDate)
	return lowerMsg
}
func priceSummary(data *simplejson.Json) string {
	summary := ""
	listPriceDetail := data.Get("PriceRemark").Get("ListPriceDetail").MustArray()
	listPriceDetail = listPriceDetail[:len(listPriceDetail)-1] // 删除尾部1个元素
	hisSummary := historySummary(data.Get("single").MustMap())
	list := append(listPriceDetail, hisSummary...)

	for _, item := range list {
		jsonItem, _ := simplejson.NewJson(utils.ToJsonByte(item))
		name := jsonItem.Get("Name").MustString()
		if name == "双11价格" {
			jsonItem.Set("Name", "双十一价格")
		} else if name == "618价格" {
			jsonItem.Set("Name", "六一八价格")
		} else if name == "30天最低价" {
			jsonItem.Set("Name", "三十天最低")
		}
		summary += fmt.Sprintf(`\n%s%s%s%s%s%s%s`, name, getSpace(8), jsonItem.Get("Price").MustString(), getSpace(8), jsonItem.Get("Date").MustString(), getSpace(8), jsonItem.Get("Difference").MustString())
	}
	return summary
}
func historySummary(single common.MapType) common.SliceType {

	var currentPrice float64
	var lowest60, lowest180, lowest360 *simplejson.Json
	flowRegexp := regexp.MustCompile(`(\[.*?\])`)
	timeRegexp := regexp.MustCompile(`\[(.*)000,(.*),"(.*)"\]`)
	list := flowRegexp.FindStringSubmatch(single["jiagequshiyh"].(string))
	for index, param := range list {
		if len(param) > 0 {
			result := timeRegexp.FindStringSubmatch(param)
			i, _ := strconv.ParseInt(result[1], 10, 64)
			date := time.Unix(i, 0).Format("2006-01-02")
			price, _ := strconv.ParseFloat(result[2], 64)
			if index == 0 {
				currentPrice := price
				lowest60Str := fmt.Sprintf(`{ Name: "六十天最低", Price: "%d", Date: "%s", Difference: %s, price: %f }`, price, date, difference(currentPrice, price), date)
				lowest180Str := fmt.Sprintf(`{ Name: "一百八最低", Price: "%d", Date: "%s", Difference: %s, price: %f }`, price, date, difference(currentPrice, price), date)
				lowest360Str := fmt.Sprintf(`{ Name: "三百六最低", Price:"%d", Date: "%s", Difference: %s, price: %f }`, price, date, difference(currentPrice, price), date)
				lowest60, _ = simplejson.NewFromReader(bytes.NewBuffer([]byte(lowest60Str)))
				lowest180, _ = simplejson.NewFromReader(bytes.NewBuffer([]byte(lowest180Str)))
				lowest360, _ = simplejson.NewFromReader(bytes.NewBuffer([]byte(lowest360Str)))
			}
			if lowest60 == nil || lowest180 == nil || lowest360 == nil {
				continue
			}
			if index < 60 && price <= lowest60.Get("price").MustFloat64() {
				lowest60.Set("price", price)
				lowest60.Set("Price", fmt.Sprintf(`¥%s`, price))
				lowest60.Set("Date", date)
				lowest60.Set("Difference", difference(currentPrice, price))
			}
			if index < 180 && price <= lowest180.Get("price").MustFloat64() {
				lowest180.Set("price", price)
				lowest180.Set("Price", fmt.Sprintf(`¥%s`, price))
				lowest180.Set("Date", date)
				lowest180.Set("Difference", difference(currentPrice, price))
			}
			if index < 360 && price <= lowest360.Get("price").MustFloat64() {
				lowest360.Set("price", price)
				lowest360.Set("Price", fmt.Sprintf(`¥%s`, price))
				lowest360.Set("Date", date)
				lowest360.Set("Difference", difference(currentPrice, price))
			}
		}
	}
	return common.SliceType{lowest60.MustMap(), lowest180.MustMap(), lowest360.MustMap()}
}
func difference(currentPrice float64, price float64) string {
	difference := sub(currentPrice, price)
	if difference == 0 {
		return "-"
	} else {
		return fmt.Sprintf("%s%.2f", If(difference > 0, "↑", "↓").(string), difference)
	}
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func sub(arg1 float64, arg2 float64) float64 {
	return add(arg1, -(arg2))
}

func add(arg1 float64, arg2 float64) float64 {
	arg1Str := strconv.FormatFloat(arg1, 'E', -1, 64)
	arg2Str := strconv.FormatFloat(arg2, 'E', -1, 64)
	arg1Arr := strings.Split(arg1Str, ".")
	arg2Arr := strings.Split(arg2Str, ".")
	d1 := ""
	if len(arg1Arr) == 2 {
		d1 = arg1Arr[1]
	}
	d2 := ""
	if len(arg2Arr) == 2 {
		d2 = arg2Arr[1]
	}
	maxLen := Max(len(d1), len(d2))
	m := math.Pow10(maxLen)
	//result := strconv.FormatFloat((arg1*m+arg2*m)/m, 'f', 2, 32)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", (arg1*m+arg2*m)/m), 64)
	return value
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func dateFormat(cellval string) string {
	cellIntVal := strings.Replace(strings.Replace(cellval, "/Date(", "", -1), "000-0000)/", "", -1)
	i, _ := strconv.ParseInt(cellIntVal, 10, 64)
	datetime := time.Unix(i, 0).Format("2006-01-02")
	return datetime
}
func getSpace(length int) string {
	blank := ""
	for index := 0; index < length; index++ {
		blank += " "
	}
	return blank
}
func insert(slice []interface{}, index int, item interface{}) []interface{} {

	tmp := append([]interface{}{}, slice[index:]...)

	//拼接插入的元素
	slice = append(slice[0:index], item)

	//与临时切片再组合得到最终的需要的切片
	slice = append(slice, tmp...)
	return slice
}
