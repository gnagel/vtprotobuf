package main

import (
	"flag"
	"strings"

	_ "github.com/planetscale/vtprotobuf/features/clone"
	_ "github.com/planetscale/vtprotobuf/features/equal"
	_ "github.com/planetscale/vtprotobuf/features/grpc"
	_ "github.com/planetscale/vtprotobuf/features/marshal"
	_ "github.com/planetscale/vtprotobuf/features/pool"
	_ "github.com/planetscale/vtprotobuf/features/size"
	_ "github.com/planetscale/vtprotobuf/features/unmarshal"
	"github.com/planetscale/vtprotobuf/generator"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var cfg generator.Config
	var features string
	var f flag.FlagSet

	f.BoolVar(&cfg.AllowEmpty, "allow-empty", false, "allow generation of empty files")
	cfg.Poolable = make(generator.ObjectSet)
	f.Var(&cfg.Poolable, "pool", "use memory pooling for this object")
	f.BoolVar(&cfg.Wrap, "wrap", false, "generate wrapper types")
	f.BoolVar(&cfg.WellKnownTypes, "wkt", true, "generate optimized code for well-known types")
	f.StringVar(&features, "features", "all", "list of features to generate (separated by '+')")

	protogen.Options{ParamFunc: f.Set}.Run(func(plugin *protogen.Plugin) error {
		gen, err := generator.NewGenerator(plugin, strings.Split(features, "+"), &cfg)
		if err != nil {
			return err
		}
		gen.Generate()
		return nil
	})
}
