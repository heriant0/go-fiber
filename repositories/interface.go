package repositories

import (
	"github.com/heriant0/go-fiber/models"
	"github.com/heriant0/go-fiber/utils"
)

type UserRepository interface {
	Find(id int) (models.User, error)
	FindByUsername(username string) (models.User, error)
	FindPaginated(paginator utils.Paginator) (utils.Paginator, error)
	Save(user *models.User) error
	Delete(user *models.User) error
}

type BookRepository interface {
	Find(id int) (models.Book, error)
	Save(book *models.Book) error
	Delete(book *models.Book) error
}
