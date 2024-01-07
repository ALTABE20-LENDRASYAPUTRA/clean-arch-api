package handler

import (
	"clean-arch-api/app/middlewares"
	"clean-arch-api/features/user"
	"clean-arch-api/utils/responses"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) GetAllUsers(c echo.Context) error {
	// panggil func di service layer
	results, errSelect := handler.userService.GetAll()
	if errSelect != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]any{
		// 	"message": "error read data. " + errSelect.Error(),
		// })
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}
	// proses mapping dari core ke response
	usersResult := CoreToResponseList(results)
	// var usersResult []UserResponse
	// for _, value := range results {
	// 	usersResult = append(usersResult, UserResponse{
	// 		ID:    value.ID,
	// 		Name:  value.Name,
	// 		Email: value.Email,
	// 	})
	// }

	// return c.JSON(http.StatusOK, map[string]any{
	// 	"message": "success",
	// 	"data":    usersResult,
	// })
	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", usersResult))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		// return c.JSON(http.StatusBadRequest, map[string]any{
		// 	"message": "error bind data. data not valid",
		// })
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(newUser); err != nil {
		c.Echo().Logger.Error("Input error :", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
			"data":    nil,
		})
	}

	//mapping dari request ke core
	userCore := RequestToCore(newUser)
	// userCore := user.Core{
	// 	Name:        newUser.Name,
	// 	Email:       newUser.Email,
	// 	Password:    newUser.Password,
	// 	Address:     newUser.Address,
	// 	PhoneNumber: newUser.PhoneNumber,
	// 	Role:        newUser.Role,
	// }

	errInsert := handler.userService.Create(userCore)
	if errInsert != nil {
		c.Logger().Error("ERROR Register, explain:", errInsert.Error())
		if strings.Contains(errInsert.Error(), "Duplicate entry") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "data yang diinputkan sudah terdaftar pada sistem",
			})
		}
		
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "terjadi permasalahan ketika memproses data",
		})
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	var userData = UserRequest{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	userCore := RequestToCore(userData)
	_, err := handler.userService.Update(userID, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var reqData = LoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	result, token, err := handler.userService.Login(reqData.Email, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error login "+err.Error(), nil))
	}
	responseData := map[string]any{
		"token": token,
		"nama":  result.Name,
	}
	return c.JSON(http.StatusOK, responses.WebResponse("success login", responseData))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized user",
		})
	}

	err := handler.userService.Delete(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error deleting user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User deleted successfully",
	})
}