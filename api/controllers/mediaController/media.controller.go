package mediaController

import (
	"net/http"

	"api.com/api/dtos"
	"api.com/api/models"
	"api.com/api/services/mediaService"
	"github.com/gin-gonic/gin"
)

func FileUpload() gin.HandlerFunc {
    return func(c *gin.Context) {
        //upload
        formFile, _, err := c.Request.FormFile("file")
        if err != nil {
            c.JSON(
                http.StatusInternalServerError,
                dtos.MediaDto{
                    StatusCode: http.StatusInternalServerError,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": "Select a file to upload"},
                })
            return
        }

        uploadUrl, err := mediaService.NewMediaUpload().FileUpload(models.File{File: formFile})
        if err != nil {
            c.JSON(
                http.StatusInternalServerError,
                dtos.MediaDto{
                    StatusCode: http.StatusInternalServerError,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": "Error uploading file"},
                })
            return
        }

        c.JSON(
            http.StatusOK,
            dtos.MediaDto{
                StatusCode: http.StatusOK,
                Message:    "success",
                Data:       map[string]interface{}{"data": uploadUrl},
            })
    }
}

func RemoteUpload() gin.HandlerFunc {
    return func(c *gin.Context) {
        var url models.Url

        //validate the request body
        if err := c.BindJSON(&url); err != nil {
            c.JSON(
                http.StatusBadRequest,
                dtos.MediaDto{
                    StatusCode: http.StatusBadRequest,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": err.Error()},
                })
            return
        }

        uploadUrl, err := mediaService.NewMediaUpload().RemoteUpload(url)
        if err != nil {
            c.JSON(
                http.StatusInternalServerError,
                dtos.MediaDto{
                    StatusCode: http.StatusInternalServerError,
                    Message:    "error",
                    Data:       map[string]interface{}{"data": "Error uploading file"},
                })
            return
        }

        c.JSON(
            http.StatusOK,
            dtos.MediaDto{
                StatusCode: http.StatusOK,
                Message:    "success",
                Data:       map[string]interface{}{"data": uploadUrl},
            })
    }
}