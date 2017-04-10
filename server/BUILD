package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "server",
    srcs = ["main.go"],
    deps = [
        "//proto",
        "@google_golang_org_grpc//:go_default_library",
    ],
)

load("@bazel_tools//tools/build_defs/docker:docker.bzl", "docker_build")

docker_build(
    name = "image",
    base = "@base_image//image:image.tar",
    entrypoint = ["/server"],
    files = [":server"],
    ports = ["50051"],
)
