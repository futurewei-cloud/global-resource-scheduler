load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "register.go",
        "types.go",
        "zz_generated.deepcopy.go",
    ],
    importpath = "k8s.io/kubernetes/pkg/controller/apis/config",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/controller/certificates/signer/config:go_default_library",
        "//pkg/controller/endpoint/config:go_default_library",
        "//pkg/controller/garbagecollector/config:go_default_library",
        "//pkg/controller/namespace/config:go_default_library",
        "//pkg/controller/podgc/config:go_default_library",
        "//pkg/controller/replicaset/config:go_default_library",
        "//pkg/controller/serviceaccount/config:go_default_library",
        "//pkg/controller/tenant/config:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//staging/src/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//staging/src/k8s.io/component-base/config:go_default_library",
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
        "//pkg/controller/apis/config/fuzzer:all-srcs",
        "//pkg/controller/apis/config/scheme:all-srcs",
        "//pkg/controller/apis/config/v1alpha1:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
