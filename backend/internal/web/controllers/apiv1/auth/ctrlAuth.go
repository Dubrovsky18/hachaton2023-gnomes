package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var HiddenUsers map[string]string

func (ctrl *Controller) register(c *fiber.Ctx) error {

	var hidden models.Hidden

	if err := c.BodyParser(&hidden); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	_, ok := ctrl.alreadyEmail(hidden.User.Email)
	if ok {
		pkg.NewErrorResponse(c, http.StatusBadRequest, "user id already in system")
		return nil
	} else {
		HiddenUsers = map[string]string{
			hidden.User.Email: hidden.User.Name,
		}

		session := &fiber.Cookie{
			Name:   hidden.User.Name,
			Value:  "hidden",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		c.Cookie(session)
		return nil
	}
}

func (ctrl *Controller) loginAuth(c *fiber.Ctx) error {

	role := c.Params("role")
	if role == "student" {
		var input models.Student
		if err := c.BodyParser(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		student, err := ctrl.services.Student.GetLogin(input.User.Email)
		if (err != nil) || (student.User.Password != input.User.Password) {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}

		session := &fiber.Cookie{
			Name:   input.User.Name,
			Value:  "student",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		c.Cookie(session)

	} else if role == "teacher" {
		var input models.Teacher
		if err := c.BodyParser(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return nil
		}
		teacher, err := ctrl.services.Teacher.GetLogin(input.User.Email)
		if (err != nil) || (teacher.User.Password != input.User.Password) {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}

		session := &fiber.Cookie{
			Name:   input.User.Name,
			Value:  "teacher",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		c.Cookie(session)
	}

	return nil
}
