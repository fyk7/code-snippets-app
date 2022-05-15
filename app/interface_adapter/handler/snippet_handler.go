package interface_adapter

import (
	"net/http"
	"strconv"

	"github.com/fyk7/code-snippets-app/app/usecase"
	"github.com/labstack/echo/v4"
)

type SnippetHandler interface {
	GetSnippetByID(c echo.Context) error
	FindSnippetByKeyWord(c echo.Context) error
	ListByTag(c echo.Context) error
	PostSnippet(c echo.Context) error
	AssociateWithTag(c echo.Context) error
	PutSnippet(c echo.Context) error
}

type snippetHandler struct {
	snippetService usecase.SnippetService
}

func NewSnippetHandler(e *echo.Echo, s usecase.SnippetService) {
	handler := &snippetHandler{
		snippetService: s,
	}
	e.GET("/snippets/:snippet_id", handler.GetSnippetByID)
	e.GET("/snippets/search", handler.FindSnippetByKeyWord)
	e.GET("/snippets/tags/:tag_id", handler.ListByTag)
	e.POST("/snippets", handler.PostSnippet)
	e.POST("/snippets/associate", handler.AssociateWithTag)
}

func (h *snippetHandler) GetSnippetByID(c echo.Context) error {
	ctx := c.Request().Context()
	snippetID, err := strconv.Atoi(c.Param("snippet_id"))
	if err != nil {
		return handleError(c, err)
	}
	snippet, err := h.snippetService.GetByID(ctx, uint64(snippetID))
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, snippet)
}

func (h *snippetHandler) FindSnippetByKeyWord(c echo.Context) error {
	ctx := c.Request().Context()
	snippetKeyword := c.QueryParam("snippet_keyword")
	snippet, err := h.snippetService.GetByKeyWord(ctx, snippetKeyword)
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, snippet)
}

func (h *snippetHandler) ListByTag(c echo.Context) error {
	ctx := c.Request().Context()
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return handleError(c, err)
	}
	snippets, err := h.snippetService.GetByKeyTagID(ctx, uint64(tagID))
	if err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusOK, snippets)
}

func (h *snippetHandler) PostSnippet(c echo.Context) error {
	var req SnippetPostReq
	if err := c.Bind(&req); err != nil {
		handleError(c, err)
	}
	ctx := c.Request().Context()
	UserID := 00000 // dummy user
	if err := h.snippetService.Create(ctx, req.ConvertToModel(), uint64(UserID)); err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusCreated, "Successfuly Created.")
}

func (h *snippetHandler) AssociateWithTag(c echo.Context) error {
	ctx := c.Request().Context()
	snippetID, err := strconv.Atoi(c.Param("snippet_id"))
	if err != nil {
		return handleError(c, err)
	}
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		return handleError(c, err)
	}
	UserID := 00000
	if err := h.snippetService.AssociateWithTag(ctx, int64(snippetID), int64(tagID), int64(UserID)); err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusCreated, "Successfuly Created.")
}

func (h *snippetHandler) PutSnippet(c echo.Context) error {
	var req SnippetPutReq
	if err := c.Bind(&req); err != nil {
		handleError(c, err)
	}
	ctx := c.Request().Context()
	UserID := 00000
	if err := h.snippetService.Update(ctx, req.ConvertToModel(), uint64(UserID)); err != nil {
		return handleError(c, err)
	}
	return c.JSON(http.StatusCreated, "Successfuly Created.")
}
