package v1

import (
	"example/01_qrcode/pkg/app"
	"example/01_qrcode/pkg/e"
	"example/01_qrcode/pkg/qrcode"
	"example/01_qrcode/service"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	QRCODE_URL = "https://www.baidu.com/"
)

func GenerateArticlePoster(c *gin.Context) {
	appG := app.Gin{C: c}
	article := &service.Article{}

	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	posterName := service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	articlePoster := service.NewArticlePoster(posterName, article, qr)
	articlePosterBgService := service.NewArticlePosterBg(
		"bg.jpg",
		articlePoster,
		&service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&service.Pt{
			X: 125,
			Y: 298,
		},
	)

	_, filePath, err := articlePosterBgService.Generate()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"poster_url":      qrcode.GetQrCodeFullUrl(posterName),
		"poster_save_url": filePath + posterName,
	})

}
