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
./generate_proto.sh
```

C# 코드 생성에는 `grpc_csharp_plugin`이 필요합니다 (`brew install grpc`로 설치).

생성되는 파일:
- `proto/*.pb.go`, `proto/*_grpc.pb.go` — Go 메시지 타입 및 서비스 인터페이스 (server)
- `../client/Assets/Scripts/Proto/*.cs` — C# 메시지 타입 및 서비스 인터페이스 (Unity client)

---

## 서버 실행

```bash
cd server
go run ./cmd/
```
