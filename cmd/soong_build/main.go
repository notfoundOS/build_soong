// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/blueprint/bootstrap"

	"android/soong"

	// These imports cause the modules to register their ModuleTypes, etc. with the soong package
	_ "android/soong/art"
	_ "android/soong/cc"
	"android/soong/common"
	_ "android/soong/genrule"
	_ "android/soong/java"
)

func main() {
	flag.Parse()

	// The top-level Blueprints file is passed as the first argument.
	srcDir := filepath.Dir(flag.Arg(0))

	ctx := soong.NewContext()

	configuration, err := common.NewConfig(srcDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	// Temporary hack
	//ctx.SetIgnoreUnknownModuleTypes(true)

	bootstrap.Main(ctx, configuration, common.ConfigFileName)
}