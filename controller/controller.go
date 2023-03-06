package controller

import (
	"fmt"
	"go-server-skeleton/model"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{}
	fmt.Println(rep)
	return r, nil
}

func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
	return
}

func (p *Controller) GetEmployee(c *gin.Context) {
	r, _ := model.NewModel()
	fmt.Println(r.GetEmployee())
	c.JSON(200, r.GetEmployee())
	return
}
