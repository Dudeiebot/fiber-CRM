package lead

import (
	"github.com/dudeiebot/fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct { //struct is basically the datatype(like string and int created by golang) you are creating on your own and it can be a combination of string and int
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead) //these kind of save us effort from marshalling and unmarshalling
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil { //you get accesss to bodyparser in your fiber and it parses the body of the lead date the user is sending to you
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead //lead is the name of the variable whereby Lead is my own datatype of struct here
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with id")
		return
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}
