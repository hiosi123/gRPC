package explain

// why choose go?
// we live in times when have so many languages
// easy answer - it is one of the fastest and efficient languages out there
// stand out feature is concurrency, multiple task simnoutinously, go and channel, network server, distribution,
// performance, run fast and consume few resources, high performance application, readability, simplicity
// built in testing support, tooling issues
// secure and

// Protocol buffer, efficiency over json
// Less CPU intensive, FASTER, mobile and microcontrollers
// Efficient Serialiation/Deserialization

// HTTP2
// HTTP1 - it opens http connection for each connection, and in the case of the 100 tile image, I would open and close 100 tcp connection
// HTTP1 does not compress headers, and plaintest headers make latency
// related to the first point http/1 actually only accepts the request and response pattern

// http/2 will open long lasting connection
// server can push the data, on one connection
// multiplexing, can send multiple request in parallel
// headers and body are both compressed
// http/2 ssl connection is required by default
// less chatter, less bandwitdth, more security, all of that for free, by using grpc framework

// unary, server streaming, client streaming, bi directional streaming

// Unary
// - client will send 1 request and server will send 1 resposne

// server streaming
// - client will send one request and then the server can rqeust 1 or more response
// ex) getting a list of nearest taxes

// client stereaming
// - client can send multiple request and server sends 1 response

// bi directional
// multi multi, response can arrive any order you want,

// stream 이란 단어로 정의한다

// Scalability in gRPC
// server: async, server can receive many requests as it want
// client: async or blocking

// google: 10 B requests/ sec
// Sercurity in gRPC
// Schema based serialization
// Easy SSL certificates initialization, tls initlization
// interceptors for Auth

// gRPC: protocol buffers: strick schmea, http2, streaming, bi directional, free design

// REST: JSON, http1, unary, client - server, (get/ post/ update/ delete)

// proto 생성
//protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto
//protoc -Igreet/proto --go_out=. --go_opt=module=github.com/hiosi123/gRPC --go-grpc_out=. --go-grpc_opt=module=github.com/hiosi123/gRPC greet/proto/dummy.proto

// mac - chmod +x ssl.sh

// 20251108
