package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

var count int = 0
var totalD int = math.MaxInt64
var pathD []Poi = []Poi{}

type Near struct {
	Name     string
	Location string
	Distance float64
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func Nearest(location string, destinations []Poi) []string {
	nearest := []Near{}
	//return neares 2
	for _, des := range destinations {
		arr := strings.Split(location, ",")
		arrd := strings.Split(des.Location, ",")

		x0, _ := strconv.ParseFloat(arr[0], 64)
		x1, _ := strconv.ParseFloat(arr[1], 64)

		y0, _ := strconv.ParseFloat(arrd[0], 64)
		y1, _ := strconv.ParseFloat(arrd[1], 64)

		d := Distance(x1, x0, y1, y0)

		// x := math.Abs(x0 - x1)
		// y := math.Abs(y0 - y1)

		// d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		nearest = append(nearest, Near{Name: des.Name, Location: des.Location, Distance: d})
		//fmt.Println(d)
	}

	sort.Slice(nearest, func(i, j int) bool {
		return nearest[i].Distance < nearest[j].Distance
	})
	// fmt.Println("nr", nearest)

	returnArr := []string{}
	for i, n := range nearest {
		if i < 1 {
			fmt.Println(location, n)
			returnArr = append(returnArr, n.Location)

		} else {
			break
		}
	}
	return returnArr
}

func PathTrack(starts []Poi, destinations []Poi) {
	//查到2个距离最近的点
	n := Nearest(starts[len(starts)-1].Location, destinations)
	// fmt.Println("near", n)
	for i, des := range destinations {

		//for _, nr := range n {
		nr := n[0]
		if des.Location == nr {
			newStarts := append(starts, des)
			newDestinations := []Poi{}
			newDestinations = append(newDestinations, destinations[0:i]...)
			newDestinations = append(newDestinations, destinations[i+1:]...)

			if len(newDestinations) != 0 {
				PathTrack(newStarts, newDestinations)
			} else {
				//规划完的路径
				//fmt.Println(newStarts)
				total := 0
				for i, _ := range newStarts {
					if i != len(newStarts)-1 {
						d, err := executePathPlan(newStarts[i].Location, newStarts[i+1].Location)
						if err != nil {
							fmt.Println(err)
						} else {
							count++
							total += d
						}
					}
				}
				fmt.Println("calculation...", newStarts, total)
				if total < totalD {
					totalD = total
					pathD = newStarts
				}
				//fmt.Println("total:", total)
			}
		} else {
			continue
		}
		//}
	}

}

type Poi struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func ExeCute() {

	//executePathPlan("116.315761,39.990097", "116.343438,40.043229")

	startPoi := []Poi{{"上海市浦东新区张江", "121.612233,31.209309"}}

	hotelPoi := []Poi{
		{"宜高快捷酒店", "121.341749,30.714454"},
		{"诚意旅馆", "121.325654,30.896209"},
		{"龙纳商务宾馆", "121.279681,30.699389"},
		{"兴萍旅馆", "121.321343,30.727819"},
		{"荣都客房", "121.324554,30.728885"},
		{"卡福伦快捷宾馆", "121.340638,30.711378"},
		{"祖杰怡居精品酒店", "121.319010,30.729829"},
		{"99优选酒店", "121.165604,30.892433"},
		{"廊苑商务宾馆", "121.192349,30.787375"},
		{"锦江农家饭庄", "121.150816,30.786995"},
		{"静洁宾馆", "121.333261,30.893259"},
		{"地园旅馆", "121.017958,30.884126"},
		{"易佰连锁酒店", "121.168163,30.893770"},
		{"梦农旅馆", "121.167591,30.893850"},
		{"新明客房", "121.332851,30.725927"},
		{"香秀宾馆", "121.332339,30.726884"},
		{"苏源客房", "121.332918,30.725619"},
		{"白牛宾馆", "121.023423,30.892945"},
		{"云英旅馆", "121.076836,30.866717"},
		{"宇琴旅馆", "121.076721,30.867369"},
		{"枫之林浴场", "121.020422,30.893238"},
		{"新义农庄", "121.025032,30.845090"},
		{"永和旅馆", "121.017009,30.884063"},
		{"桥枫旅馆", "121.018965,30.884704"},
		{"好再来旅店", "121.017762,30.884095"},
		{"东亚旅馆", "121.072590,30.865155"},
		{"闵枫旅馆", "121.072627,30.865103"},
		{"庆怡宾馆", "121.015358,30.885758"},
		{"上海豪鹰宾馆", "121.014808,30.890512"},
		{"新长岭大酒店", "121.014162,30.898840"},
		{"桥英旅馆", "121.026236,30.892365"},
		{"山麓酒店", "121.017945,30.888129"},
		{"金皇朝轻奢酒店", "121.014366,30.891412"},
		{"悠怡精品酒店", "121.014943,30.893713"},
		{"景宴旅馆", "121.016018,30.896504"},
		{"馨盏民宿", "121.013835,30.886405"},
		{"君临旅馆", "121.078184,30.869245"},
		{"恩佳旅馆", "121.076594,30.867504"},
		{"富君客房", "121.013798,30.888396"},
		{"兴雅旅馆", "121.082281,30.868087"},
		{"和月宾馆", "121.076722,30.872562"},
		{"100易佰酒店", "121.015158,30.887617"},
		{"三桥别院", "121.016021,30.888422"},
		{"99旅馆连锁", "121.013601,30.888057"},
		{"荷风嬉鱼度假村", "121.017225,30.922266"},
		{"清水旅馆", "121.014319,30.889544"},
		{"芳心园", "121.250582,30.821053"},
		{"梦浓客房", "121.235062,30.825989"},
		{"鑫光客房", "121.246780,30.815219"},
		{"美婷旅馆", "121.160916,30.894640"},
		{"99旅馆连锁", "121.162944,30.895825"},
		{"永城旅馆", "121.161482,30.892114"},
		{"辰奕客房", "121.318622,30.726582"},
		{"安驿旅馆", "121.317468,30.729020"},
		{"龙翔宾馆", "121.344660,30.753932"},
		{"三音酒店", "121.336941,30.752878"},
		{"金山汤泉", "121.343895,30.747763"},
		{"喆啡酒店", "121.352423,30.730901"},
		{"小江南宾馆", "121.343291,30.707872"},
		{"蓝驿宾馆", "121.364234,30.763366"},
		{"海泊漫居", "121.374863,30.731141"},
		{"海趣客房", "121.339711,30.717346"},
		{"湘港大酒店", "121.337389,30.717711"},
		{"缘？渔歌", "121.377418,30.734235"},
		{"林泉渔家傲", "121.376637,30.733287"},
		{"杭州湾酒店", "121.375526,30.731968"},
		{"灿宇宾馆", "121.450266,31.214308"},
		{"上海锦山酒店", "121.182820,30.885887"},
		{"碧水云天", "121.172970,30.886623"},
		{"MT主题酒店", "121.166743,30.885771"},
		{"维也纳国际酒店", "121.165749,30.885902"},
		{"豪庭时尚宾馆/豪庭宾馆", "121.162533,30.885963"},
		{"99旅馆连锁", "121.161922,30.884656"},
		{"富苑酒店", "121.340954,30.853796"},
		{"旅馆", "121.327338,30.729452"},
		{"金居客房", "121.319026,30.725120"},
		{"钱龙大酒店", "121.318539,30.733079"},
		{"开馨旅馆", "121.313051,30.725060"},
		{"钱圩宾馆", "121.250485,30.777476"},
		{"新街旅馆", "121.247940,30.776900"},
		{"鸿卫旅馆", "121.308225,30.725577"},
		{"商务宾馆", "121.320852,30.736192"},
		{"驿恋宾馆", "121.314523,30.746848"},
		{"尚客优连锁酒店", "121.317769,30.747779"},
		{"君逑时尚宾馆", "121.323416,30.725148"},
		{"君鳅时尚宾馆", "121.323416,30.725148"},
		{"速8酒店", "121.322981,30.726987"},
		{"青城旅馆", "121.322893,30.727279"},
		{"浦江连锁旅店", "121.322117,30.729507"},
		{"古城客房", "121.322007,30.729815"},
		{"绿竹巷主题客栈", "121.341970,30.741991"},
		{"绿竹巷主题客栈", "121.341970,30.741991"},
		{"古井坊", "121.376255,30.732799"},
		{"百发客房", "121.341066,30.712684"},
		{"锦江之星", "121.340279,30.713765"},
		{"宜高快捷酒店", "121.341749,30.714454"},
		{"枫泾商务宾馆", "121.013148,30.889523"},
		{"朱行大酒店", "121.341590,30.849301"},
		{"水生宾馆", "121.340816,30.849117"},
		{"青檐艺宿", "121.164138,30.776018"},
		{"新世纪商行", "121.192029,30.785605"},
		{"景乐客房", "121.191871,30.789538"},
		{"舒乐旅馆", "121.187139,30.785165"},
		{"街道招待所", "121.192175,30.785729"},
		{"博海农艺休闲馆", "121.182636,30.817421"},
		{"园宿听风塘", "121.164138,30.776018"},
		{"江南莲湘", "121.164138,30.776018"},
		{"函七旅馆", "121.173135,30.782802"},
		{"立期酒家", "121.189314,30.786388"},
		{"美静旅馆", "121.315144,30.724636"},
		{"东街旅馆", "121.314893,30.724531"},
		{"卫康旅馆", "121.309089,30.723446"},
		{"尚客优选酒店", "121.339552,30.852480"},
		{"海滩大酒店", "121.340195,30.707355"},
		{"君悦宾馆", "121.170049,30.892979"},
		{"东乐缘住宿", "121.169572,30.893145"},
		{"智尚酒店", "121.169214,30.892989"},
		{"豪毅宾馆", "121.350218,30.758821"},
		{"7天连锁酒店", "121.348314,30.726160"},
		{"布丁酒店", "121.344902,30.725399"},
		{"速8酒店", "121.338012,30.724584"},
		{"格林豪泰上海金山城市沙滩商务酒店", "121.338168,30.724109"},
		{"茵阁假日酒店", "121.334954,30.723756"},
		{"海鸥大厦", "121.352139,30.729227"},
		{"茉雅酒店", "121.334779,30.723777"},
		{"安母酒店", "121.339400,30.747657"},
		{"格林豪泰快捷酒店", "121.340835,30.747749"},
		{"易佰良品酒店", "121.337393,30.746509"},
		{"和喜快捷宾馆", "121.345969,30.718371"},
		{"普吉汤", "121.344281,30.716983"},
		{"普吉汤", "121.344281,30.716983"},
		{"麗枫酒店", "121.344816,30.717492"},
		{"鸿煌宾馆", "121.169608,30.893901"},
		{"金粮客房", "121.236552,30.825912"},
		{"心苑旅馆", "121.182775,30.831847"},
		{"顺达浴场", "121.178169,30.828250"},
		{"南文宾馆", "121.177223,30.824541"},
		{"惠意客房", "121.231896,30.831311"},
		{"新溪民宿", "121.239779,30.827096"},
		{"钱吕宾馆", "121.181108,30.831447"},
		{"99旅馆连锁", "121.342999,30.730213"},
		{"神仙居沐浴会所", "121.343152,30.729762"},
		{"海焱快捷宾馆", "121.342378,30.733228"},
		{"易佰旅店", "121.344300,30.724420"},
		{"赣川旅馆", "121.339468,30.885435"},
		{"碧丽宫大酒店", "121.323650,30.881589"},
		{"城华旅馆", "121.321438,30.880813"},
		{"万尚客快捷宾馆", "121.320812,30.880827"},
		{"振兴旅馆", "121.319612,30.880038"},
		{"汶庭酒店", "121.421436,30.795097"},
		{"枫憬酒店", "121.345797,30.735417"},
		{"莲纳酒家旅馆", "121.249135,30.777012"},
		{"锦水客房", "121.249653,30.776879"},
		{"智选假日酒店", "121.336359,30.753609"},
		{"维也纳酒店（金山新城店）", "121.337800,30.750910"},
		{"颐家精选宾馆", "121.338206,30.750933"},
		{"丰家旅馆", "121.363074,30.762600"},
		{"净馨一站式民宿", "121.375444,30.731873"},
		{"听海精品民宿", "121.337047,30.717690"},
		{"好久不见", "121.376255,30.732800"},
		{"半朵悠莲", "121.356657,30.761077"},
		{"老井客栈", "121.376070,30.733190"},
		{"格林豪泰快捷酒店", "121.335285,30.754057"},
		{"尚客优连锁酒店", "121.339400,30.747657"},
		{"荣乐宾馆", "121.373094,30.765576"},
		{"梦旭旅馆", "121.372799,30.765615"},
		{"天绮酒店", "121.349887,30.730875"},
		{"维也纳国际酒店", "121.332857,30.726654"},
		{"大屋里", "121.364594,30.735725"},
		{"湖悦旅馆", "121.376361,30.733043"},
		{"一片叶子的故事主题民宿", "121.376361,30.733043"},
		{"萍聚阁", "121.376361,30.733043"},
		{"蓝心舍客栈", "121.376361,30.733043"},
		{"渔舍客栈", "121.376361,30.733043"},
		{"琴轩居", "121.376361,30.733043"},
		{"藤缘", "121.376361,30.733043"},
		{"玲珑村舍", "121.376361,30.733043"},
		{"渔家客栈", "121.376361,30.733043"},
		{"春峰旅馆", "121.351506,30.791843"},
		{"惠晨旅馆", "121.351506,30.791843"},
		{"璞园", "121.338872,30.774028"},
		{"卡福伦宾馆", "121.340638,30.711378"},
		{"金山宾馆", "121.343742,30.714950"},
		{"骏怡连锁酒店", "121.324800,30.723595"},
		{"非酷不住新Fun潮宿", "121.167966,30.906695"},
		{"红枫缘浴场", "121.014837,30.890322"},
		{"欣博大浴场", "121.371532,30.763904"},
		{"翔茂花园度假酒店", "121.304131,30.888444"},
		{"永昌宾馆", "121.317468,30.887140"},
		{"永新旅社", "121.323130,30.878964"},
		{"魅之梦旅馆", "121.328725,30.882633"},
		{"山悦旅馆", "121.322087,30.880706"},
		{"复兴旅社", "121.319546,30.880670"},
		{"99优选酒店", "121.317658,30.874252"},
		{"玺悦旅馆", "121.340462,30.848663"},
		{"芭蕉客栈", "121.328339,30.895842"},
		{"歆南宾馆", "121.323556,30.881549"},
		{"歆南宾馆", "121.323556,30.881549"},
		{"车站旅社", "121.337329,30.847328"},
		{"悠有宾馆", "121.313642,30.887934"},
		{"吉祥宾馆", "121.313652,30.888067"},
		{"如意旅馆", "121.313169,30.888108"},
		{"顺兴旅馆", "121.314290,30.883416"},
		{"怡杰旅馆", "121.314318,30.883513"},
		{"新顺宾馆", "121.313303,30.887785"},
		{"恋依旅馆", "121.239476,30.890910"},
		{"华婷旅馆", "121.237269,30.890781"},
		{"亭鑫旅馆", "121.312932,30.890088"},
		{"妙泉", "121.321199,30.883493"},
		{"颐家宾馆", "121.320265,30.884861"},
		{"硕茂旅馆", "121.240220,30.894241"},
		{"老地方旅馆", "121.240560,30.894229"},
		{"星连心旅馆", "121.342890,30.846930"},
		{"朱港旅馆", "121.344928,30.846711"},
		{"凯乐旅馆", "121.338038,30.847899"},
		{"高家旅馆", "121.344567,30.846769"},
		{"上引国际酒店", "121.349171,30.837247"},
		{"山阳田园度假村", "121.347440,30.786409"},
		{"山悦商务宾馆", "121.330026,30.882055"},
		{"锦江都城酒店", "121.371560,30.728723"},
		{"途安假日酒店", "121.321088,30.890164"},
		{"隆亭旅馆", "121.320539,30.879884"},
		{"相福浴室", "121.337603,30.845492"},
		{"嘉都酒店", "121.176750,30.897120"},
		{"茉雅精品酒店", "121.164760,30.897130"},
		{"绿洋水汇", "121.177813,30.897198"},
		{"青皮树酒店", "121.303482,30.733543"},
		{"贝壳酒店", "121.331633,30.745906"},
		{"铠石宾馆", "121.337844,30.723116"},
		{"晥沪宾馆", "121.339621,30.713123"},
		{"志宇旅馆", "121.340021,30.721628"},
		{"新世佳快捷客房", "121.337481,30.724634"},
		{"上海市金山区祥正旅馆", "121.340478,30.714149"},
		{"欧城宾馆", "121.334944,30.729170"},
		{"征伟宾馆", "121.335744,30.729893"},
		{"浦江快捷连锁酒店", "121.328849,30.724574"},
		{"如家酒店", "121.349899,30.730104"},
		{"海立方国际酒店", "121.349778,30.729862"},
		{"易佰良品酒店", "121.346183,30.729219"},
		{"贝壳酒店", "121.342369,30.729059"},
		{"尚客优酒店", "121.342069,30.729141"},
		{"城市便捷酒店", "121.351371,30.730270"},
		{"柏丽艾尚酒店", "121.351254,30.730224"},
		{"忆Hotel酒店", "121.332906,30.727072"},
		{"金汤浴", "121.359972,30.729100"},
		{"国标客房", "121.337100,30.709239"},
		{"文商时尚宾馆", "121.160316,30.895151"},
		{"华风旅馆", "121.159630,30.894234"},
		{"星盛客房", "121.159375,30.894249"},
		{"居馨旅馆", "121.159304,30.894267"},
		{"如家酒店", "121.345440,30.711180"},
		{"驿宾乔恩宾馆", "121.342980,30.710024"},
		{"上海金山滨海铂骊酒店", "121.343219,30.704926"},
		{"振林旅馆", "121.001709,30.893275"},
		{"板桥客房", "121.285537,30.801748"},
		{"长乐旅馆", "121.014113,30.888284"},
		{"新农宾馆", "121.206750,30.894474"},
		{"富春浴室", "121.205124,30.893616"},
		{"贵发宾馆", "121.077924,30.870822"},
		{"鸿运来商务旅馆", "121.324605,30.720509"},
		{"东方旅馆", "121.324739,30.719836"},
		{"鸿福旅馆", "121.324118,30.722675"},
		{"上海市金山区郝房118宾馆", "121.346960,30.809099"},
		{"双乐旅馆", "121.290162,30.806221"},
		{"云酒店睿柏上海金山张堰镇店", "121.296897,30.798759"},
		{"贝壳酒店", "121.289494,30.803423"},
		{"维也纳智好酒店(松金公路店)", "121.290202,30.804114"},
		{"板桥客房", "121.285141,30.801342"},
		{"留溪旅社", "121.286329,30.802660"},
		{"张堰商务宾馆", "121.287932,30.804177"},
		{"齐美旅馆", "121.291212,30.807038"},
		{"金山假日酒店", "121.310571,30.836092"},
		{"东港旅社", "121.412062,30.795761"},
		{"上海中洪大院酒店", "121.018004,30.922452"},
		{"汉姆连锁酒店", "121.166560,30.896270"},
		{"唐沐浴室", "121.156081,30.897062"},
		{"7天连锁酒店", "121.160332,30.895851"},
		{"云甸旅馆", "121.339212,30.847694"},
		{"格林豪泰酒店", "121.158088,30.891152"},
	}
	PathTrack(startPoi, hotelPoi)
	fmt.Println("最短路径:", pathD, "共", totalD, "米", "共运算", count, "次")
}

type PathPlan struct {
	Status string `json:"status"`
	Route  *Route `json:"route"`
}

type Route struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Paths       []*Path `json:"paths"`
}

type Path struct {
	Distance string `json:"distance"`
}

func executePathPlan(origin, destination string) (int, error) {
	_url := "https://restapi.amap.com/v5/direction/walking"

	params := &url.Values{}
	params.Set("key", "11ef02e45ea082f955f20e5d455cb722")
	params.Set("origin", origin)
	params.Set("destination", destination)

	req, _ := http.NewRequest("GET", _url, nil)
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.URL.RawQuery = params.Encode()

	res, err := http.DefaultClient.Do(req)
	pathPlan := &PathPlan{}
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		defer res.Body.Close()
		resp, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(resp, pathPlan)
		//fmt.Println(string(resp))

		if pathPlan.Status == "1" {
			if len(pathPlan.Route.Paths) > 0 {
				distance, _ := strconv.Atoi(pathPlan.Route.Paths[0].Distance)
				return distance, nil
			} else {
				return 0, errors.New("error distance")
			}
		} else {
			fmt.Println(pathPlan, params)
			return 0, errors.New("error distance")
		}
	}
}
