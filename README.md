# go-graphql-tutorial

Just an example for `Graph Query language`.

## GQL Stacks
- gqlgen
- gin

## Federation?
- The idea is distributed graphs
- Building a new unified and single Graph API (by combining multiple graph schemas)
- With GraphQL federation, you tell the gateway where it needs to look for the different objects and what URLs they live at. The subgraphs provide metadata that the gateway uses to automatically stitch everything together. This is a low-maintenance approach that gives your team a lot of flexibility.
- Have a gateway layer
- Split or decentralized yet interoperable among teams working in different locations
- Schemas are split across multiple services, and as they grow or scale over time, federation via the “gateway” layer

### Develop Guides
- Learning about schema https://graphql.org/learn/schema/
- https://gqlgen.com/getting-started/
- https://gqlgen.com/recipes/federation/

**References**

- [https://blog.logrocket.com/the-what-when-why-and-how-of-federated-graphql/](https://blog.logrocket.com/the-what-when-why-and-how-of-federated-graphql/)
- [https://www.apollographql.com/docs/federation/](https://www.apollographql.com/docs/federation/)

## Getting Started

### cmd/gql
GraphQL

1. setup module
```shell
go mod init <your_module>
```
2. Initialise gqlgen config and generate models
```shell
go run github.com/99designs/gqlgen init
```

The project skeleton will be created. you can read more about skeleton structure and file from [here](go run github.com/99designs/gqlgen init)

3. Removed unnecessary code from generated file. Here you are going to leave `gqlgen.yml` on your project.
```shell
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
```
4. Modify the `schema` which are includes schema, mutation and query.
5. Run `go run github.com/99designs/gqlgen generate` to re-generate the schema
6. Make sure all schema has been updated by `git diff`

### cmd/gql-server
GraphlQL Server and Playground

```shell
go run main.go
```

## Playground Queries

**Mutation**
```shell
mutation createTodo {
  createTodo(input: { text: "golf todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}
```

**Query**
```shell
query findTodos {
  todos {
    user {
      id
      name
    }
    text
  }
}
```