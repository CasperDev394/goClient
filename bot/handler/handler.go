package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/CasperDev394/goClient/getinfo"
	"github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	c.R.GET("/bot", h.Bot)

	g := c.R.Group("/bot")
	simple := c.R.Group("/bot/simple")
	coin := c.R.Group("/bot/coins")

	g.GET("/ping", h.Ping)
	simple.GET("/supported_vs_currencies", h.SimpleSupportedVSCurrencies)
	simple.GET("/price")
	simple.GET("/token_price/:id", h.TokenPrice)
	coin.GET("/list", h.CoinList)
	coin.GET("/:id/history", h.CoinHistory)
	coin.GET("/:id/ohlc", h.CoinOhlc)
}

func (h *Handler) Bot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Ping(c *gin.Context) {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.Ping()
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (h *Handler) SimpleSupportedVSCurrencies(c *gin.Context) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.SimpleSupportedVSCurrencies()
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (h *Handler) CoinList(c *gin.Context) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.CoinList()
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (h *Handler) CoinHistory(c *gin.Context) {
	id := c.Param("id")
	//date, _ := c.Params.Get("date")
	var queryParams struct {
		Date string `form:"date" json:"date"`
	}
	if c.BindQuery(&queryParams) != nil {
		log.Println("Ошибка параметров")
	}

	log.Println("H date - ", queryParams.Date)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.CoinHistory(id, queryParams.Date)
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"id":     c.Param("id"),
	})
}

func (h *Handler) CoinOhlc(c *gin.Context) {
	id := c.Param("id")
	//date, _ := c.Params.Get("date")
	var queryParams struct {
		Days       string `form:"days" json:"days"`
		VsCurrency string `form:"vs_currency" json:"vs_currency"`
	}
	if c.BindQuery(&queryParams) != nil {
		log.Println("Ошибка параметров")
	}

	log.Println("H date - ", queryParams.Days, queryParams.VsCurrency)

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	r := getinfo.NewClient(httpClient)
	result, err := r.CoinOhlc(id, queryParams.Days, queryParams.VsCurrency)
	if err != nil {
		log.Printf("Fail %s", err)
	}
	//log.Println(result)
	//log.Println(reflect.TypeOf(result))

	c.JSON(http.StatusOK, gin.H{
		"result": result,
		"id":     c.Param("id"),
	})
}

func (h *Handler) TokenPrice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": c.Param("id"),
	})

}
