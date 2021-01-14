# squaaat-api


# cli

### go build & deploy
``` bash
GOARCH=amd64 GOOS=linux go build -o dist/sq 
```

``` bash
terraform init ./terraform # if didn't run
terraform apply ./terraform
```


### secrets 환경변수

- create
SSM 에 값 생성하기
``` bash
./scripts/secrets/create.sh -r ap-northeast-2 -a squaaat-api -e alpha
```

- printout
SSM 에 있는 값 파일로 만들기
``` bash
./scripts/secrets/printout.sh -r ap-northeast-2 -a squaaat-api -e alpha -o ./
```

- update
SSM 에 있는 값 파일로 만든거를 기반으로 내용 업데이트 하기
``` bash
./scripts/secrets/update.sh -r ap-northeast-2 -a squaaat-api -e alpha -i ./
```

### swagger

- create swagger yml

``` bash
swagger generate spec -o ./swagger.yml --scan-models
```

- validate swagger yml

``` bash
swagger validate ./swagger.yml
```

### gorm

- DB 초기화
``` bash
go run main.go gorm create
```

- DB 삭제
``` bash
go run main.go gorm clean
```

- DB 초기화 & 삭제
``` bash
go run main.go gorm re-create
```

- Migration 코드 생성
``` bash
go run main.go gorm migrate create -v (default: yyyymmddHHMM)
```

- Migration 코드 실행 
``` bash
go run main.go gorm migrate sync
``` 

