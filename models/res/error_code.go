package res

import (
	"blog_server/global"
	"encoding/json"
	"os"
)

type ErrorCode int

const (
	SettingsError    ErrorCode = 1001
	ParamsError      ErrorCode = 1002
	FileSizeExceeded ErrorCode = 1010
)

type ErrorMap map[ErrorCode]string

var ErrMap = ErrorMap{}

func ReadErrorCodeJson() map[ErrorCode]string {
	// 读取json文件
	byteData, err := os.ReadFile("models/res/error_code.json")
	if err != nil {
		global.Log.Fatal(err)
		return nil
	}

	err = json.Unmarshal(byteData, &ErrMap)
	if err != nil {
		global.Log.Error(err)
		return nil
	}
	return ErrMap
}
