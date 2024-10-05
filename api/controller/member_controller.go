package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"momssi-apig-app/api/form"
	"momssi-apig-app/internal/domain/member"
	"net/http"
	"time"
)

type MemberController struct {
	service member.Service
}

func NewMemberController(service member.Service) *MemberController {
	return &MemberController{
		service: service,
	}
}

func (mc *MemberController) successResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: form.NoError,
		Message:   form.GetCustomMessage(form.NoError),
		Result:    data,
	})
}

func (mc *MemberController) failResponse(c *gin.Context, statusCode int, errorCode int, err error) {

	logMessage := form.GetCustomErrMessage(errorCode, err.Error())
	c.Errors = append(c.Errors, &gin.Error{
		Err:  fmt.Errorf(logMessage),
		Type: gin.ErrorTypePrivate,
	})

	c.JSON(statusCode, form.ApiResponse{
		ErrorCode: errorCode,
		Message:   form.GetCustomMessage(errorCode),
	})
}

func (mc *MemberController) SignUp(c *gin.Context) {

	req := member.SignUpRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.failResponse(c, http.StatusBadRequest, form.ErrParsing, fmt.Errorf("sign up json parsing err : %v", err))
		return
	}

	memberId, err := mc.service.SignUp(req)
	if err != nil {
		if errors.Is(err, form.GetCustomErr(form.ErrDuplicatedEmail)) {
			mc.failResponse(c, http.StatusBadRequest, form.ErrDuplicatedEmail, fmt.Errorf("duplicated email : %w", err))
		} else {
			mc.failResponse(c, http.StatusInternalServerError, form.ErrInternalServerError, fmt.Errorf("sign up occur err : %w", err))
		}
		return
	}

	res := member.SignUpRes{
		MemberId: memberId,
	}

	mc.successResponse(c, http.StatusOK, form.ApiResponse{
		ErrorCode: 0,
		Message:   "success",
		Result:    res,
	})

}

func (mc *MemberController) Login(c *gin.Context) {

	req := member.LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		mc.failResponse(c, http.StatusBadRequest, form.ErrParsing, fmt.Errorf("sign up json parsing err : %v", err))
		return
	}

	accessToken, err := member.GenerateJWT(req.Email, time.Minute*5)
	if err != nil {
		mc.failResponse(c, http.StatusUnauthorized, form.ErrInvalidToken, form.GetCustomErr(form.ErrMissingToken))
		return
	}

	refreshToken, err := member.GenerateJWT(req.Email, time.Hour*12)
	if err != nil {
		mc.failResponse(c, http.StatusUnauthorized, form.ErrInvalidToken, form.GetCustomErr(form.ErrMissingToken))
		return
	}

	if err := mc.service.LoginSuccess(req.Email, refreshToken); err != nil {
		mc.failResponse(c, http.StatusInternalServerError, form.ErrInternalServerError, form.GetCustomErr(form.ErrInternalServerError))
		return
	}

	mc.successResponse(c, http.StatusOK, form.ApiResponse{
		ErrorCode: 0,
		Message:   "success",
		Result: member.LoginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	})

}
