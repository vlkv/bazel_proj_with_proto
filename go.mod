module gitlab.com/myorg/myproj

go 1.18

require gitlab.com/myorg/myproj/proto v0.0.0

require (
	github.com/golang/protobuf v1.5.2
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace gitlab.com/myorg/myproj/proto v0.0.0 => ./proto/gitlab.com/myorg/myproj/proto
