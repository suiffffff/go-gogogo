package main

import "github.com/gin-gonic/gin"

type Bob struct {
	Name  string    `json:"name"`
	Score []float64 `json:"score"`
}

// 下载go get -u github.com/gin-gonic/gin
func main() {
	r := gin.Default()
	r.GET("/talk", func(c *gin.Context) {
		msg := c.Query("msg")
		if msg == "ping" {
			c.JSON(200, gin.H{"data": "pong"})
		} else if msg == "helloserver" {
			c.JSON(200, gin.H{"data": "helloclient"})
		} else {
			c.JSON(400, gin.H{"data": "i do not konw"})
		}

	})
	r.Static("/cat.jpg", "./cat.jpg")
	r.POST("/ChaChaBob", func(c *gin.Context) {
		var hibob Bob
		if err := c.ShouldBindJSON(&hibob); err != nil {
			c.JSON(400, gin.H{
				"error":   "JSON 格式不正确",
				"details": err.Error(),
			})
			return
		}
		var sum, average float64
		var number int
		for _, i := range hibob.Score {
			sum += i
			number++
		}
		average = sum / float64(number)
		c.JSON(200, gin.H{
			"averaage": average,
		})
	})
	//代码手写，但大部分课件内容靠ai答疑，我有点不敢问学长，模型用的Gemini2.5Pro，若代码有误那应该是ai的误导，望指出
	//下面的应该是JavaScript的内容，这个也确实不了解，没办法，只能先学会用了
	//curl -Method POST -Uri http://localhost:8080/ChaChaBob -ContentType "application/json" -Body '{"name": "Bob", "score": [69,91,78]}'
	r.Run()
}
