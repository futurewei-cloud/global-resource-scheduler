load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "types.go",
    ],
    importpath = "k8s.io/kubernetes/globalscheduler/pkg/apis/scheduler/client",
    visibility = ["//visibility:public"],
    deps = [
        "//globalscheduler/pkg/apis/scheduler/client/clientset/versioned:go_default_library",
        "//globalscheduler/pkg/apis/scheduler/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//globalscheduler/pkg/apis/scheduler/client/clientset/versioned:all-srcs",
        "//globalscheduler/pkg/apis/scheduler/client/informers/externalversions:all-srcs",
        "//globalscheduler/pkg/apis/scheduler/client/listers/scheduler/v1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
