package routers

import (
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/repository"
	"github.com/Jack-Music-Streaming/server/src/services"
)

var (
	userRepository = repository.NewUserRepository(database.DB)
	userService    = services.NewUserService(userRepository)
)
