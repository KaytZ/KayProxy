package provider

import (
	"bytes"
	"encoding/json"
	"github.com/Kaytz/KayProxy/common"
	"github.com/Kaytz/KayProxy/provider/jd"
	"github.com/Kaytz/KayProxy/provider/yyw"
	"github.com/Kaytz/KayProxy/utils"
	"github.com/dsnet/compress/brotli"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

func JdPrice(httpInfo *common.HttpInfo) bool {
	return jd.AddPrice(httpInfo)
}
func YywOffline(httpInfo *common.HttpInfo) bool {
	return yyw.YywOffline(httpInfo)
}

func RequiresBodyAfter(response *http.Response, httpInfo *common.HttpInfo) {
	var appleMethod interface{}
	if regexp.MustCompile(`^https?://api\.m\.jd\.com/client\.action\?functionId=(wareBusiness|serverConfig)`).MatchString(httpInfo.EndPoint) {
		appleMethod = JdPrice
	}
	if regexp.MustCompile(`^http:\/\/115\.com\/lx.*$`).MatchString(httpInfo.EndPoint) || regexp.MustCompile(`^https?:\/\/webapi\.115\.com\/user\/check_sign.*$`).MatchString(httpInfo.EndPoint) {
		appleMethod = YywOffline
	}
	if appleMethod != nil {
		log.Println("Exp Match: ==> True ==> " + httpInfo.EndPoint)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		tmpBody := make([]byte, len(body))
		copy(tmpBody, body)

		if len(body) > 0 {
			encode := response.Header.Get("Content-Encoding")
			decryptBody := body
			if len(encode) > 0 {
				if strings.Contains(encode, "gzip") || strings.Contains(encode, "deflate") {
					decryptBody, _ = utils.UnGzip(decryptBody)
				}
				if strings.Contains(encode, "br") {
					gr, err := brotli.NewReader(bytes.NewReader(decryptBody), nil)
					if err != nil {
						panic(err)
					}
					gb, gerr := ioutil.ReadAll(gr)
					if err := gr.Close(); gerr == nil {
						gerr = err
					} else if gerr != nil && err == nil {
						panic("nil on Close after non-nil error")
					}
					decryptBody = gb
				}
			}
			result := utils.ParseJson(decryptBody)
			httpInfo.JsonBody = result
			httpInfo.Body = decryptBody
			modified := false

			ret1 := Apply(appleMethod, []interface{}{httpInfo})
			for _, v := range ret1 {
				modified = v.Bool()
			}

			if modified {
				response.Header.Del("transfer-encoding")
				response.Header.Del("content-encoding")
				response.Header.Del("content-length")
				if httpInfo.ChangeOriginal == true {
					response.Body = ioutil.NopCloser(bytes.NewBuffer(httpInfo.Body))
				} else {
					//httpInfo.JsonBody = httpInfo.JsonBody
					//log.Println("NeedRepackage")
					modifiedJson, _ := json.Marshal(httpInfo.JsonBody)
					response.Body = ioutil.NopCloser(bytes.NewBuffer(modifiedJson))
				}
			} else {
				//log.Println("NotNeedRepackage")
				responseHold := ioutil.NopCloser(bytes.NewBuffer(tmpBody))
				response.Body = responseHold
			}
			//log.Println(utils.ToJson(httpInfo.JsonBody))
		} else {
			responseHold := ioutil.NopCloser(bytes.NewBuffer(tmpBody))
			response.Body = responseHold
		}
	} else {
		log.Println("Exp Match: ==> False ==> " + httpInfo.EndPoint)
	}
}

func Apply(f interface{}, args []interface{}) []reflect.Value {
	fun := reflect.ValueOf(f)
	in := make([]reflect.Value, len(args))
	for k, param := range args {
		in[k] = reflect.ValueOf(param)
	}
	r := fun.Call(in)
	return r
}
