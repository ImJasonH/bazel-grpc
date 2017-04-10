workspace(name = "bazel_grpc")

# ================================================================
# go rules
# ================================================================

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "6fab60252e89cb603afce15d0d0321758895ffd2",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "new_go_repository")

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

# ====
# external dependencies
# ====

new_go_repository(
    name = "google_golang_org_grpc",
    importpath = "google.golang.org/grpc",
    remote = "https://github.com/grpc/grpc-go",
    tag = "v1.2.1",
    vcs = "git",
)

# TODO: Do I really need a new_go_repository for each external dependency? I
# just want to use reflection from the grpc package.
new_go_repository(
    name = "google_golang_org_grpc_reflection",
    importpath = "google.golang.org/grpc/reflection",
    remote = "https://github.com/grpc/grpc-go",
    tag = "v1.2.1",
    vcs = "git",
)
