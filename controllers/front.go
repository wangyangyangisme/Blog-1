package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
	"time"
)

type FrontController struct {
	beego.Controller
}

func (this *FrontController) Prepare() {
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
}

func (this *FrontController) Home() {
	value := "2014-04-20T15:13:09+08:00"
	t, _ := time.Parse(time.RFC3339, value)
	timeStr := fmt.Sprintf("%s %d, %d", t.Month().String(), t.Day(), t.Year())
	fmt.Println(timeStr)

	// Main Nav
	this.Data["HomeActive"] = "active"

	// Data Source
	this.Data["Entries"] = models.PublishedEntries()

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["LeftPage"] = "disabled"

	this.TplNames = "entry-list.tpl"
	// this.Ctx.WriteString("<h1>Hello world!</h1>")

	renderTemplate(this.Ctx, "views/entry-list.amber", this.Data)

	// fmt.Println(this.Data)
	// compiler := amber.New()
	// err := compiler.ParseFile("views/entry-list.amber")
	// if err != nil {
	// 	panic(err)
	// }
	// tpl, err := compiler.Compile()
	// if err != nil {
	// 	panic(err)
	// }

	// var content bytes.Buffer
	// _, err = json.Marshal(this.Data)
	// err = tpl.Execute(&content, this.Data)
	// if err != nil {
	// 	panic(err)
	// }
	// this.Ctx.WriteString(content.String())
}

func (this *FrontController) Collections() {
	this.TplNames = "collection-list.tpl"

	// Main Nav
	this.Data["CollectionActive"] = "active"

	// Data Source
	collections, err := models.AllCollections()
	if err != nil {
		panic(err)
	}
	this.Data["Collections"] = collections

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["LeftPage"] = "disabled"

	renderTemplate(this.Ctx, "views/collection-list.amber", this.Data)
}

func (this *FrontController) Collection() {
	cid := this.Ctx.Input.Param(":id")
	collection, err := models.CollectionById(cid)
	if err != nil {
		panic(err)
	}

	entries, err := models.EntriesByCollection(cid)
	if err != nil {
		// panic(err)
	}

	this.TplNames = "entry-list.tpl"
	this.Data["Title"] = collection.Title
	this.Data["Subtitle"] = collection.Subtitle
	this.Data["Entries"] = entries

	renderTemplate(this.Ctx, "views/entry-list.amber", this.Data)
}

func (this *FrontController) Entry() {
	eid := this.Ctx.Input.Param(":id")
	entry, _ := models.EntryById(eid)
	// fmt.Println(entry)

	this.TplNames = "entry.tpl"

	fmt.Println(entry.Content)
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true

	renderTemplate(this.Ctx, "views/entry.amber", this.Data)
}
