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

	} else if role == "teacher" {
		var input models.Teacher
		if err := c.BodyParser(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return nil
		}
		teacher, err :=
	}

	uuid, err := ctrl.studentService.Client.CreateUser(input)

	if err != nil {
		pkg.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	pkg.NewJsonInterfaceResponse(c, http.StatusOK, "", "Auth-srv", map[string]interface{}{
		"id": uuid,
	})

	return nil
}
