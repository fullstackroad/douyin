package controller

import (
	"net/http"

	"github.com/RaymondCode/simple-demo/common/status_code"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	status_code.Status
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, status_code.Status{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, status_code.Status{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Status: status_code.Status{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Status: status_code.Status{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
