module github.com/temporalio/ringpop-go

go 1.11

require (
	github.com/apache/thrift v0.15.0
	github.com/benbjohnson/clock v0.0.0-20160125162948-a620c1cc9866
	github.com/cactus/go-statsd-client/statsd v0.0.0-20200423205355-cb0885a1018c
	github.com/dgryski/go-farm v0.0.0-20140601200337-fc41e106ee0e
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20141108142129-dee209f2455f
	github.com/samuel/go-thrift v0.0.0-20210915161234-7b67f98e972f // indirect
	github.com/sirupsen/logrus v1.0.2-0.20170726183946-abee6f9b0679
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/uber-common/bark v1.0.0
	github.com/temporalio/tchannel-go // TODO: use gomod version for tchannel-go with Thrift v0.15.0
	github.com/vektra/mockery v0.0.0-20160406211542-130a05e1b51a // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
	golang.org/x/tools v0.1.8 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
