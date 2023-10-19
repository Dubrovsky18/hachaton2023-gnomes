package initializers

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/config"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/web/router"
	"github.com/gofiber/fiber/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HttpServer", func() {
	Describe("InitializeHTTPServer()", func() {
		var (
			r   *fiber.App
			cfg *config.HTTPConfig
		)

		BeforeEach(func() {
			r = router.NewRouter()
			cfg = InitializeAppConfig().HTTP
		})

		It("should initialize HTTP server", func() {
			srv := InitializeHTTPServer(r, cfg)

			Expect(srv).NotTo(BeNil())
		})
	})
})
