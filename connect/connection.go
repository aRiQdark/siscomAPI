package connect

import (
	"gin-gonic-gorm/configuration"
	"gin-gonic-gorm/handler"
	"gin-gonic-gorm/route"

	"github.com/gin-gonic/gin"
)

func Connection() {
    app := gin.Default()
    
    handlers := &handler.Handlers{
        MemberHandler:  &handler.MemberHandler{},
        HistoryHandler: &handler.Historyhandler{},
    }
    
    route.Route(app, handlers)
    app.Run(configuration.Port)
}
