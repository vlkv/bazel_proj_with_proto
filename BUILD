
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:proto disable_global
# gazelle:prefix gitlab.com/myorg/myproj
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
