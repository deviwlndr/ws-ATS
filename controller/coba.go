package controller

import (
    
	"github.com/deviwlndr/undangan_rapat"
	"github.com/gofiber/fiber/v2"
	
	
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := _714220054.GetAllUndanganRapat()
	return c.JSON(ipaddr)
}

