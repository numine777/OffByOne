load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//go:def.bzl", "nogo")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

# gazelle:prefix github.com/numine777/OffByOne
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

nogo(
    name = "my_vet",
    vet = True,
    visibility = ["//visibility:public"],
)

go_library(
    name = "OffByOne_lib",
    srcs = [
        "api.go",
        "main.go",
    ],
    importpath = "github.com/numine777/OffByOne",
    visibility = ["//visibility:private"],
    deps = [
        "//core",
        "@io_gorm_driver_sqlite//:sqlite",
        "@io_gorm_gorm//:gorm",
    ],
)

go_binary(
    name = "OffByOne",
    embed = [":OffByOne_lib"],
    visibility = ["//visibility:public"],
)

go_image(
    name = "app_image",
    binary = ":OffByOne",
)
