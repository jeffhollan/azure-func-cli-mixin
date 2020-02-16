package functions

import (
	"get.porter.sh/porter/pkg/mixin"
	"get.porter.sh/porter/pkg/porter/version"
	"github.com/jeffhollan/azure-func-cli-mixin/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "functions",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Jeff Hollan",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
