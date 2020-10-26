package common

type MapType = map[string]interface{}
type SliceType = []interface{}
type HttpInfo struct {
	Path             string
	Params           map[string]interface{}
	JsonBody         map[string]interface{}
	Body             []byte //decryptBody
	ChangeOriginal   bool
	Web              bool
	Encrypted        bool
	Forward          bool
	EndPoint         string
	SearchPath       string
	SearchKey        string
}

var (
	ProxyIp     = "127.0.0.1"
	ProxyDomain = map[string]string{
		"music.163.com":            "59.111.181.35",
		"interface.music.163.com":  "59.111.181.35",
		"interface3.music.163.com": "59.111.181.35",
		"apm.music.163.com":        "59.111.181.35",
		"apm3.music.163.com":       "59.111.181.35",
		"m.jd.com":                 "111.13.29.202",
		"api.m.jd.com":             "111.13.149.100",
		"115.com":                  "119.23.87.59",
	}
	HostDomain = map[string]string{
		"music.163.com":           "59.111.181.35",
		"interface.music.163.com": "59.111.181.35",
		"m.jd.com":                "111.13.29.202",
		"api.m.jd.com":            "111.13.149.100",
		"115.com":                 "119.23.87.59",
	}
)
