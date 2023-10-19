package status

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/build"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controller", func() {
	var (
		statusCtrl *Controller
	)

	BeforeEach(func() {
		info := build.NewInfo()
		statusCtrl = NewController(info)
	})

	It("controller should not be nil", func() {
		Expect(statusCtrl).NotTo(BeNil())
	})

	Describe("GetStatus()", func() {
		It("should return status", func() {
			app := fiber.New()
			app.Get("/api/v1/status", func(c *fiber.Ctx) error {
				statusCtrl.GetStatus(c)
				return nil
			})

			req := httptest.NewRequest("GET", "/api/v1/status", nil)
			resp, _ := app.Test(req)

			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
