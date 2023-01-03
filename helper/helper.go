package helper

import (
	"context"
	config "echo-cloudinary-api/configs"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input interface{}) (string, error) {
	//*Defined a timeout of 10 seconds when connecting to Cloudinary.
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    //* create cloudinary instance
    cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
    if err != nil {
        return "", err
    }

    //* upload file
    uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
    if err != nil {
        return "", err
    }
    fmt.Println(uploadParam)
    return uploadParam.SecureURL, nil
}