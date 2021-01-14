package util

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"
)

type Params map[string]string

// map本来已经是引用类型了，所以不需要 *Params
//@author sam@2020-06-18 10:06:38
func (p Params) SetString(k, s string) Params {
	p[k] = s
	return p
}

func (p Params) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p Params) SetInt64(k string, i int64) Params {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p Params) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// 判断key是否存在
func (p Params) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}
//-----------------------------------------------------------------------------------------------------------------
/*
<xml><return_code><![CDATA[FAIL]]></return_code>
<return_msg><![CDATA[商户号该产品权限未开通，请前往商户平台>产品中心检查后重试]]></return_msg>
</xml>

(1)<xml>           (StartElement xml)==>{Name:{Space: Local:xml} Attr:[]}
(2)<return_code>   (StartElement return_code)==>{Name:{Space: Local:return_code} Attr:[]}
(3)<![CDATA[FAIL]]> (CharData FAIL)==>[70 65 73 76]
(4)</return_code>  (EndElement return_code)==>{Name:{Space: Local:return_code}}
(5)                (CharData  "\n") ==>[10]
(6)<return_msg>    (StartElement return_msg) ==> {Name:{Space: Local:return_msg} Attr:[]}
(7)<![CDATA[商户号该产品权限未开通，请前往商户平台>产品中心检查后重试]]>  (CharData ...) ==>[229 149 134 230 136 183 229 143 183 232 175 165 228 186 ...]
(8)</return_msg>    (EndElement return_msg)==>{Name:{Space: Local:return_msg}}
(9)                (CharData  "\n") ==>[10]
(10)</xml>        (EndElement xml)   ==>{Name:{Space: Local:xml}}

*/


//xml转map
//todo 参照go web编程第7章,这里只是针对微信接口那种简易xml格式进行解析的额，复杂的还得进一步书写,比如发文回调就不适用
//@author sam@2020-06-18 11:42:37
func XmlToMap(xmlStr string) Params {
	//生命一个map,存放xml数据
	params := make(Params)
	//根据给定的XML数据生成相应的解码器
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	var (
		key   string
		value string
	)
	//每进行一次迭代，就从解码器里面获取一个token
	//token是针对<>进行一个个元素迭代的额
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		//fmt.Printf("%+v\r\n",t)
		//检查token的类型
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
			//fmt.Println("StartElement",key)
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
			//fmt.Printf("%q",content)
		case xml.EndElement://结束标签
			//不做任何处理
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
		//fmt.Println("------------------------")
	}
	return params
}

//将Map转为xml
//todo 注意所有传递的节点值务必使用CDATA保护，以防止特殊字符被误解析了
//@author sam@2020-06-18 11:40:59
func MapToXml(params Params) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)
	return buf.String()
}




