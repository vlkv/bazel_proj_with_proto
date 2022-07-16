load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

# This is an API workspace, having public visibility by default makes perfect sense.
package(default_visibility = ["//visibility:public"])

proto_library(
    name = "myproto_proto",
    srcs = [
        "myproto.proto",
        "myproto_api.proto",
    ],
    deps = [],
)

go_proto_library(
    name = "myproto_go_proto",
    importpath = "github.com/vlkv/bazel_proj_with_proto/proto",
    protos = [":myproto_proto"],
)
