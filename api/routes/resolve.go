// once we use shorten.go file to shorten the actual URL then the shortened URL needs to redirect to the actual link, so we are saving the actual URL in our database and we will use the shortened URL

package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

}
