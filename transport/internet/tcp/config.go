package tcp

import (
	"github.com/F4RD1N/xray-core/common"
	"github.com/F4RD1N/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
