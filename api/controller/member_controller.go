package controller

import "github.com/gin-gonic/gin"

type MemberController struct {
}

func NewMemberController() *MemberController {
	return &MemberController{}
}

func (mc *MemberController) SignUp(c *gin.Context) {
}
