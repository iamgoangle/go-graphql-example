# go-graphql-tutorial

Just an example for `Graph Query language`.

## GQL Stacks
- gqlgen is a library for creating GraphQL applications in Go.

## Federation?
- https://www.apollographql.com/blog/announcement/apollo-federation-f260cf525d21/
- 

## Getting Started

1. setup module
```shell
go mod init <your_module>
```
2. Add `tools.go`
```shell
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

go mod tidy
```
3. Initialise gqlgen config and generate models
```shell
go run github.com/99designs/gqlgen init
```

The project skeleton will be created. you can read more about skeleton structure and file from [here](go run github.com/99designs/gqlgen init)

4. Removed unnecessary code from generated file. Here you are going to leave `gqlgen.yml` on your project.
```shell
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
```