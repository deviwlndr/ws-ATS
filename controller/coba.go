package controller

import (
   
	//"fmt"
	"github.com/aiteung/musik"
	cek"github.com/deviwlndr/undangan_rapat/module"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/deviwlndr/ws-ATS/config"
	//"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"
	//"net/http"
	
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllUndanganRapat(c *fiber.Ctx) error {
	ps := cek.GetAllUndanganRapat
	return c.JSON(ps)
}

