package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
	"unsafe"
)

type UserResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	RequestDate time.Time `json:"request_date"`
}
type UserRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Response struct {
	Id        uuid.UUID    `json:"response_id"`
	Status    int          `json:"status"`
	Responses UserResponse `json:"data"`
	Length    int          `json:"length"`
}

var usersResponse []UserResponse
var usersRequest []UserRequest
var responses []Response

// postUser  godoc
// @Summary postUser
// @Description post a new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body UserResponse true "User"
// @Success 200 {object} UserResponse
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /user [post]
func PostUser(ctx *gin.Context) {
	var newUserResponse UserResponse
	var newUserRequest UserRequest
	var newResponse Response

	UserId := uuid.New()
	newUserResponse.Id = UserId
	newUserResponse.RequestDate = time.Now()

	if err := ctx.BindJSON(&newUserResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Bad Request": err})
		return
	}

	newUserRequest.Name = newUserResponse.Name
	newUserRequest.Surname = newUserResponse.Surname

	RequestId := uuid.New()
	newResponse.Id = RequestId
	newResponse.Responses = newUserResponse
	statusOK := http.StatusOK
	newResponse.Status = statusOK
	length := unsafe.Sizeof(newUserResponse)
	lengthInInt := int(length)
	newResponse.Length = lengthInInt

	responses = append(responses, newResponse)
	usersResponse = append(usersResponse, newUserResponse)
	usersRequest = append(usersRequest, newUserRequest)

	ctx.IndentedJSON(http.StatusCreated, newResponse)
}

// getResponsesAndRequest godoc
// @Summary getUsers
// @Description get all responses and requests
// @Tags Responses and Requests
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Success 200 {object} UserRequest
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /users [get]
func GetUsers(ctx *gin.Context) {

	if usersResponse == nil {
		ctx.IndentedJSON(http.StatusOK, []UserResponse{})
		return
	}

	ctx.JSON(http.StatusOK, "Responses")
	ctx.IndentedJSON(http.StatusOK, responses)
	ctx.JSON(http.StatusOK, "Requests")
	ctx.IndentedJSON(http.StatusOK, usersRequest)

}
