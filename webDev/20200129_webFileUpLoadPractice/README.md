# 檔案上傳實作

## 技術重點
1. 前端html的form格式
```html
<!-- 一定要設定enctype="multipart/form-data" -->
<form action="/upload" method="POST" enctype="multipart/form-data">
<!-- input type 選file -->
Files: <input type="file" name="file" ><br><br>
<input type="submit" value="Submit">
</form>
```

2. 後端
```go
// 上傳相片
r.POST("/upload", func(c *gin.Context){
	// 由前端req提取檔案的方式
	file, err := c.FormFile("file")
}
```