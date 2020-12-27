# squaaat-api


# cli
### gorm

- DB 초기화
``` go
go run main.go gorm init
```

- DB 초기화
``` go
go run main.go gorm clean
```

- Migration 코드 생성
``` go
go run main.go gorm migrate create -v (default: yyyymmddHHMM)
```

- Migration 코드 실행 
``` go
go run main.go gorm migrate sync
``` 

