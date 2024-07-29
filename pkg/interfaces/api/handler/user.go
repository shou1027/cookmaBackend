package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shou1027/cookmaBackend/pkg/myerror"
	"github.com/shou1027/cookmaBackend/pkg/usecase"
)

type Handler interface {
	HandleSignup(c *gin.Context)
	HandleLogin(c *gin.Context)
	HandleLogout(c *gin.Context)
}

type handler struct {
	useCase usecase.UseCase
}

func NewHandler(userUseCase usecase.UseCase) Handler {
	return &handler{
		useCase: userUseCase,
	}
}

func (h *handler) HandleSignup(c *gin.Context) {
	type (
		request struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8`
		}
		response struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}
	)

	requestBody := new(request)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.useCase.Signup(c.Request.Context(), requestBody.Username, requestBody.Email, requestBody.Password)
	if err != nil {
		switch e := err.(type) {
		case *myerror.InternalServerError:
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Err.Error()})
			return
		case *myerror.BadRequestError:
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, &response{
		ID:       user.GetId(),
		Username: user.GetName(),
		Email:    user.GetEmail(),
	})
}

func (h *handler) HandleLogin(c *gin.Context) {
	type (
		request struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}
		response struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		}
	)

	requestBody := new(request)

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	signedString, user, err := h.useCase.Login(c.Request.Context(), requestBody.Email, requestBody.Password)

	if err != nil {
		switch e := err.(type) {
		case *myerror.InternalServerError:
			c.JSON(http.StatusInternalServerError, gin.H{"error": e.Err.Error()})
			return
		case *myerror.BadRequestError:
			c.JSON(http.StatusBadRequest, gin.H{"error": e.Err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.SetCookie("jwt", signedString, 60*60*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &response{
		ID:       user.GetId(),
		Username: user.GetName(),
	})
}

func (h *handler) HandleLogout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successfuc"})
}
