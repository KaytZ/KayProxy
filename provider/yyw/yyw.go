package yyw

import (
	"github.com/Kaytz/KayProxy/common"
	"github.com/Kaytz/KayProxy/utils"
	"github.com/bitly/go-simplejson"
	"strings"
)

var (
	path1 = "/user/check_sign"
)

func YywOffline(httpInfo *common.HttpInfo) bool {
	modified := false

	if strings.Contains(httpInfo.EndPoint, path1) {
		jsonBody, err := simplejson.NewJson(utils.ToJsonByte(httpInfo.JsonBody))
		if err != nil {
			return false
		}
		jsonBody.Get("data").Set("is_new_sign", false)
		m, err := jsonBody.Map()
		if err != nil {
			return false
		}
		httpInfo.JsonBody = m
		modified = true
	} else {
		bodyStr := string(httpInfo.Body)
		bodyStr = utils.ReplaceAll(bodyStr, "UDown", "XXXXXXXXX") // 使重定向判断条件失效
		clearJS := `<script type="text/javascript">
    $(function(){
        function actionSheetAction(index) {
            if (index == 4) { return; }
            $.ajax({
                url: '/web/lixian/?ct=lixian&ac=task_clear',
                dataType: 'json',
                data: {flag: index},
                type: 'POST',
                cache: false,
                success: function (r) {
                    window.location.reload();
                },
                error: function(){
                    console.log('清空失败');
                }
            });
        }
        function clearAction() { 
            OOFJS.common.actionSheet('选择清空的操作', ['清空已完成任务', '清空全部任务', '清空失败任务', '清空进行中任务'], actionSheetAction);
        }
        setTimeout(function(){OOFJS.common.addRightBarItem('清空', clearAction);}, 200);
    });
    </script>`
		bodyStr = utils.ReplaceAll(bodyStr, "</body>", clearJS+"\n</body>") // 注入清空任务相关JS
		httpInfo.Body = []byte(bodyStr)
		httpInfo.ChangeOriginal = true
		modified = true
	}
	return modified
}
