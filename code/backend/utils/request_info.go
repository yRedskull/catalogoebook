package utils

import (
	"code/structs_utils"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

func CaptureRequestInfo(c *gin.Context) structs_utils.RequestInfo {

	ip := c.ClientIP()

	userAgent := c.GetHeader("User-Agent")
	platform := c.GetHeader("Sec-Ch-Ua-Platform")

	queryParams := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		queryParams[key] = strings.Join(values, ", ")
	}

	bodyBytes, _ := io.ReadAll(c.Request.Body)

	body := string(bodyBytes)

	// Captura o Referer (se disponível)
	referer := c.GetHeader("Referer")

	// Captura os cookies enviados pelo cliente
	cookies := make(map[string]string)
	for _, cookie := range c.Request.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}

	// Retorna a estrutura populada com os dados
	return structs_utils.RequestInfo{
		IP:        ip,
		UserAgent: userAgent,
		Platform: platform,
		Query:     queryParams,
		Body:      body,
		Referer:   referer,
		Cookies:   cookies,
	}
}