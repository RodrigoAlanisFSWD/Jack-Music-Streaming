package routers

import (
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/Jack-Music-Streaming/server/src/repository"
	"github.com/Jack-Music-Streaming/server/src/services"
)

var userRepository models.UserRepository
var userService models.UserService

func InitializeRotuers() {
	userRepository = repository.NewUserRepository(database.DB)
	userService = services.NewUserService(userRepository)
}
