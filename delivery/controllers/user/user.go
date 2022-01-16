package user

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist-app/delivery/formatter"
	mw "todolist-app/delivery/middlewares"
	request "todolist-app/delivery/request"
	"todolist-app/delivery/response"
	"todolist-app/entities"
	"todolist-app/repository/user"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	repo user.User
}

func NewController(repo user.User) *UserController {
	return &UserController{repo}
}

func (c *UserController) GetUserByID(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	user_id, _ := strconv.Atoi(id)

	user, _ := c.repo.GetUserByID(user_id)

	if user.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "user not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	if user.ID != uint(currentUserID) {
		res.Status = http.StatusUnauthorized
		res.Message = "can't get another user data"
		res.Data = nil

		return e.JSON(http.StatusUnauthorized, res)
	}

	formatData := formatter.FormatUser(user)

	res.Status = http.StatusOK
	res.Message = "success get user"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *UserController) CreateUser(e echo.Context) error {
	res := response.Response{}

	input := request.CreateUserInput{}

	e.Bind(&input)

	e.Validate(&input)

	user := entities.User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	newUser, err := c.repo.CreateUser(user)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = newUser

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatUser(newUser)

	res.Status = http.StatusOK
	res.Message = "success create user"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *UserController) UpdateUser(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	user := entities.User{}

	user_id, _ := strconv.Atoi(id)

	user, _ = c.repo.GetUserByID(user_id)

	if user.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "user not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	input := request.UpdateUserInput{}

	e.Bind(&input)

	e.Validate(&input)

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	if user.ID != uint(currentUserID) {
		res.Status = http.StatusUnauthorized
		res.Message = "can't update another user data"
		res.Data = nil

		return e.JSON(http.StatusUnauthorized, res)
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	updatedUser, _ := c.repo.UpdateUser(user)

	formatData := formatter.FormatUser(updatedUser)

	res.Status = http.StatusOK
	res.Message = "success update user"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *UserController) DeleteUser(e echo.Context) error {
	res := response.Response{}

	id := e.Param("id")

	user := entities.User{}

	user_id, _ := strconv.Atoi(id)

	user, _ = c.repo.GetUserByID(user_id)

	if user.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "user not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	currentUserID := mw.NewAuth().ExtractTokenUserID(e)

	if user.ID != uint(currentUserID) {
		res.Status = http.StatusUnauthorized
		res.Message = "can't delete another user data"
		res.Data = nil

		return e.JSON(http.StatusUnauthorized, res)
	}

	oldUser := user

	_, err := c.repo.DeleteUser(user)

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = fmt.Sprint(err.Error())
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "no data"
		res.Data = oldUser

		return e.JSON(http.StatusBadRequest, res)
	}

	formatData := formatter.FormatUser(oldUser)

	res.Status = http.StatusOK
	res.Message = "success delete user"
	res.Data = formatData

	return e.JSON(http.StatusOK, res)
}

func (c *UserController) Login(e echo.Context) error {
	res := response.Response{}

	input := request.LoginUserFormInput{}

	e.Bind(&input)

	e.Validate(&input)

	email := input.Email
	password := input.Password

	user, _ := c.repo.GetUserByEmail(email)

	if user.ID == 0 {
		res.Status = http.StatusNotFound
		res.Message = "user not found on that email"
		res.Data = input.Email

		return e.JSON(http.StatusNotFound, res)
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		res.Status = http.StatusUnprocessableEntity
		res.Message = "Password doesn't match"
		res.Data = nil

		return e.JSON(http.StatusUnprocessableEntity, res)
	}

	token, err := mw.NewAuth().GenerateToken(int(user.ID))

	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Can't get token"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Status = http.StatusOK
	res.Message = "login success"
	res.Data = token

	return e.JSON(http.StatusOK, res)
}
