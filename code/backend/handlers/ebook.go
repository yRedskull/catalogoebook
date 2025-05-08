package handlers

import (
	"code/database"
	/* "code/engine" */
	"code/pixel"
	"code/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)


func ViewEbook(c *gin.Context) {
	username, _ := c.Params.Get("username")

	manager, err := database.FindByOneManager("username", username)

	if err != nil || manager == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Erro na Requisição": "Entre em contato com o anunciante!",
		})
		return
	}

	id := manager["id"].(string)
	contact := manager["contact"].(string)
	first_name := utils.CapitalizeSentence(utils.GetFirstName(manager["name"].(string)))

	link_wts := utils.GenerateLinkWts(contact, "Olá, vim pelo ebook da Bonífica e gostaria de saber mais sobre os produtos.")

	data := map[string]any{
		"first_name": first_name,
		"link_wts": link_wts,
		"source_pdf": "https://filess3.s3.sa-east-1.amazonaws.com/Cat%C3%A1logo+bnf+2025.pdf",
		"thumb_pdf": "https://filess3.s3.sa-east-1.amazonaws.com/Cat%C3%A1logo+bnf+2025_page-0001.jpg",

	}

	pixel_manager, exists_pixel := manager[id].(bson.M)["pixel"]

	if exists_pixel {
		pixel_id, exists_pixel_id := pixel_manager.(bson.M)["id"]
		access_token, exists_access_token := pixel_manager.(bson.M)["access_token"]

		if exists_pixel_id && exists_access_token {
			fbc := pixel.ExtractFbc(c)
			test_event_code, _ := c.GetQuery("test_event_code")
			request_info := utils.CaptureRequestInfo(c)

			go pixel.PageView(pixel_id.(string), access_token.(string), fbc, test_event_code, request_info)

		}
	}

	c.HTML(http.StatusOK, "ebook.html", data)
}
