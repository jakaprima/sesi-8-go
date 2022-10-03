package controller

import (
	"net/http"
	"strconv"

	"golang-swagger/httputil"
	"golang-swagger/models"

	"github.com/gin-gonic/gin"
)

// PingPeoples godoc
// @Summary      ping people
// @Description  do ping
// @Tags         peoples
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "pong"
// @Failure      400  {string}  string  "ok"
// @Failure      404  {string}  string  "ok"
// @Failure      500  {string}  string  "ok"
// @Router       /peoples/ping [get]
func (c *Controller) PingPeoples(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong people")
}

// ShowPeople godoc
// @Summary      Show an people
// @Description  get string by ID
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "People ID"
// @Success      200  {object}  models.People
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /peoples/{id} [get]
func (c *Controller) ShowPeople(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	people, err := models.PeopleOne(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, people)
}

// ListPeoples godoc
// @Summary      List peoples
// @Description  get peoples
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(string)
// @Success      200  {array}   models.People
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /peoples [get]
func (c *Controller) ListPeoples(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	peoples, err := models.PeoplesAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, peoples)
}

// Add People godoc
// @Summary      Add an people
// @Description  add by json people
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        people  body      models.AddPeople  true  "Add people"
// @Success      200      {object}  models.People
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /peoples [post]
func (c *Controller) AddPeople(ctx *gin.Context) {
	var addPeople models.AddPeople
	if err := ctx.ShouldBindJSON(&addPeople); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addPeople.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	people := models.People{
		Name:    addPeople.Name,
		Address: addPeople.Address,
	}
	lastID, err := people.InsertPeople()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	people.ID = lastID
	ctx.JSON(http.StatusOK, people)
}

// UpdatePeople godoc
// @Summary      Update an people
// @Description  Update by json people
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "People ID"
// @Param        people  body      models.UpdatePeople  true  "Update people"
// @Success      200      {object}  models.People
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /peoples/{id} [patch]
func (c *Controller) UpdatePeople(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updatePeople models.UpdatePeople
	if err := ctx.ShouldBindJSON(&updatePeople); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	people := models.People{
		ID:   aid,
		Name: updatePeople.Name,
	}
	err = people.UpdatePeopleAction()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, people)
}

// DeletePeople godoc
// @Summary      Delete an people
// @Description  Delete by people ID
// @Tags         peoples
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "People ID"  Format(int64)
// @Success      204  {object}  models.People
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /peoples/{id} [delete]
func (c *Controller) DeletePeople(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = models.DeletePeopleAction(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
