workspace(name = "com_github_off_by_one")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    url = "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# Declare Go direct dependencies.
go_repository(
    name = "io_gorm_driver_sqlite",
    importpath = "gorm.io/driver/sqlite",
    sum = "h1:SStaH/b+280M7C8vXeZLz/zo9cLQmIGwwj3cSj7p6l4=",
    version = "v1.2.6",
)

go_repository(
    name = "io_gorm_gorm",
    importpath = "gorm.io/gorm",
    sum = "h1:lYREBgc02Be/5lSCTuysZZDb6ffL2qrat6fg9CFbvXU=",
    version = "v1.22.5",
)

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()
