workspace(name = "bazel_grpc")

# ================================================================
# go rules
# ================================================================

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "6fab60252e89cb603afce15d0d0321758895ffd2",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")

go_repositories()

# ================================================================
# protobuf rules
# ================================================================

git_repository(
    name = "org_pubref_rules_protobuf",
    remote = "https://github.com/pubref/rules_protobuf.git",
    commit = "d6f96942fefc5abf459a7942dd92cf3b54eca983",
)

load("@org_pubref_rules_protobuf//protobuf:rules.bzl", "proto_repositories")

proto_repositories()

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")

go_proto_repositories()

# ================================================================
# docker dependencies
# ================================================================

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    commit = "06bb5b29088e55a29d09e37e1c88f1cc44aaa806",
)

load(
   "@io_bazel_rules_docker//docker:docker.bzl",
   "docker_pull",
   "docker_repositories",
)

# Consumers shouldn't need to do this themselves once WORKSPACE is
# instantiated recursively.
docker_repositories()

docker_pull(
    name = "base_image",
    registry = "index.docker.io",
    repository = "library/debian",
    digest = "sha256:d3892ebd85bb21b161e9d170beefd6c444d9bd58aba3998e1f67aa842fbedb9d", # wheezy
)
