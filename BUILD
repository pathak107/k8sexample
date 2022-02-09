load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/pathak107/k8Example
gazelle(name = "gazelle")

go_binary(
    name = "rtsp-simple-server",
    embed = [":k8Example_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "k8Example_lib",
    srcs = ["main.go"],
    importpath = "github.com/pathak107/k8Example",
    visibility = ["//visibility:private"],
    deps = [
        "//grpcserverservice",
        "//k8grpc",
        "//k8service",
        "//restServices",
        "@org_golang_google_grpc//:go_default_library",
    ],
)
