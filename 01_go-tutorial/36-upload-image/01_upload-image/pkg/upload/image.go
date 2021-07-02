package upload

import (
	"example/uploadimage/pkg/file"
	"example/uploadimage/pkg/setting"
	"example/uploadimage/pkg/util"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

/*

ImagePrefixUrl = http://127.0.0.1:8000
ImageSavePath = upload/images/

RuntimeRootPath = runtime/

*/
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetFullImagePath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func GetImageName(name string) string {
	//get ext
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func checkImageExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func checkImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println("checkImageSize err", err)
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("CheckImage os.Getwd err:%v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("CheckImage IsNotExistMkDir err: %v ", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
