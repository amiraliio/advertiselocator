package controllers

import (
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/amiraliio/advertiselocator/helpers"
	"github.com/amiraliio/advertiselocator/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//TODO file types

func UploadMedia(request echo.Context) (err error) {
	requestedMediaType := request.Param("mediaType")
	mediaTypes := []string{"image", "video", "file", "audio"}
	sort.Strings(mediaTypes)
	indexOfMediaType := sort.SearchStrings(mediaTypes, requestedMediaType)
	if indexOfMediaType >= len(mediaTypes) || mediaTypes[indexOfMediaType] != requestedMediaType {
		return helpers.ResponseError(request, nil, http.StatusUnprocessableEntity, "CM-1000", "Media Type", "requested media type must be one the [ image, video, file, audio]")
	}
	file, err := request.FormFile("media")
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1001", "Upload File", "file not uploaded")
	}
	mediaSize, err := strconv.ParseInt(os.Getenv(strings.ToUpper(requestedMediaType)+"_SIZE"), 0, 10)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1002", "Get File Size", "cannot estimate file size")
	}
	fileSizeInMegaByte, err := helpers.ConvertByte(file.Size, "MB")
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1003", "Get File Size", "cannot estimate file size")
	}
	if int64(fileSizeInMegaByte) > mediaSize {
		return helpers.ResponseError(request, nil, http.StatusBadRequest, "CM-1004", "File Size", "media size must not be more than "+string(mediaSize))
	}
	sourceFile, err := file.Open()
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1005", "Open File", "cannot open the file")
	}
	defer sourceFile.Close()
	authData := helpers.AuthData(request)
	storagePath := helpers.Path("storage")
	filePath := "/temp/images/" + authData.UserID.Hex() + "/" + uuid.New().String()
	fileName := "/" + file.Filename
	if _, err := os.Stat(storagePath + filePath); os.IsNotExist(err) {
		if err = os.MkdirAll(storagePath+filePath, 0755); err != nil {
			return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1006", "Create Directory", "cannot create directory")
		}
	}
	destination, err := os.Create(storagePath + filePath + fileName)
	if err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1007", "Write File", "cannot write in the directory")
	}
	defer destination.Close()
	if _, err := io.Copy(destination, sourceFile); err != nil {
		return helpers.ResponseError(request, err, http.StatusBadRequest, "CM-1008", "Move File", "cannot move file to the directory")
	}
	imageModel := new(models.Image)
	imageModel.OriginalURL = filePath + fileName
	imageModel.URL = filePath + fileName
	imageModel.Size = models.OriginalSize
	return helpers.ResponseOk(request, http.StatusOK, imageModel)
}

func GetMedia(request echo.Context) (err error) {
	return nil
}
