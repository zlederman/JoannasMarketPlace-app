package main


import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"os"
	
)

func main()  {

    // Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("client/build", true)))
	fmt.Println(parseJSON())
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", homePage)
		api.GET("/items",)
		api.POST("/order")
	}

	// Start and run the server
	router.Run(":5000")
}

func homePage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
				"message": "pong",
	})
}

func parseJSON() (products, error){
	r, err := os.Open("data.json")
	if err != nil {
		panic(err)

	}
	d:= json.NewDecoder(r)
	var p products
	if err := d.Decode(&p); err != nil{
		return nil, err
	}
	return p, nil
}
func items(c *gin.Context){
	 p, err := parseJSON()
	 if(err != nil){
		 panic(err)
	 }
	 payload, err := json.Marshal(&p)
	
	 c.JSON(http.StatusOK, gin.H{
		 "payload":payload,
	 })
	//obtain json values the package them into an array to be sent as a get request

}

type Product struct {
	Price int    `json:"price"`
	Type  string `json:"type"`
}
type products 
