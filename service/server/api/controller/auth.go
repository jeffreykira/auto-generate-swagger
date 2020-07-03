package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jeffreykira/log-management/service/server/api/middleware"
	"github.com/jeffreykira/log-management/service/server/api/model"
	"github.com/jeffreykira/log-management/service/server/api/response"
)

// SigninValid godoc
// @Summary sign in
// @Description Sign in to get jwt token
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body model.Signin true "signin information"
// @Success 200 {object} response.SigninSuccess
// @Failure 400 {object} response.HTTPError
// @Failure 401 {object} response.HTTPError
// @Router /auth/signin [post]
func (c *Controller) SigninValid(ctx *gin.Context) {
	var reqBody model.Signin
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, response.HTTPError{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	if reqBody.Username != "jeffrey" || reqBody.Password != "kira" {
		ctx.JSON(http.StatusUnauthorized, response.HTTPError{
			Code:    401,
			Message: "validation error",
		})
		return
	}

	now := time.Now()
	jwtID := reqBody.Username + strconv.FormatInt(now.Unix(), 10)

	claims := model.Claims{
		Role: reqBody.Username,
		StandardClaims: jwt.StandardClaims{
			Audience:  reqBody.Username,
			ExpiresAt: now.Add(60 * time.Minute).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    "CMS-Agent",
			NotBefore: now.Unix(),
			Subject:   reqBody.Username,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(middleware.JwtSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.HTTPError{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.SigninSuccess{
		Token: token,
	})
}

//TestHandler ...
func TestHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	ctx.JSON(http.StatusOK, gin.H{
		"role": role,
	})
}
