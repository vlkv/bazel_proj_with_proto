
Doc about go_proto_library rule: https://github.com/bazelbuild/rules_go/blob/master/proto/core.rst#go-proto-library

* Build all `bazelisk build //...`
* Test all `bazelisk test //...`
* Run main binary `bazelisk run //:bazel_proj_with_proto`


To run single test use `--test_filter` arg:
```
bazelisk run --test_filter 'TestOne' //:bazel_proj_with_proto_test
```
