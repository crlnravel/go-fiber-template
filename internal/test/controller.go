package test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/crlnravel/go-fiber-template/internal/common"
)

type Controller interface {
	GetTestMsg(c *fiber.Ctx) error
}

type controller struct{}

func NewController() Controller {
	return &controller{}
}

// GetTestMsg godoc
// @Summary      Get Test Message
// @Description  Get test message.
// @Tags         Test
// @Produce      json
// @Success      200  {object} common.GeneralSuccessResponse  "OK"
// @Failure		 400  {object} common.ErrorResponse		"Bad Request"
// @Failure		 401  {object} common.ErrorResponse		"Unauthenticated"
// @Failure		 403  {object} common.ErrorResponse		"Forbidden"
// @Failure		 500  {object} common.ErrorResponse		"Internal Server Error"
// @Router       /api/v1/test/ [get]
// GetTestMsg return test message.
func (ctr *controller) GetTestMsg(c *fiber.Ctx) error {
	return c.JSON(common.GeneralSuccessResponse{
		OK:      true,
		Message: "test",
	})
}
