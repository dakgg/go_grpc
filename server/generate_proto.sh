#!/usr/bin/env bash
set -euo pipefail

# proto/*.proto -> Go(server) + C#(Unity client) 코드 동시 생성
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# protoc-gen-go / protoc-gen-go-grpc 가 설치되는 경로를 PATH에 추가
export PATH="$(go env GOPATH)/bin:$PATH"

PROTO_DIR="proto"
CSHARP_OUT="../client/Assets/Scripts/Proto"

if ! command -v grpc_csharp_plugin >/dev/null 2>&1; then
  echo "grpc_csharp_plugin을 찾을 수 없습니다. 'brew install grpc'로 설치 후 다시 실행하세요." >&2
  exit 1
fi

mkdir -p "$CSHARP_OUT"

# Go: server/proto/*.pb.go, *_grpc.pb.go
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       "$PROTO_DIR"/*.proto

# C#: ../client/Assets/Scripts/Proto/*.cs
protoc --csharp_out="$CSHARP_OUT" \
       --grpc_out="$CSHARP_OUT" \
       --plugin=protoc-gen-grpc="$(command -v grpc_csharp_plugin)" \
       "$PROTO_DIR"/*.proto

echo "Go  -> $SCRIPT_DIR/$PROTO_DIR/"
echo "C#  -> $SCRIPT_DIR/$CSHARP_OUT/"
