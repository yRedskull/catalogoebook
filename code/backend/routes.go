package main

import (
	"code/handlers"
	"code/token_app"
	/* "net/http" */
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	mu sync.Mutex
)

func Routes(r *gin.Engine) {

	mu.Lock()
	token_app.JWT_KEY = []byte(os.Getenv("JWT_KEY"))
	mu.Unlock()

	r.StaticFile("/robots.txt", "../frontend/static/robots.txt")

	/* r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/entrar")
	})

	r.GET("/entrar", handlers.PageLogin)
	r.POST("/login", handlers.VerifyLogin)

	r.GET("/criar-conta", handlers.PageCreateAccount)
	r.POST("/create-account", handlers.CreateAccount) */

	r.GET("/bnf/:username", handlers.ViewEbook)
	/* r.POST("/bonifica/:username", handlers.Vi) */
	
	
	/* r.POST("/wts/:username/create-lead", handlers.CreateLeadHandler) */
	/* r.PUT("/wts/:username/edit", handlers.EditAttendantsHandler)
	r.DELETE("/wts/:username/delete", handlers.DeleteAttendantHandler) */

	/* r.GET("/dashboard", token_app.TokenAuthMiddleware(), handlers.PageDashboardReport)
	r.GET("/dashboard/gerenciador", token_app.TokenAuthMiddleware(), handlers.PageDashboardManager)
	r.GET("/dashboard/mensagens", token_app.TokenAuthMiddleware(), handlers.PageDashboardMessage)
	r.GET("/dashboard/leads", token_app.TokenAuthMiddleware(), handlers.PageDashboardLead)
	r.GET("/dashboard/pixel", token_app.TokenAuthMiddleware(), handlers.PageDashboardPixel)

	r.GET("/logout", handlers.Logout)

	r.GET("/dashboard/attendants/read", token_app.TokenAuthMiddleware(), handlers.ReadAttendantsHandler)
	r.POST("/dashboard/attendants/create", token_app.TokenAuthMiddleware(), handlers.CreateAttendantsHandler)
	r.PUT("/dashboard/attendants/edit", token_app.TokenAuthMiddleware(), handlers.EditAttendantsHandler)
	r.PUT("/dashboard/attendants/toggle-status", token_app.TokenAuthMiddleware(), handlers.ToggleStatusAttendantHandler)
	r.DELETE("/dashboard/attendants/delete", token_app.TokenAuthMiddleware(), handlers.DeleteAttendantHandler)

	r.GET("/dashboard/messages/read", token_app.TokenAuthMiddleware(), handlers.ReadMessageHandler)
	r.POST("/dashboard/messages/create", token_app.TokenAuthMiddleware(), handlers.CreateMessageHandler)
	r.PUT("/dashboard/messages/edit", token_app.TokenAuthMiddleware(), handlers.EditMessageHandler)
	r.DELETE("/dashboard/messages/delete", token_app.TokenAuthMiddleware(), handlers.DeleteMessageHandler)

	r.GET("/dashboard/leads/read", token_app.TokenAuthMiddleware(), handlers.ReadLeadHandler)

	r.GET("/dashboard/pixel/read", token_app.TokenAuthMiddleware(), handlers.ReadPixelHandler)
	r.PUT("/dashboard/pixel/set", token_app.TokenAuthMiddleware(), handlers.SetPixelHandler)

	r.GET("/dashboard/historic/redirects", token_app.TokenAuthMiddleware(), ReadLogRedirectsHandler)

	r.GET("/dashboard/ws", token_app.TokenAuthMiddleware(), handlers.SocketDashboardHandler)

	r.GET("/dashboard/username/read", token_app.TokenAuthMiddleware(), handlers.ReadUsernameHandler)
	r.PUT("/dashboard/username/edit", token_app.TokenAuthMiddleware(), handlers.EditUsernameHandler)
 */
}
