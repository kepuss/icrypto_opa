package main

import (
	"fmt"
	"context"
	//"github.com/open-policy-agent/opa/storage/inmem"
	//"github.com/open-policy-agent/opa/plugins"
	//"github.com/open-policy-agent/opa/server"
	//"github.com/open-policy-agent/opa/runtime"
	"os"
	"github.com/open-policy-agent/opa/util"
	//"github.com/open-policy-agent/opa/runtime"
	"github.com/open-policy-agent/opa/runtime"

)

func main() {
	logLevel := util.NewEnumFlag("info", []string{"debug", "info", "error"})
	logFormat := util.NewEnumFlag("json", []string{"text", "json", "json-pretty"})
	ctx := context.Background()
	params := icrypto_runtime.NewParams()
	params.Addrs = &[]string{":8182"}
	params.Logging = runtime.LoggingConfig{
		Level:  logLevel.String(),
		Format: logFormat.String(),
	}

	rt, err := runtime.NewRuntime(ctx, params)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	rt.StartServer(ctx)
}
}