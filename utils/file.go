package utils

import (
	"ReMotion-C7/config"
	"fmt"
	"io"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFileToStorage(c *fiber.Ctx, name string) (string, error) {

	srcFile, err := c.FormFile(name)
	if err != nil {
		return "", err
	}

	file, err := srcFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	if seeker, ok := file.(io.Seeker); ok {
		seeker.Seek(0, io.SeekStart)
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), srcFile.Filename)
	fmt.Println(fileName)

	bucket := config.LoadEnvConfig("BUCKET")

	_, err = config.GetStorageClient().UploadFile(bucket, fileName, file)
	if err != nil {
		return "", err
	}

	publicUrl := config.GetStorageClient().GetPublicUrl(bucket, fileName)

	return publicUrl.SignedURL, nil

}
