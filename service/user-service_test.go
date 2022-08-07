package service

import (
	"fmt"
	"testing"

	"github.com/PutraFajarF/go-project-api/config"
	"github.com/PutraFajarF/go-project-api/dto"
	"github.com/PutraFajarF/go-project-api/repository"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestUserService_Update(t *testing.T) {
	dbPost := config.SetupDatabaseConnection()
	config.CloseDatabaseConnection(dbPost)
	repositoryUser := repository.NewUserRepository(dbPost)
	services := NewUserService(repositoryUser)
	users := dto.UserUpdateDTO{
		ID:       1,
		Name:     "Febri",
		Email:    "febri@gmail.com",
		Password: "Febriss",
	}
	userServices := services.Update(users)
	fmt.Println(userServices)
}
