package controllers

import (
	"context"
	"database/sql"
	"demo/blog/core"
	"demo/blog/db"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (a *ArticleController) Create(c *gin.Context) {
	m := &struct {
		Title   string   `json:"title" binding:"required"`
		Content string   `json:"content" binding:"required"`
		Tag     []string `json:"tag" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(m); err != nil {
		c.String(http.StatusBadRequest, "the body should be Article %v", err)
		return
	}
	log.Printf("%v", m)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	if _, err := db.GetDB().ExecContext(ctx, "insert into t_article(title,content,`tag`) values(?,?,?);",
		m.Title, m.Content, strings.Join(m.Tag, ",")); err != nil {
		c.String(http.StatusOK, "create fail %v", err)
		return
	}

	c.String(http.StatusOK, "create ok")
}

func (a *ArticleController) Get(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	m := &core.Article{}
	err := db.GetDB().QueryRowContext(ctx, "select id,title,content,tag,createdTime from t_article where id=?;", id).Scan(&m.Id, &m.Title, &m.Content, &m.Tag, &m.CreatedTime)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusOK, m)
	case err != nil:
		c.String(http.StatusInternalServerError, "query error: %v", err)
	default:
		c.JSON(http.StatusOK, m)
	}
}
