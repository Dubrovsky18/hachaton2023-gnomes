package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetProfiles(c *gin.Context) {
	jsonData, err := json.Marshal(HiddenUsers)
	if err != nil {
		log.Fatalf(err.Error())
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
func (ctrl *Controller) ChangeUsers(c *gin.Context) {
	var input models.Hidden
	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Role == "student" {
		err := ctrl.services.Student.Create(models.Student{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	} else if input.Role == "teacher" {
		err := ctrl.services.Teacher.Create(models.Teacher{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	} else if input.Role == "admin" {
		err := ctrl.services.Admin.Create(models.Admin{
			User: input.User,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	} else {
		c.String(http.StatusBadRequest, "Invalid Role")
		return
	}

	delete(HiddenUsers, input.User.Email)

	c.Status(http.StatusOK)
}
