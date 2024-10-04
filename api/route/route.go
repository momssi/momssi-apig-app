package route

import (
	"github.com/gin-gonic/gin"
	"momssi-apig-app/api/controller"
)

type RouterConfig struct {
	Engine           *gin.Engine
	MemberController *controller.MemberController
}

func (r *RouterConfig) Setup() {
	r.SetupMember()
}

func (r *RouterConfig) SetupMember() {
	member := r.Engine.Group("/member")
	member.POST("/sign-up", r.MemberController.SignUp)
	member.POST("/login", r.MemberController.Login)
}
