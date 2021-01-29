package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func main() {
	// 印出server IP位置
	fmt.Println(GetLocalIP())
	//  路由註冊
	r := gin.Default()

	// 首頁
	r.GET("/", func(c *gin.Context){
		f := "index.html"
		c.File(f)
	})
	// 上傳相片
	r.POST("/upload", func(c *gin.Context){
		file, err := c.FormFile("file")
		if err !=nil {
			fmt.Println(err.Error())
			c.JSON(404, gin.H{
				"status" : "false",
			})
		} else {
			// Upload the file to specific dst.
			dst := "./pic/" + file.Filename
			fmt.Println(dst)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(404, gin.H{
					"status" : "false",
				})				
			} else {
				c.JSON(200, gin.H{
					"status" : "ok",
					"filename" : file.Filename,
				})				
			}
		}
	})
	// 查看相片
  	r.GET("/show/:f", func(c *gin.Context) {
            f := "pic/" + c.Param("f")
            c.File(f)
	})


	r.Run()

}
