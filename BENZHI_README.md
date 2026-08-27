# graindry

graindry 是一套工业粮食干燥塔工业过程控制系统，用于区段含水梯度、风门配风与塔级状态联锁。

## Requirements

- Go 1.22+ (container image uses golang:1.22)
- `GOTOOLCHAIN=local` recommended on host when using a pinned toolchain

## Build

```bash
export GOTOOLCHAIN=local
go build ./...
```

## Run

```bash
export GOTOOLCHAIN=local
go run ./cmd/graindry
```

Open http://127.0.0.1:8080/

Default ingest secret: `dev-ingest-secret`.

## Test

```bash
export GOTOOLCHAIN=local
go test ./... -count=1
```

## Docker (benzhi)

Must build **linux/amd64** and **linux/arm64**:

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh graindry linux/amd64
./build_benzhi_docker.sh graindry linux/arm64
docker run -it graindry:latest
# inside container:
export GOTOOLCHAIN=local
go version
go build ./...
go test ./... -count=1
```
