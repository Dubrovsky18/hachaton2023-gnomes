package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) loginAuth(c *fiber.Ctx) error {

	role := c.Params("role")
	if role == "student" {
		var input models.Student
		if err := c.BodyParser(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return err
		}
		student, err := ctrl.studentService.GetLogin(input.User.Email)
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
		teacher, err := ctrl.teacherService.GetLogin(input.User.Email)
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

	pkg.NewJsonInterfaceResponse(c, http.StatusOK, "", "Auth-srv", map[string]interface{}{
		"id": uuid,
	})

	return nil
}
