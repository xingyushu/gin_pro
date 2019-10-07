package main  

import   (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
) 


func main() {
	r :=gin.Default()
    r.POST("/upload",func(c *gin.Context) {
    	form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(200, fmt.Sprintf("%d files uploaded!", len(files)))
    })

    r.Run()
}


curl -X POST http://localhost:8080/upload \
  -F "upload[]=@/learn.txt" \
  -F "upload[]=@/antdnote.txt" \
  -H "Content-Type: multipart/form-data"