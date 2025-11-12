GitHub README 스타일로 변환해드리겠습니다!

```markdown
# 🚀 Go & gRPC 학습 가이드

> Go 언어와 gRPC를 활용한 고성능 분산 시스템 개발 학습 노트

## 📋 목차

- [Go 언어를 선택하는 이유](#-go-언어를-선택하는-이유)
- [Protocol Buffer](#-protocol-buffer)
- [HTTP/1 vs HTTP/2](#-http1-vs-http2)
- [gRPC 스트리밍 타입](#-grpc-스트리밍-타입)
- [gRPC의 확장성](#-grpc의-확장성)
- [gRPC의 보안](#-grpc의-보안)
- [gRPC vs REST 비교](#-grpc-vs-rest-비교)
- [Proto 파일 생성 명령어](#-proto-파일-생성-명령어)

---

## 🎯 Go 언어를 선택하는 이유

Go는 현대의 수많은 프로그래밍 언어 중에서도 특별한 위치를 차지하고 있습니다.

### 주요 장점

- ⚡ **속도와 효율성**: 가장 빠르고 효율적인 언어 중 하나
- 🔄 **동시성 처리**: 고루틴(goroutine)과 채널(channel)을 통해 여러 작업을 동시에 처리
- 🌐 **네트워크 서버 및 분산 시스템**에 최적화
- 🚀 **고성능**: 빠른 실행 속도와 적은 리소스 소비
- 📖 **가독성과 단순성**: 코드가 읽기 쉽고 이해하기 쉬움
- 🧪 **내장된 테스팅 지원**: 테스트 작성이 용이
- 🛠️ **강력한 도구 생태계**
- 🔒 **보안성**

---

## 📦 Protocol Buffer

JSON 대비 효율성이 뛰어난 데이터 직렬화 포맷입니다.

### 장점

- ⚙️ CPU 사용량이 적음
- ⚡ 더 빠른 처리 속도
- 📱 모바일 및 마이크로컨트롤러 환경에 적합
- 🔄 효율적인 직렬화/역직렬화 성능

---

## 🌐 HTTP/1 vs HTTP/2

### HTTP/1의 한계

- ❌ 각 요청마다 새로운 TCP 연결을 열고 닫음 (예: 이미지 100개 = TCP 연결 100개)
- ❌ 헤더를 압축하지 않아 평문 헤더가 지연 시간을 증가시킴
- ❌ 요청-응답 패턴만 지원

### HTTP/2의 장점

- ✅ **장기 지속 연결**: 한 번 열린 연결을 계속 사용
- ✅ **서버 푸시**: 서버가 하나의 연결로 데이터를 능동적으로 전송 가능
- ✅ **멀티플렉싱**: 하나의 연결에서 여러 요청을 병렬로 전송
- ✅ **압축**: 헤더와 바디 모두 압축
- ✅ **SSL 필수**: 기본적으로 SSL 연결 요구
- ✅ **효율성**: 적은 통신량, 적은 대역폭, 높은 보안성 (gRPC 프레임워크 사용 시 자동 제공)

---

## 📡 gRPC 스트리밍 타입

### 1️⃣ Unary (단일)

클라이언트가 1개의 요청을 보내고, 서버가 1개의 응답을 반환

```
Client ---(1 request)---> Server
Client <--(1 response)--- Server
```

### 2️⃣ Server Streaming (서버 스트리밍)

클라이언트가 1개의 요청을 보내고, 서버가 1개 이상의 응답을 반환

```
Client ---(1 request)-----> Server
Client <--(N responses)----- Server
```

**사용 예시**: 근처 택시 목록 조회

### 3️⃣ Client Streaming (클라이언트 스트리밍)

클라이언트가 여러 개의 요청을 보내고, 서버가 1개의 응답을 반환

```
Client ---(N requests)----> Server
Client <--(1 response)------ Server
```

### 4️⃣ Bi-directional Streaming (양방향 스트리밍)

클라이언트와 서버 모두 여러 개의 요청/응답을 주고받음

```
Client <---(N requests/responses)---> Server
```

응답 순서는 자유롭게 결정 가능

> 💡 **참고**: 스트리밍은 `stream` 키워드로 정의합니다.

---

## 📈 gRPC의 확장성 (Scalability)

- **서버**: 비동기 방식으로 원하는 만큼 많은 요청을 받을 수 있음
- **클라이언트**: 비동기 또는 블로킹 방식 선택 가능

### 실제 사례

> 🌟 Google은 초당 **100억 건**의 요청을 처리

---

## 🔐 gRPC의 보안 (Security)

- **스키마 기반 직렬화**: 타입 안정성 보장
- **SSL 인증서 초기화 용이**: TLS 설정이 간편
- **인터셉터를 통한 인증**: 요청/응답 중간에 인증 로직 삽입 가능

---

## ⚖️ gRPC vs REST 비교

| 구분 | gRPC | REST |
|------|------|------|
| 📄 **데이터 포맷** | Protocol Buffers (엄격한 스키마) | JSON |
| 🌐 **프로토콜** | HTTP/2 | HTTP/1 |
| 📡 **통신 방식** | Unary, Streaming, 양방향 | Unary (요청-응답) |
| 🏗️ **아키텍처** | 클라이언트-서버, 양방향 가능 | 클라이언트-서버 |
| 🔧 **메서드** | 자유로운 디자인 | GET/POST/PUT/DELETE |

---

## 🛠️ Proto 파일 생성 명령어

### 기본 명령어

```bash
protoc -Igreet/proto \
  --go_out=. \
  --go-grpc_out=. \
  greet/proto/dummy.proto
```

### 모듈 경로 지정

```bash
protoc -Igreet/proto \
  --go_out=. \
  --go_opt=module=github.com/hiosi123/gRPC \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/hiosi123/gRPC \
  greet/proto/dummy.proto
```

---

## 📚 참고 자료

- [gRPC 공식 문서](https://grpc.io/docs/)
- [Protocol Buffers 가이드](https://developers.google.com/protocol-buffers)
- [Go 공식 문서](https://go.dev/doc/)

---

## 📝 License

MIT License

---

**Made with ❤️ by [Your Name]**
```

