package initializers

import (
	dependencies "github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/dependencies"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	Describe("InitializeRouter()", func() {
		var (
			c *dependencies.Container
		)

		BeforeEach(func() {
			c = &dependencies.Container{}
		})

		It("should initialize router", func() {
			r := InitializeRouter(c)

			Expect(r).NotTo(BeNil())
		})
	})
})
