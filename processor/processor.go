package processor

import (
	"github.com/Kaytz/KayProxy/common"
	"github.com/Kaytz/KayProxy/network"
	"github.com/Kaytz/KayProxy/provider"
	"github.com/Kaytz/KayProxy/utils"
	"net/http"
	"strings"
)

func RequestBefore(request *http.Request) *common.HttpInfo {
	httpInfo := &common.HttpInfo{Path: request.URL.Path}
	httpInfo.EndPoint = GetURL(request)
	if request.Method == http.MethodPost {
		if strings.Contains(httpInfo.Path, "/........") {

		} else if strings.Index(httpInfo.Path, "/weapi/") == 0 || strings.Index(httpInfo.Path, "/api/") == 0 {
			request.Header.Set("X-Real-IP", "118.66.66.66")
			httpInfo.Web = true
			httpInfo.Path = utils.ReplaceAll(httpInfo.Path, `^\/weapi\/`, "/api/")
			httpInfo.Path = utils.ReplaceAll(httpInfo.Path, `\?.+$`, "")
			httpInfo.Path = utils.ReplaceAll(httpInfo.Path, `\/\d*$`, "")
		}
	}
	return httpInfo
}
func Request(request *http.Request, remoteUrl string) (*http.Response, error) {
	clientRequest := network.ClientRequest{
		Method:    request.Method,
		RemoteUrl: remoteUrl,
		Host:      request.Host,
		Header:    request.Header,
		Body:      request.Body,
		Proxy:     true,
	}
	return network.Request(&clientRequest)
}
func RequestAfter(request *http.Request, response *http.Response, httpInfo *common.HttpInfo) {
	if response.StatusCode == 200 {

		provider.RequiresBodyAfter(response, httpInfo)
		//
		//defer response.Body.Close()
		//body, _ := ioutil.ReadAll(response.Body)
		//tmpBody := make([]byte, len(body))
		//copy(tmpBody, body)
		//
		//if len(body) > 0 {
		//	encode := response.Header.Get("Content-Encoding")
		//	decryptBody := body
		//	if len(encode) > 0 {
		//		if strings.Contains(encode, "gzip") || strings.Contains(encode, "deflate") {
		//			decryptBody, _ = utils.UnGzip(decryptBody)
		//		}
		//		if strings.Contains(encode, "br") {
		//			decryptBody, _ = dec.DecompressBuffer(body, nil)
		//			//decryptBody, _ = ioutil.ReadAll(reader)
		//		}
		//	}
		//	//var decryptBody, err = network.GetResponseBody(response)
		//	//if err != nil {
		//	//	log.Println("GetResponseBody fail")
		//	//	return
		//	//}
		//	result := utils.ParseJson(decryptBody)
		//	httpInfo.JsonBody = result
		//	httpInfo.Body = decryptBody
		//	modified := false
		//
		//	if regexp.MustCompile(`^https?://api\.m\.jd\.com/client\.action\?functionId=(wareBusiness|serverConfig)`).MatchString(httpInfo.EndPoint) {
		//		modified = provider.JdPrice(httpInfo)
		//	}
		//	if regexp.MustCompile(`^http:\/\/115\.com\/lx.*$`).MatchString(httpInfo.EndPoint) || regexp.MustCompile(`^https?:\/\/webapi\.115\.com\/user\/check_sign.*$`).MatchString(httpInfo.EndPoint) {
		//		modified = provider.YywOffline(httpInfo)
		//	}
		//	if modified {
		//		response.Header.Del("transfer-encoding")
		//		response.Header.Del("content-encoding")
		//		response.Header.Del("content-length")
		//		if httpInfo.ChangeOriginal == true {
		//			response.Body = ioutil.NopCloser(bytes.NewBuffer(httpInfo.Body))
		//		} else {
		//			//httpInfo.JsonBody = httpInfo.JsonBody
		//			//log.Println("NeedRepackage")
		//			modifiedJson, _ := json.Marshal(httpInfo.JsonBody)
		//			response.Body = ioutil.NopCloser(bytes.NewBuffer(modifiedJson))
		//		}
		//	} else {
		//		//log.Println("NotNeedRepackage")
		//		responseHold := ioutil.NopCloser(bytes.NewBuffer(tmpBody))
		//		response.Body = responseHold
		//	}
		//	//log.Println(utils.ToJson(httpInfo.JsonBody))
		//} else {
		//	responseHold := ioutil.NopCloser(bytes.NewBuffer(tmpBody))
		//	response.Body = responseHold
		//}
	} else {
		//log.Println("Not Process")
	}
}
func GetURL(r *http.Request) (Url string) {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host, r.RequestURI}, "")
}
