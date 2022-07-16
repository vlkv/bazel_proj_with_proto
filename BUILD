load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:proto disable_global
# gazelle:prefix github.com/vlkv/bazel_proj_with_proto
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "myproj_lib",
    srcs = ["main.go"],
    importpath = "github.com/vlkv/bazel_proj_with_proto",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_golang_protobuf//proto",
        "//proto:myproto_go_proto",
    ],

)

go_binary(
    name = "myproj",
    embed = [":myproj_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "myproj_test",
    srcs = ["main_test.go"],
    embed = [":myproj_lib"],
    deps = ["@com_github_stretchr_testify//assert"],
)
