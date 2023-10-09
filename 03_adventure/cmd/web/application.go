package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/milkymilky0116/gophercises-practice/03_adventure/internal/models"
)

type Application struct {
	Filename string
	Router   *httprouter.Router
	Stories  map[string]models.AdventureStory
}
