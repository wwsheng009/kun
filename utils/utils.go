package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
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

		if err, ok := v.(error); ok {
			color.Red(err.Error())
			continue
		}

		switch v.(type) {

		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			color.Cyan(fmt.Sprintf("%v", v))
			continue

		case string, []byte:
			color.Green(fmt.Sprintf("%s", v))
			continue

		default:
			var res interface{}
			txt, err := jsoniter.Marshal(v)
			if err != nil {
				color.Red(err.Error())
				continue
			}

			jsoniter.Unmarshal(txt, &res)
			// s, _ := f.Marshal(res)
			s, _ := UnescapeJsonMarshal(res)
			fmt.Printf("%s\n", s)
		}
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
