package utils

import (
	"bytes"
	"ecommerce_site/src/core/entities"
	"fmt"

	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ProcessImages(fileHeaders []*multipart.FileHeader) ([]entities.Data, error) {
	var result []entities.Data

	url := "https://api.imgbb.com/1/upload?key=35ab53548bf387fc865eded79efa0a19"
	method := "POST"

	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)

		part, err := writer.CreateFormFile("image", filepath.Base(fileHeader.Filename))
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}

		err = writer.Close()
		if err != nil {
			return nil, err
		}

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var responseData map[string]interface{}
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			return nil, err
		}

		data := entities.Data{
			//	ID:  responseData["data"].(map[string]interface{})["id"].(string),
			URL: responseData["data"].(map[string]interface{})["url"].(string),
		}

		result = append(result, data)
	}

	return result, nil
}

func ProcessImageSign(fileHeader *multipart.FileHeader) (entities.Data, error) {
	var result entities.Data

	url := "https://api.imgbb.com/1/upload?key=35ab53548bf387fc865eded79efa0a19"
	method := "POST"

	file, err := fileHeader.Open()
	if err != nil {
		return result, err
	}
	defer file.Close()

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	part, err := writer.CreateFormFile("image", filepath.Base(fileHeader.Filename))
	if err != nil {
		return result, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return result, err
	}

	err = writer.Close()
	if err != nil {
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return result, err
	}

	//result.ID = responseData["data"].(map[string]interface{})["id"].(string)
	result.URL = responseData["data"].(map[string]interface{})["url"].(string)

	return result, nil
}
func GetUploadedFiles(c *gin.Context) ([]*multipart.FileHeader, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files, ok := form.File["image"]
	if !ok || len(files) == 0 {
		return nil, nil
	}

	var uploadedFiles []*multipart.FileHeader
	for _, file := range files {
		if file.Size == 0 {
			return nil, fmt.Errorf("Uploaded file is empty")
		}
		uploadedFiles = append(uploadedFiles, file)
	}

	return uploadedFiles, nil
}
