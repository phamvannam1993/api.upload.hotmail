package util

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Response ...
type Response struct {
	HTTPCode int         `json:"-"`
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
}

func sendResponse(c echo.Context, data Response) error {
	return c.JSON(data.HTTPCode, echo.Map{
		"data":    data.Data,
		"message": data.Message,
		"success": data.Success,
	})
}

func getResponse(httpCode int, data interface{}, message string) Response {
	success := true
	if httpCode != http.StatusOK {
		success = false
	}

	return Response{
		HTTPCode: httpCode,
		Success:  success,
		Data:     data,
		Message:  message,
	}
}

// Response200 success
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Thành công!"
	}

	if data == nil {
		data = echo.Map{}
	}

	resp := getResponse(http.StatusOK, data, message)
	return sendResponse(c, resp)
}

// Response400 bad request
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không hợp lệ!"
	}

	if data == nil {
		data = echo.Map{}
	}

	resp := getResponse(http.StatusBadRequest, data, message)
	return sendResponse(c, resp)
}

// Response401 unauthorized
func Response401(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Bạn không có quyền thực hiện hành động này!"
	}

	if data == nil {
		data = echo.Map{}
	}

	resp := getResponse(http.StatusUnauthorized, data, message)
	return sendResponse(c, resp)
}

// Response404 not found
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Dữ liệu không tìm thấy!"
	}

	if data == nil {
		data = echo.Map{}
	}

	resp := getResponse(http.StatusNotFound, data, message)
	return sendResponse(c, resp)
}

// ResponseRouteValidation ...
// WARNING: only match with lib "github.com/go-ozzo/ozzo-validations/v4"
func ResponseRouteValidation(c echo.Context, err error) error {
	errors := strings.Split(err.Error(), ";")
	totalErrors := len(errors)

	// If have no errors
	if totalErrors == 0 {
		return Response400(c, nil, "")
	}

	// Get message
	message := strings.Split(errors[0], ":")[1]
	// Trim unnecessary spaces
	message = strings.Trim(message, " ")

	// Return
	return Response400(c, nil, message)
}
