package all

import (
	"github.com/F4RD1N/xray-core/main/commands/all/api"
	"github.com/F4RD1N/xray-core/main/commands/all/convert"
	"github.com/F4RD1N/xray-core/main/commands/all/tls"
	"github.com/F4RD1N/xray-core/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
