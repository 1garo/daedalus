load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/1garo/daedalus
gazelle(name = "gazelle")

go_library(
    name = "daedalus_lib",
    srcs = ["main.go"],
    importpath = "github.com/1garo/daedalus",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "daedalus",
    embed = [":daedalus_lib"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "daedalus_test",
    srcs = ["main_test.go"],
    embed = [":daedalus_lib"],
)
