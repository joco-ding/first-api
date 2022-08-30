package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type filmTS struct {
	ID       string `json:"id"`
	Judul    string `json:"judul"`
	Kategori string `json:"kategori"`
}

var (
	films = map[string]filmTS{
		"1": {ID: "1", Judul: "Gods of Egypt", Kategori: "Laga & Petualangan"},
		"2": {ID: "2", Judul: "Spider-man Homecoming", Kategori: "Laga & Petualangan"},
		"3": {ID: "3", Judul: "Now You See Me", Kategori: "Kriminal"},
	}
)

func getAllFilm(c *gin.Context) {
	c.JSON(http.StatusOK, films)
}

func getDetail(c *gin.Context) {
	id := c.Param("id")

	if f, exists := films[id]; exists {
		c.JSON(http.StatusOK, f)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"ok": false, "message": "film tidak ditemukan"})
}

func addFilm(c *gin.Context) {
	var newfilm filmTS
	if err := c.ShouldBindJSON(&newfilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "message": "data film tidak valid"})
		return
	}
	if _, exists := films[newfilm.ID]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "message": "id film sudah digunakan"})
		return
	}
	films[newfilm.ID] = newfilm
	c.JSON(http.StatusOK, newfilm)
}

func editFilm(c *gin.Context) {
	id := c.Param("id")
	var newfilm filmTS
	if err := c.ShouldBindJSON(&newfilm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "message": "data film tidak valid"})
		return
	}
	if _, exists := films[id]; exists {
		newfilm.ID = id
		films[id] = newfilm
		c.JSON(http.StatusOK, newfilm)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"ok": false, "message": "data film tidak ditemukan"})
}

func delFilm(c *gin.Context) {
	id := c.Param("id")

	if _, exists := films[id]; exists {
		delete(films, id)
		c.JSON(http.StatusOK, gin.H{"ok": true, "message": "film berhasil dihapus"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"ok": false, "message": "film tidak ditemukan"})
}

func main() {
	router := gin.Default()
	router.GET("/films", getAllFilm)
	router.GET("/film/:id", getDetail)
	router.PATCH("/film/:id", editFilm)
	router.DELETE("/film/:id", delFilm)
	router.POST("/film", addFilm)
	router.Run(":8081")
}
