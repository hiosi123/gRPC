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
// Kubectl

// kubectl get all
// (Section 5): list all objects that you’ve created. Pods at first, later,
// ReplicaSets, Deployments and Services

// kubectl apply –f <yaml file>
// (Section 5): either creates or updates resources depending on the
// contents of the yaml file

// kubectl apply –f .
// (Section 7): apply all yaml files found in the current directory
// Kubernetes - Quick Command Reference 2

// kubectl describe pod <name of pod>
// (Section 5): gives full information about the specified pod

// kubectl exec –it <pod name> <command>
// (Section 5): execute the specified command in the pod’s container.
// Doesn’t work well in Cygwin.

// kubectl get (pod | po | service | svc | rs | replicaset | deployment |
// deploy)
// (Section 6): get all pods or services. Later in the course, replicasets and
// deployments.

// kubectl get po --show-labels
// (Section 6): get all pods and their labels

// kubectl get po --show-labels -l {name}={value}
// (Section 6): get all pods matching the specified name:value pair

// kubectl delete po <pod name>
// (Section 8): delete the named pod. Can also delete svc, rs, deploy

// kubectl delete po --all
// (Section 8): delete all pods (also svc, rs, deploy)
// Deployment Management

// kubectl rollout status deploy <name of deployment>
// (Section 9): get the status of the named deployment

// kubectl rollout history deploy <name of deployment>
// (Section 9): get the previous versions of the deployment

// kubectl rollout undo deploy <name of deployment>
// (Section 9): go back one version in the deployment. Also optionally --
// to-revision=<revision number>
// We recommend this is used only in stressful emergency situations! Your
// YAML will now be out of date with the live deployment!
