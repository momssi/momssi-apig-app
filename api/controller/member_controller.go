package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		if errors.Is(err, form.GetCustomErr(form.ErrDuplicatedUsername)) {
			mc.failResponse(c, http.StatusBadRequest, form.ErrDuplicatedUsername, fmt.Errorf("duplicated username : %w", err))
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &member.Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(member.JWTKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		mc.failResponse(c, http.StatusInternalServerError, form.ErrFailGenerateJWTKey, form.GetCustomErr(form.ErrFailGenerateJWTKey))
		return
	}

	mc.successResponse(c, http.StatusOK, form.ApiResponse{
		ErrorCode: 0,
		Message:   "success",
		Result: member.LoginRes{
			AccessToken: tokenString,
		},
	})

}
