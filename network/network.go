package network

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/dsnet/compress/brotli"

	"github.com/Kaytz/KayProxy/common"
	"github.com/Kaytz/KayProxy/config"
	"github.com/Kaytz/KayProxy/utils"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

func init() {

	insecure := flag.Bool("insecure-ssl", false, "Accept/Ignore all server SSL certificates")
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := ioutil.ReadFile(*config.CaCrtFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", *config.CaCrtFile, err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}
	// Trust the augmented cert pool in our client
	caConfig := &tls.Config{
		InsecureSkipVerify: *insecure,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          300,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		MaxConnsPerHost:       100,
		TLSClientConfig:       caConfig,
	}
	httpClient = &http.Client{
		Transport: tr,
	}
}

type ClientRequest struct {
	Method               string
	RemoteUrl            string
	Host                 string
	ForbiddenEncodeQuery bool
	Header               http.Header
	Body                 io.Reader
	Cookies              []*http.Cookie
	Proxy                bool
	ConnectTimeout       time.Duration
}

func Request(clientRequest *ClientRequest) (*http.Response, error) {
	//log.Println("remoteUrl:" + clientRequest.RemoteUrl)
	method := clientRequest.Method
	remoteUrl := clientRequest.RemoteUrl
	host := clientRequest.Host
	header := clientRequest.Header
	body := clientRequest.Body
	proxy := clientRequest.Proxy
	cookies := clientRequest.Cookies
	connectTimeout := clientRequest.ConnectTimeout
	if connectTimeout == 0 {
		connectTimeout = 10 * time.Second
	}
	var resp *http.Response
	request, err := http.NewRequest(method, remoteUrl, body)
	if err != nil {
		log.Printf("NewRequest fail:%v\n", err)
		return resp, nil
	}
	if !clientRequest.ForbiddenEncodeQuery {
		request.URL.RawQuery = request.URL.Query().Encode()
	}
	if len(host) > 0 {
		request.Host = host
		request.Header.Set("host", host)
	}
	if len(request.URL.Scheme) == 0 {
		if request.TLS != nil {
			request.URL.Scheme = "https"
		} else {
			request.URL.Scheme = "http"
		}
	}
	tr := httpClient.Transport.(*http.Transport)
	tr.TLSClientConfig = nil
	// if request.URL.Scheme == "https" || request.TLS != nil {
	// fix redirect to https://music.163.com
	if _, ok := common.HostDomain[request.Host]; ok {
		tr.TLSClientConfig = &tls.Config{}
		// verify music.163.com certificate
		tr.TLSClientConfig.ServerName = request.Host //it doesn't contain any IP SANs
		// redirect to music.163.com will need verify self
		tr.TLSClientConfig.InsecureSkipVerify = true
	}
	// }

	if proxy { //keep headers&cookies for Direct
		if header != nil {
			request.Header = header
		}
		for _, value := range cookies {
			request.AddCookie(value)
		}
	}
	accept := "application/json, text/plain, */*"
	acceptEncoding := "gzip, deflate"
	acceptLanguage := "zh-CN,zh;q=0.9"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"

	if header != nil {
		accept = header.Get("accept")
		if len(accept) == 0 {
			accept = "application/json, text/plain, */*"
		}
		acceptEncoding = header.Get("accept-encoding")
		if len(acceptEncoding) == 0 {
			acceptEncoding = "gzip, deflate"
		}
		acceptLanguage = header.Get("accept-language")
		if len(acceptLanguage) == 0 {
			acceptLanguage = "zh-CN,zh;q=0.9"
		}
		userAgent = header.Get("user-agent")
		if len(userAgent) == 0 {
			userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
		}
		Range := header.Get("range")
		if len(Range) > 0 {
			request.Header.Set("range", Range)
		}
	}

	request.Header.Set("accept", accept)
	request.Header.Set("accept-encoding", acceptEncoding)
	request.Header.Set("accept-language", acceptLanguage)
	request.Header.Set("user-agent", userAgent)
	resp, err = httpClient.Do(request)
	if err != nil {
		//log.Println(request.Method, request.URL.String(), host)
		log.Printf("http.Client.Do fail:%v\n", err)
		return resp, err
	}

	return resp, err

}
func GetResponseBody(response *http.Response) ([]byte, error) {
	encode := response.Header.Get("Content-Encoding")
	decryptBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read  body fail")
		return decryptBody, err
	}
	response.Body.Close()
	if len(encode) > 0 {
		if strings.Contains(encode, "gzip") || strings.Contains(encode, "deflate") {
			decryptBody, err = utils.UnGzip(decryptBody)
			if err != nil {
				log.Println("read  body fail")
				return decryptBody, err
			}
		}
		if strings.Contains(encode, "br") {
			gr, _ := brotli.NewReader(bytes.NewReader(decryptBody), nil)
			gb, err := ioutil.ReadAll(gr)
			if err != nil {
				log.Println("read  body fail")
				return gb, err
			}
			decryptBody = gb
		}
	}
	return decryptBody, nil
}
