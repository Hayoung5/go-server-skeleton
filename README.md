# go-server-skeleton (server skeleton using mvc model with go-lang)

Server-client 아키텍처 패턴에서 server 파트를 표현한 예제 코드와 폴더 구조 입니다.
<br/>예시 코드는 Mongo DB 사용을 가정하고 작성되어 있습니다.

## 폴더 구조
폴더 구조는 아래와 같습니다.
```
./go-server-skeleton
├── LICENSE
├── Makefile
├── README.md
├── config
│   ├── config.go
│   └── config.toml
├── controller
│   └── controller.go
├── go.mod
├── go.sum
├── logger
│   └── logger.go
├── logs
│   └── go_logger_2022-12-19.log
├── main.go
├── model
│   └── model.go
└── router
    └── router.go

6 directories, 13 files
```

주요 파일 및 디렉터리 설명은 다음과 같습니다.
* main.go : 서버의 엔트리 포인트 역할
* model : db의 출력 형태를 정의하고, 데이터를 핸들링
* controller: model과 view를 컨트롤 하는 구성으로 api 입출력의 시작점
* router: 서버의 라우트를 정의
* config: 서버의 설정 파일을 포함
* logger: 로그작성에 대한 정의
* log: 로그 파일 저장 

**프로세스의 흐름**
request → router → controller → model → controller → response

## 사용 방법
1. 고 설치 및 GOPATH 설정 후 go 폴더 아래 bin, pkg, src 폴더를 생성
2. src 밑으로 해당 레파지토리를 clone함
3. 패키지 초기화 및 실행

패키지 모듈 초기화
```
go mod init
```
패키지 모듈 정렬, 설정 및 go.sum 내용의 패키지 재설치
```
go mod tidy
```

실행
```
go run main.go
```

컴파일 및 실행
```
go bulid -o ready main.go
./ready
```
