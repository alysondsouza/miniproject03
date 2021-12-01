module replicaNode

go 1.17

replace utils => ../utils

replace service => ../service

replace server => ../server

replace client => ../client

replace auction => ../auction

require (
	auction v0.0.0-00010101000000-000000000000
	client v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20211118161319-6a13c67c3ce4 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.42.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
