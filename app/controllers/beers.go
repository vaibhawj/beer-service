package controllers

import (
	"github.com/revel/revel"
	"github.com/vaibhawj/beer-service/app/services"
)

type Beers struct {
	*revel.Controller
}

func (c Beers) List() revel.Result {
	return c.RenderJSON(services.ListBeers())
}
