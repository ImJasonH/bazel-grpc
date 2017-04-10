package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary")

load("@org_pubref_rules_protobuf//go:rules.bzl", "GRPC_COMPILE_DEPS")

go_binary(
    name = "client",
    srcs = ["main.go"],
    deps = [
        "//proto",
    ] + GRPC_COMPILE_DEPS,
)
