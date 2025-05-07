package nodejs

import (
	"github.com/dop251/goja_nodejs/require"
)

func RegisterModules(req *require.Registry) {
	req.RegisterNativeModule("https", HTTPSModule)

}
