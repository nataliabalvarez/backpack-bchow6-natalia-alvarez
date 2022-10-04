package main
import "github.com/gin-gonic/gin"

func main(){
	router := gin.Default() //enginee de gin

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"Mensaje": "bienvenides a go web \n",
		})
	})

	router.Run() //dentro de (:puerto)
}