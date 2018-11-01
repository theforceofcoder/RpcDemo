RPC 應用範例
===================
邊學邊寫邊設計  
RPC demo across languages

#### System requirement:  
-------------------------
- go 1.11.1
- python 3.6
- gRPC

#### 我的系統
-----------------------  
macOS Mojave 10.14

#### for go
-----------------------
1. Homebrew安裝 protobuf
  * brew install protobuf
1. install gRPC
  * go get -u google.golang.org/grpc
1. **Protocol Buffers**建立RPC service和Request Response物件
1. 安裝go protobuf compile
  * go get -u github.com/golang/protobuf/protoc-gen-go
1. 編譯.proto檔案轉成go程式，proto檔案放在rpctoute目錄裡，-I 後面要放目錄名稱及檔案路徑和檔案名稱，plugins=grpc:目錄名稱
  * protoc -I rpcroute/ rpcroute/route.proto --go_out=plugins=grpc:rpcroute
  * 產生出的檔案是route.pb.go
1. gprc framework建立rpc應用範例，  
1. 經由Protocol Buffers使用不同程式語言建立rpc client



#### for Python
-----------------------
我是用conda建立Python虛擬環境  
[下載Anaconda](https://conda.io/docs/user-guide/install/download.html)  
[conda管理環境](https://conda.io/docs/user-guide/tasks/manage-environments.html)  

1. 安裝gRPC for Python
  * pip install grpcio
1. 安裝gRPC tool，過程中會自動安裝相依package，也會連帶安裝protobuf
  * pip install grpcio-tools googleapis-common-protos
  * conda安裝方式可以去 [Anaconda](https://anaconda.org/)搜尋相對應作業系統的package，一樣是搜尋grpcio-tools



[GRPC官網](https://grpc.io/)  
GRPC 跨程式語言應用範例
