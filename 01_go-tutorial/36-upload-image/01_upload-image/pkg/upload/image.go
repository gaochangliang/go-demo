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

PrefixUrl = http://127.0.0.1:8000
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

	ext := path.Ext(name)                     //获取文件后缀 e.g. png
	fileName := strings.TrimSuffix(name, ext) //文件名去掉后缀
	fileName = util.EncodeMD5(fileName)       //返回名字加密后16进制的hash

	return fileName + ext
}

func GetImageFullPath() string {
	//log.Println("setting.AppSetting.RuntimeRootPath --", setting.AppSetting.RuntimeRootPath)	   runtime/
	//log.Println("GetImagePath() - ", GetImagePath())		 upload/images/
	return setting.AppSetting.RuntimeRootPath + GetImagePath() //通常由根目录下文件夹加上存储目录  runtime/ + upload/images/
}

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func CheckImageExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println("checkImageSize err", err)
	}
	return size <= setting.AppSetting.ImageMaxSize
}

//src -> runtime/upload/images/
func CheckImage(src string) error {
	dir, err := os.Getwd() //获取项目的工作目录也就是根目录
	//fmt.Println(" CheckImage dir ",dir)
	if err != nil {
		return fmt.Errorf("CheckImage os.Getwd err:%v", err)
	}

	//不存在目录就创建
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("CheckImage IsNotExistMkDir err: %v ", err)
	}

	//是否有权限
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
