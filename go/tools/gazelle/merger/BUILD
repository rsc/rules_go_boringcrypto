load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "fix.go",
        "merger.go"
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@io_bazel_rules_go//go/tools/gazelle/config:go_default_library",
        "@com_github_bazelbuild_buildtools//build:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "fix_test.go",
        "merger_test.go"
    ],
    library = ":go_default_library",
    deps = ["@com_github_bazelbuild_buildtools//build:go_default_library"],
    size = "small",
)
