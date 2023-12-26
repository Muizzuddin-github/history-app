package utility

import (
	"bytes"
	"crud/src/responsebody"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func ReadByte (file *multipart.FileHeader) ([]byte, error){
	fileOpen, err := file.Open()
	if err != nil{
		return nil,errors.New(err.Error())
	}

	imageByte, err := io.ReadAll(fileOpen)
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return imageByte,nil
}

func UploadImageApi(source []byte,filename string) (responsebody.ImageHostAPI, error) {

	image := new(bytes.Buffer)
	writer := multipart.NewWriter(image)
	file, err := writer.CreateFormFile("source",filename)
	if err != nil{
		return responsebody.ImageHostAPI{}, errors.New(err.Error())
	}

	file.Write(source)
	writer.Close()

	req,err := http.NewRequest(http.MethodPost,os.Getenv("FREE_IMAGE_HOST_URL"),image)
	if err != nil{
		return responsebody.ImageHostAPI{}, errors.New(err.Error())
	}

	req.Header.Set("Content-Type",writer.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil{
		return responsebody.ImageHostAPI{}, errors.New(err.Error())
	}
	body := res.Body
	defer body.Close()

	dataByte, err := io.ReadAll(body)
	if err != nil{
		return responsebody.ImageHostAPI{}, nil
	}


	dataJson := responsebody.ImageHostAPI{}

	err = json.Unmarshal(dataByte, &dataJson)
	if err != nil{
		return responsebody.ImageHostAPI{}, errors.New(err.Error())
	}

	return dataJson, nil
}