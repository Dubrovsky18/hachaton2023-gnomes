package auth

import (
	"encoding/json"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func (ctrl *Controller) GetProfiles(c *fiber.Ctx) error {
	jsonData, err := json.Marshal(HiddenUsers)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return c.Send(jsonData)
}

func (ctrl *Controller) ChangeUsers(c *fiber.Ctx) error {

	var input models.Hidden
	if err := c.BodyParser(&input); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	if input.Role == "student" {
		err := ctrl.services.Student.Create(models.Student{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
	} else if input.Role == "teacher" {
		err := ctrl.services.Teacher.Create(models.Teacher{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
	} else if input.Role == "admin" {
		err := ctrl.services.Admin.Create(models.Admin{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
	} else {
		err := c.Send([]byte("Invalid Role"))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return err
		}
	}

	delete(HiddenUsers, input.User.Email)

	c.Status(http.StatusOK)

	return nil
}
