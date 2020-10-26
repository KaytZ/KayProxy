package test

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}
func TestTime(t *testing.T) {
	cellIntVal := strings.Replace(strings.Replace("/Date(1603209600000-0000)/", "/Date(", "", -1), "000-0000)/", "", -1)
	t.Log(cellIntVal)
	i, _ := strconv.ParseInt(cellIntVal, 10, 64)
	t.Log(i)
	datetime := time.Unix(i, 0).Format("2006-01-02")
	t.Log(datetime)
}
func TestReg(t *testing.T) {
	flysnowRegexp := regexp.MustCompile(`(\[.*?\])`)
	rexExec := regexp.MustCompile(`\[(.*)000,(.*),"(.*)"\]`)
	params := flysnowRegexp.FindStringSubmatch("[1576598400000,19.90,\"\"],[1576684800000,19.90,\"\"],[1576771200000,19.90,\"\"],[1576857600000,19.90,\"\"],[1576944000000,19.90,\"\"],[1577030400000,19.90,\"\"],[1577116800000,19.90,\"\"],[1577203200000,19.90,\"\"],[1577289600000,19.90,\"\"],[1577376000000,19.90,\"\"],[1577462400000,19.90,\"\"],[1577548800000,19.90,\"\"],[1577635200000,19.90,\"\"],[1577721600000,19.90,\"\"],[1577808000000,19.90,\"\"],[1577894400000,19.90,\"\"],[1577980800000,19.90,\"\"],[1578067200000,19.90,\"\"],[1578153600000,19.90,\"\"],[1578240000000,19.90,\"\"],[1578326400000,19.90,\"\"],[1578412800000,19.90,\"\"],[1578499200000,19.50,\"\"],[1578585600000,19.50,\"\"],[1578672000000,19.50,\"\"],[1578758400000,19.50,\"\"],[1578844800000,19.50,\"\"],[1578931200000,19.90,\"\"],[1579017600000,19.90,\"\"],[1579104000000,19.90,\"\"],[1579190400000,19.90,\"\"],[1579276800000,19.90,\"\"],[1579363200000,19.90,\"\"],[1579449600000,19.90,\"\"],[1579536000000,19.90,\"\"],[1579622400000,19.90,\"\"],[1579708800000,19.90,\"\"],[1579795200000,19.90,\"\"],[1579881600000,19.90,\"\"],[1579968000000,19.50,\"\"],[1580054400000,19.50,\"\"],[1580140800000,19.50,\"\"],[1580227200000,19.50,\"\"],[1580313600000,19.50,\"\"],[1580400000000,19.90,\"\"],[1580486400000,19.90,\"\"],[1580572800000,19.90,\"\"],[1580659200000,19.90,\"\"],[1580745600000,19.90,\"\"],[1580832000000,19.90,\"\"],[1580918400000,19.90,\"\"],[1581004800000,18.90,\"购买1件,plus价格18.9\"],[1581091200000,19.90,\"\"],[1581177600000,19.90,\"\"],[1581264000000,19.90,\"\"],[1581350400000,19.90,\"\"],[1581436800000,19.90,\"\"],[1581523200000,19.90,\"\"],[1581609600000,19.90,\"\"],[1581696000000,19.90,\"\"],[1581782400000,21.9000,\"\"],[1581868800000,21.9000,\"\"],[1581955200000,21.90,\"\"],[1582041600000,21.90,\"\"],[1582128000000,21.90,\"\"],[1582214400000,21.90,\"\"],[1582300800000,21.90,\"\"],[1582387200000,21.90,\"\"],[1582473600000,21.90,\"\"],[1582560000000,21.90,\"\"],[1582646400000,21.90,\"\"],[1582732800000,21.90,\"\"],[1582819200000,21.90,\"\"],[1582905600000,21.90,\"\"],[1582992000000,21.90,\"\"],[1583078400000,21.90,\"\"],[1583164800000,21.90,\"\"],[1583251200000,21.90,\"\"],[1583337600000,21.90,\"\"],[1583424000000,21.90,\"\"],[1583510400000,21.90,\"\"],[1583596800000,21.90,\"\"],[1583683200000,21.90,\"\"],[1583769600000,21.90,\"\"],[1583856000000,21.90,\"\"],[1583942400000,21.90,\"\"],[1584028800000,21.90,\"\"],[1584115200000,21.90,\"\"],[1584201600000,21.90,\"\"],[1584288000000,19.90,\"\"],[1584374400000,19.90,\"\"],[1584460800000,19.90,\"\"],[1584547200000,23.9000,\"\"],[1584633600000,23.9000,\"\"],[1584720000000,23.9000,\"\"],[1584806400000,23.9000,\"\"],[1584892800000,23.9000,\"\"],[1584979200000,19.9000,\"\"],[1585065600000,19.90,\"购买1件,plus价格19.9\"],[1585152000000,23.90,\"\"],[1585238400000,23.90,\"\"],[1585324800000,23.90,\"\"],[1585411200000,23.90,\"\"],[1585497600000,23.90,\"\"],[1585584000000,19.90,\"\"],[1585670400000,19.90,\"\"],[1585756800000,19.90,\"\"],[1585843200000,19.90,\"\"],[1585929600000,19.90,\"\"],[1586016000000,19.90,\"\"],[1586102400000,19.90,\"\"],[1586188800000,23.90,\"\"],[1586275200000,23.90,\"\"],[1586361600000,23.90,\"\"],[1586448000000,23.90,\"\"],[1586534400000,23.90,\"\"],[1586620800000,23.90,\"\"],[1586707200000,23.90,\"\"],[1586793600000,19.90,\"\"],[1586880000000,19.90,\"\"],[1586966400000,19.90,\"\"],[1587052800000,19.90,\"\"],[1587139200000,19.90,\"\"],[1587225600000,19.90,\"\"],[1587312000000,19.90,\"\"],[1587398400000,19.90,\"\"],[1587484800000,19.90,\"\"],[1587571200000,23.90,\"\"],[1587657600000,23.90,\"\"],[1587744000000,23.90,\"\"],[1587830400000,23.90,\"\"],[1587916800000,23.90,\"\"],[1588003200000,23.90,\"\"],[1588089600000,23.90,\"\"],[1588176000000,23.90,\"\"],[1588262400000,23.90,\"\"],[1588348800000,23.90,\"\"],[1588435200000,23.90,\"\"],[1588521600000,23.90,\"\"],[1588608000000,23.90,\"\"],[1588694400000,23.90,\"\"],[1588780800000,23.90,\"\"],[1588867200000,23.90,\"\"],[1588953600000,23.90,\"\"],[1589040000000,15.90,\"47.76元（合15.92元/件）\"],[1589126400000,23.90,\"\"],[1589212800000,23.90,\"\"],[1589299200000,23.90,\"\"],[1589385600000,23.90,\"\"],[1589472000000,23.90,\"\"],[1589558400000,23.90,\"\"],[1589644800000,23.90,\"\"],[1589731200000,23.90,\"\"],[1589817600000,23.90,\"\"],[1589904000000,23.90,\"\"],[1589990400000,23.90,\"\"],[1590076800000,23.90,\"\"],[1590163200000,23.90,\"\"],[1590249600000,23.90,\"\"],[1590336000000,19.90,\"购买1件,plus价格19.9\"],[1590422400000,23.90,\"\"],[1590508800000,23.90,\"\"],[1590595200000,23.90,\"\"],[1590681600000,23.90,\"\"],[1590768000000,23.90,\"\"],[1590854400000,23.90,\"\"],[1590940800000,23.9000,\"\"],[1591027200000,23.9000,\"\"],[1591113600000,23.9000,\"\"],[1591200000000,19.90,\"购买1件,plus价格19.9\"],[1591286400000,23.90,\"\"],[1591372800000,19.90,\"购买1件,plus价格19.9\"],[1591459200000,23.90,\"\"],[1591545600000,23.90,\"\"],[1591632000000,23.90,\"\"],[1591718400000,23.90,\"\"],[1591804800000,23.90,\"\"],[1591891200000,23.90,\"\"],[1591977600000,23.90,\"\"],[1592064000000,19.9000,\"\"],[1592150400000,19.90,\"\"],[1592236800000,19.90,\"\"],[1592323200000,19.90,\"\"],[1592409600000,19.90,\"\"],[1592496000000,19.90,\"\"],[1592582400000,19.90,\"\"],[1592668800000,23.9000,\"\"],[1592755200000,23.90,\"\"],[1592841600000,23.90,\"\"],[1592928000000,23.90,\"\"],[1593014400000,23.90,\"\"],[1593100800000,23.90,\"\"],[1593187200000,23.90,\"\"],[1593273600000,23.90,\"\"],[1593360000000,23.90,\"\"],[1593446400000,23.90,\"\"],[1593532800000,23.90,\"\"],[1593619200000,23.90,\"\"],[1593705600000,23.90,\"\"],[1593792000000,19.90,\"\"],[1593878400000,19.90,\"\"],[1593964800000,19.90,\"\"],[1594051200000,19.90,\"\"],[1594137600000,19.90,\"\"],[1594224000000,19.90,\"\"],[1594310400000,19.90,\"\"],[1594396800000,19.90,\"\"],[1594483200000,19.90,\"\"],[1594569600000,19.90,\"\"],[1594656000000,19.90,\"\"],[1594742400000,19.50,\"购买1件,plus价格19.5\"],[1594828800000,19.90,\"\"],[1594915200000,19.90,\"\"],[1595001600000,19.90,\"\"],[1595088000000,19.90,\"\"],[1595174400000,19.90,\"\"],[1595260800000,19.90,\"\"],[1595347200000,19.50,\"\"],[1595433600000,19.50,\"\"],[1595520000000,19.50,\"\"],[1595606400000,19.50,\"\"],[1595692800000,19.9000,\"\"],[1595779200000,19.90,\"\"],[1595865600000,19.90,\"\"],[1595952000000,19.90,\"\"],[1596038400000,19.90,\"\"],[1596124800000,19.90,\"\"],[1596211200000,19.90,\"\"],[1596297600000,19.90,\"\"],[1596384000000,19.90,\"\"],[1596470400000,19.90,\"\"],[1596556800000,19.90,\"\"],[1596643200000,19.90,\"\"],[1596729600000,19.90,\"\"],[1596816000000,19.90,\"\"],[1596902400000,19.90,\"\"],[1596988800000,19.90,\"\"],[1597075200000,19.90,\"\"],[1597161600000,19.90,\"\"],[1597248000000,19.90,\"\"],[1597334400000,19.90,\"\"],[1597420800000,19.90,\"\"],[1597507200000,19.90,\"\"],[1597593600000,19.90,\"\"],[1597680000000,19.90,\"\"],[1597766400000,19.90,\"\"],[1597852800000,19.90,\"\"],[1597939200000,19.90,\"\"],[1598025600000,19.90,\"\"],[1598112000000,19.90,\"\"],[1598198400000,19.90,\"\"],[1598284800000,19.90,\"\"],[1598371200000,19.90,\"\"],[1598457600000,19.90,\"\"],[1598544000000,19.90,\"\"],[1598630400000,19.90,\"\"],[1598716800000,19.90,\"\"],[1598803200000,19.90,\"\"],[1598889600000,19.90,\"\"],[1598976000000,19.90,\"\"],[1599062400000,19.90,\"\"],[1599148800000,19.90,\"\"],[1599235200000,19.90,\"\"],[1599321600000,19.90,\"\"],[1599408000000,19.90,\"\"],[1599494400000,19.90,\"\"],[1599580800000,19.90,\"\"],[1599667200000,19.90,\"\"],[1599753600000,19.90,\"\"],[1599840000000,19.90,\"\"],[1599926400000,19.90,\"\"],[1600012800000,19.90,\"\"],[1600099200000,19.90,\"\"],[1600185600000,19.90,\"\"],[1600272000000,19.90,\"\"],[1600358400000,19.90,\"\"],[1600444800000,19.90,\"\"],[1600531200000,19.50,\"购买1件,plus价格19.5\"],[1600617600000,19.90,\"\"],[1600704000000,19.90,\"\"],[1600790400000,19.90,\"\"],[1600876800000,19.90,\"\"],[1600963200000,19.90,\"\"],[1601049600000,19.90,\"\"],[1601136000000,19.90,\"\"],[1601222400000,19.90,\"\"],[1601308800000,19.90,\"\"],[1601395200000,19.90,\"\"],[1601481600000,19.90,\"\"],[1601568000000,19.90,\"\"],[1601654400000,19.90,\"\"],[1601740800000,19.90,\"\"],[1601827200000,19.90,\"\"],[1601913600000,19.90,\"\"],[1602000000000,19.90,\"\"],[1602086400000,19.90,\"\"],[1602172800000,19.90,\"\"],[1602259200000,19.90,\"\"],[1602345600000,19.90,\"\"],[1602432000000,19.90,\"\"],[1602518400000,17.50,\"\"],[1602604800000,17.50,\"\"],[1602691200000,17.50,\"\"],[1602777600000,17.50,\"\"],[1602864000000,17.50,\"\"],[1602950400000,17.50,\"\"],[1603036800000,17.50,\"\"],[1603123200000,17.50,\"\"],[1603209600000,17.50,\"\"],[1603296000000,17.9,\"京东秒杀价:17.9\"],[1603388998714,18.90,\"购买1件,plus价格18.9\"],[1603419969964,18.90,\"购买1件,plus价格18.9\"]")

	for index, param := range params {
		t.Log(strconv.Itoa(index) + "  " + param)
		if len(param) > 0 {
			result := rexExec.FindStringSubmatch(param)
			t.Log(result[1])
			i, _ := strconv.ParseInt(result[1], 10, 64)
			t.Log(i)
			date := time.Unix(i, 0).Format("2006-01-02")
			t.Log(date)
			t.Log(result[2])
		}
	}
}
func TestFiex(t *testing.T) {
	v := 3.1415926535
	s1 := strconv.FormatFloat(v, 'f', 2, 32)
	t.Log(s1)
}