package auth

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (ctrl *Controller) loginOAuth2(c *fiber.Ctx) error {
	role := c.Params("role")
	if role == "student" {
		var input models.Student
		if err := c.BodyParser(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return nil
		}

	}

	var ok bool
	input.Phone, ok = helpers.IsValidRussianPhoneNumber(input.Phone)
	if !ok {
		pkg.NewErrorResponse(c, http.StatusBadRequest, "number not in Russian format")
		return nil
	}

	_, err := ctrl.studentService.Client.GetUserFromMobile(input.Phone)
	if err == nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, "user id already in system")
		return nil
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
