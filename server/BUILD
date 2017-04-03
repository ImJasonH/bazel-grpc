package(default_visibility = ["//visibility:public"])

load("@bazel_tools//tools/build_defs/docker:docker.bzl", "docker_build")

go_binary(
    name = "server",
    srcs = ["main.go"],
    deps = ["//proto"],
)

docker_build(
    name = "image",
    entrypoint = ["/server"],
    files = [":server"],
    ports = ["50051"],
)
