load("@gazelle//:def.bzl", "gazelle")
load("@rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/yashsoodini/bazelinit
gazelle(name = "gazelle")

go_library(
    name = "bazelinit_lib",
    srcs = ["main.go"],
    importpath = "github.com/yashsoodini/bazelinit",
    visibility = ["//visibility:private"],
    deps = ["//cmd"],
)

go_binary(
    name = "bazelinit",
    embed = [":bazelinit_lib"],
    visibility = ["//visibility:public"],
)
