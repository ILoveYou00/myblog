package v1

import (
	"github.com/ILoveYou00/myblog/pkg/upload"
	"github.com/ILoveYou00/myblog/util/app"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		app.ResponseError(c, app.ERROR)
	}
	if image == nil {
		app.ResponseError(c, app.INVALID_PARAMS)
	} else {

		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || upload.CheckImageSize(file) {
			app.ResponseError(c, app.ERROR_UPLOAD_CHECK_IMAGE_FORMAT)
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				app.ResponseError(c, app.ERROR_AUTH_CHECK_TOKEN_FAIL)
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				app.ResponseError(c, app.ERROR_AUTH_CHECK_TOKEN_FAIL)
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
				app.ResponseSuccess(c, data)
			}
		}
	}
}
