package auth

import (
	"net/http"

	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg"
	"github.com/gin-gonic/gin"
)

var HiddenUsers map[string]string

func (ctrl *Controller) register(c *gin.Context) {
	var hidden models.Hidden

	if err := c.ShouldBindJSON(&hidden); err != nil {
		pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, ok := ctrl.alreadyEmail(hidden.User.Email)
	if ok {
		pkg.NewErrorResponse(c, http.StatusBadRequest, "user id already in system")
		return
	} else {
		HiddenUsers = map[string]string{
			hidden.User.Email: hidden.User.Name,
		}

		session := &http.Cookie{
			Name:   hidden.User.Name,
			Value:  "hidden",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		http.SetCookie(c.Writer, session)
	}
}
func (ctrl *Controller) loginAuth(c *gin.Context) {
	role := c.Param("role")
	if role == "student" {
		var input models.Student
		if err := c.ShouldBindJSON(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		student, err := ctrl.services.Student.GetLogin(input.User.Email)
		if err != nil || student.User.Password != input.User.Password {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		session := &http.Cookie{
			Name:   input.User.Name,
			Value:  "student",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		http.SetCookie(c.Writer, session)

	} else if role == "teacher" {
		var input models.Teacher
		if err := c.ShouldBindJSON(&input); err != nil {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		teacher, err := ctrl.services.Teacher.GetLogin(input.User.Email)
		if err != nil || teacher.User.Password != input.User.Password {
			pkg.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		session := &http.Cookie{
			Name:   input.User.Name,
			Value:  "teacher",
			MaxAge: 24 * 60 * 60,
			Path:   "/",
		}
		http.SetCookie(c.Writer, session)
	}
}
