package upload

import (
	"fmt"
	"github.com/ILoveYou00/myblog/config"
	"github.com/ILoveYou00/myblog/util"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//GetImageFullUrl 获取图片完整访问 URLGetImageFullUrl：获取图片完整访问 URL
func GetImageFullUrl(name string) string {
	return config.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

//GetImageName 获取图片名称
func GetImageName(name string) string {
	//获取文件扩展名
	ext := path.Ext(name)
	//根据扩展名进行分割
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//GetImagePath 获取图片路径
func GetImagePath() string {
	return config.AppSetting.ImageSavePath
}

//GetImageFullPath 获取图片完整路径
func GetImageFullPath() string {
	return config.AppSetting.RuntimeRootPath + GetImagePath()
}

//CheckImageExt 检查图片后缀
func CheckImageExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range config.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

//CheckImageSize 检查图片大小
func CheckImageSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	size := len(content)
	if err != nil {
		log.Println(err)
		return false
	}

	return size <= config.AppSetting.ImageMaxSize
}

//CheckImage 检查图片
func CheckImage(src string) error {
	// Getwd 返回对应的根路径名
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	//判断是否存在该文件
	err = IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	//检查文件权限
	_, err = os.Stat(src)
	if os.IsPermission(err) == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}

//IsNotExistMkDir 判断文件是否存在，如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	_, err := os.Stat(src)
	if os.IsNotExist(err) == true {
		//创建文件
		err := os.MkdirAll(src, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
