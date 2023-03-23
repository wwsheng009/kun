package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	jsoniter "github.com/json-iterator/go"
)

// DD The DD function dumps the given variables and ends execution of the script
func DD(values ...interface{}) {
	Dump(values...)
	os.Exit(0)
}

// Dump The Dump function dumps the given variables:
func Dump(values ...interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 4
	for _, v := range values {
		var res interface{}
		if err, ok := v.(error); ok {
			fmt.Printf("%s\n", err.Error())
			continue
		}

		// to Map
		if value, ok := v.(interface{ Map() map[string]interface{} }); ok {
			v = value.Map()
		}

		txt, err := jsoniter.Marshal(v)
		if err != nil {
			fmt.Printf("%#v\n%s\n", v, err)
			continue
		}
		jsoniter.Unmarshal(txt, &res)
		// s, _ := f.Marshal(res)
		// fmt.Printf("%s\n", s)
		s, _ := UnescapeJsonMarshal(res)
		fmt.Printf("%s", s) //返回值已带有\n
	}
}

// https://blog.csdn.net/zhuhanzi/article/details/106156174
// 修正console.log输出html符号异常
func UnescapeJsonMarshal(jsonRaw interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	//带有缩进的格式化
	encoder.SetIndent("", "    ")
	err := encoder.Encode(jsonRaw)
	return buffer.Bytes(), err
}
