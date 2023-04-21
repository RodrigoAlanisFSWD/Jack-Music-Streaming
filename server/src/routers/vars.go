package routers

import (
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
	"github.com/Jack-Music-Streaming/server/src/repository"
	"github.com/Jack-Music-Streaming/server/src/services"
)

var userRepository models.UserRepository
var userService models.UserService
var profileRepository models.ProfileRepository
var profileService models.ProfileService
var filesRepository models.FilesRepository
var songRepository models.SongRepository
var songService models.SongService
var jwtRepository models.JWTRepository
var authService models.AuthService
var playlistRepository models.PlaylistRepository
var playlistService models.PlaylistService

func InitializeRotuers() {
	userRepository = repository.NewUserRepository(database.DB)
	userService = services.NewUserService(userRepository)
	profileRepository = repository.NewProfileRepository(database.DB)
	filesRepository = repository.NewFilesRepository(database.DB)
	songRepository = repository.NewSongRepository(database.DB)
	jwtRepository = repository.NewJWTRepository(database.DB)
	profileService = services.NewProfileService(profileRepository, filesRepository)
	authService = services.NewAuthService(jwtRepository, userService, profileService)
	songService = services.NewSongService(songRepository, filesRepository)
	playlistRepository = repository.NewPlaylistRepository(database.DB)
	playlistService = services.NewPlaylistService(playlistRepository, songService)
}
