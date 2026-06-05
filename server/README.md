# go-grpc server

## 사전 준비

### 1. protoc 설치
```bash
brew install protobuf
```

### 2. Go 플러그인 설치
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 3. PATH 설정 (~/.zshrc에 추가)
```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

---

## protoc 실행

`server/` 디렉토리에서 실행:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/auth.proto
```

생성되는 파일:
- `proto/auth.pb.go` — 메시지 타입
- `proto/auth_grpc.pb.go` — 서비스 인터페이스 및 클라이언트

---

## 서버 실행

```bash
cd server
go run ./cmd/
```
