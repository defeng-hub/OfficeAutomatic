package utils

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"os"
	"path"
	"strconv"
	"time"
)

func ExportFile(houzhui string) (filepath string, url string, err error) {
	day := time.Now().Format("20060102")
	err = os.MkdirAll("./uploads/excel/"+day, os.ModePerm)
	if err != nil {
		return "", "", err
	}
	param := strconv.Itoa(int(time.Now().Unix()))

	url = global.GVA_CONFIG.System.RouterPrefix + "/uploads/excel/" + day + "/" + param + "-" + "export" + houzhui
	filepath = path.Join("./uploads/excel/"+day, param+"-"+"export"+houzhui)
	return filepath, url, err
}
