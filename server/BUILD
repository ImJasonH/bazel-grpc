package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary")

load("@org_pubref_rules_protobuf//go:rules.bzl", "GRPC_COMPILE_DEPS")

go_binary(
    name = "server",
    srcs = ["main.go"],
    deps = [
        "//proto",
    ] + GRPC_COMPILE_DEPS,
 )

load("@bazel_tools//tools/build_defs/docker:docker.bzl", "docker_build")

docker_build(
    name = "image",
    base = "@base_image//image:image.tar",
    entrypoint = ["/server"],
    files = [":server"],
    ports = ["50051"],
)
