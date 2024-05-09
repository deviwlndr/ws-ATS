package controller

import (
   
	
	"github.com/aiteung/musik"
	cek"github.com/deviwlndr/undangan_rapat/module"
	
	//"github.com/deviwlndr/ws-ATS/config"
	
	"github.com/gofiber/fiber/v2"
	
	
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetRapatMakrab(c *fiber.Ctx) error {
	ps := cek.GetAllRapatMakrab
	return c.JSON(ps)
}

