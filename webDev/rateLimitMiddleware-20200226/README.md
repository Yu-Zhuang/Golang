# Dcard backend 實習生 應徵作業
## Rate Limit Middleware

### 實測與使用

#### 實測
1. 設定與啟動Redis (一行指令即可)
> 簡單快速完成此步驟 -> 使用docker一行指令即可:
```sh
sudo docker run redis -d -p 6379:6479
```

2. 至 ./bin 目錄下依照作業系統環境選擇啟動執行檔 
> 預設rate limit= 3 ; duration= 1 min 方便進行測試, 需更改可在./conf/conf.go中更改後重新go build即可

#### 使用
測試範例中將該middleware寫在
1. ./controller/middleware.go
2. ./logic/logic.go
若需要移植使用可參考範例檔案中middleware使用方式

### 實作技術列表
1. language: Go(gin framwork)
2. database: Redis

### 實作細節闡述