package v1public

import (
	"net/http"
	model "seifwu/app/models"
	param "seifwu/app/params"
	service "seifwu/app/services"
	"seifwu/global/response"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

var signUpParam param.SignUpParam

// SignUp 注册
func SignUp(c *gin.Context) {

	// 参数错误处理
	if err := c.ShouldBind(&signUpParam); err != nil {
		errResult := util.UnifiedValidation(c, err, "40001", signUpParam)

		response.Response(c, http.StatusBadRequest, "40001", nil, errResult, nil)
		return
	}

	user := model.User{
		UserName: signUpParam.Username,
		NickName: signUpParam.Username,
		Email:    signUpParam.Email,
		Password: signUpParam.Password,
	}

	newUser, err := service.SignUpService(&user)
	if err != nil {
		response.Response(c, http.StatusBadRequest, "40001", nil, err, nil)
		return
	}

	response.Response(c, http.StatusOK, "0", gin.H{"user": newUser}, "创建成功", nil)
}
