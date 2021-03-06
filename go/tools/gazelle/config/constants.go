/* Copyright 2017 The Bazel Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

const (
	// RulesGoRepoName is the canonical name of the rules_go repository. It must
	// match the workspace name in WORKSPACE.
	RulesGoRepoName = "io_bazel_rules_go"
	// RulesGoDefBzlLabel is the canonical label for def.bzl, where all
	// Go-related rules are defined.
	RulesGoDefBzlLabel = "@io_bazel_rules_go//go:def.bzl"
	// DefaultLibName is the name of the default go_library rule in a Go
	// package directory. It must be consistent to DEFAULT_LIB in go/private/common.bf.
	DefaultLibName = "go_default_library"
	// DefaultTestName is a name of an internal test corresponding to
	// DefaultLibName. It does not need to be consistent to something but it
	// just needs to be unique in the Bazel package
	DefaultTestName = "go_default_test"
	// DefaultXTestName is a name of an external test corresponding to
	// DefaultLibName.
	DefaultXTestName = "go_default_xtest"
	// DefaultProtosName is the name of a filegroup created
	// whenever the library contains .pb.go files
	DefaultProtosName = "go_default_library_protos"
	// DefaultCgoLibName is the name of the default cgo_library rule in a Go package directory.
	DefaultCgoLibName = "cgo_default_library"
)
