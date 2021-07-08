package export

import "example/gen-excel/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.ExcelSetting.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.ExcelSetting.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.ExcelSetting.RuntimeRootPath + GetExcelPath()
}
