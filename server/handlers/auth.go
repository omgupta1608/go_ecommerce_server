package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/omgupta1608/aftershoot_task/db"
	"github.com/omgupta1608/aftershoot_task/types"
	"github.com/omgupta1608/aftershoot_task/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	var body types.LoginBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	user, err := db.Conn.GetUserByEmail(c, body.Email)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	cc := types.CustomClaims{
		UserId:   user.ID.String(),
		IsActive: true,
		Email:    user.Email,
	}

	token, err := utils.GenerateJWTToken(cc)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Logged In", map[string]any{
		"access_token": token,
		"user_id":      user.ID.String(),
		"email":        user.Email,
	})
}

func RegisterHandler(c *gin.Context) {
	var body types.RegisterBody
	if err := c.BindJSON(&body); err != nil {
		utils.SendError(c, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}
	user, err := db.Conn.CreateUser(c, db.CreateUserParams{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	})

	cc := types.CustomClaims{
		UserId:   user.ID.String(),
		IsActive: true,
		Email:    user.Email,
	}

	token, err := utils.GenerateJWTToken(cc)

	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(c, "Registered", map[string]any{
		"access_token": token,
		"user_id":      user.ID.String(),
		"email":        user.Email,
	})
}