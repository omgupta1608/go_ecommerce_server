package utils

import (
	"errors"
	"strconv"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/omgupta1608/aftershoot_task/types"
)

func PrintToConsole(msg string, reportType string) {
	switch reportType {
	case "error":
		color.Red(msg)
	case "info":
		color.Blue(msg)
	case "log":
		color.Cyan(msg)
	case "success":
		color.Green(msg)
	}
}

func GetVersion() string {
	return "v1"
}

func Find(arr []string, val string) int {
	for i, s := range arr {
		if s == val {
			return i
		}
	}
	return -1
}

func AreEqualArray(a []string, b []string) bool {
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func GetPaginationOpts(size string, page string) (int, int, error) {
	limit, err := strconv.Atoi(size)

	if err != nil {
		return ONE_MILLION, 0, err
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return ONE_MILLION, 0, err
	}

	return limit, (offset - 1) * limit, nil
}

func StringToUUID(val string) (res uuid.UUID) {
	res, err := uuid.Parse(val)

	if err != nil {
		return
	}
	return
}

func GetUserFromContext(c *gin.Context) (*types.CustomClaims, error) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		return nil, errors.New("Bad Token")
	}

	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		return nil, err
	}

	user_email := c.GetString("user_email")
	if user_email == "" {
		return nil, errors.New("Bad Token")
	}

	is_user_active := c.GetBool("is_user_active")
	if !is_user_active {
		return nil, errors.New("Bad Token")
	}

	return &types.CustomClaims{
		UserId:   user_uuid,
		IsActive: is_user_active,
		Email:    user_email,
	}, nil
}
