package user

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/gin-example/interface/http/user/base"
	"github.com/hiagomf/gin-example/interface/http/user/contact"
)

func Router(r *gin.RouterGroup) {
	base.Router(r.Group("base"))
	contact.Router(r.Group("contact"))
}
