---
sidebar_position: 6
description: "Template using Go templates"
---

# Templates

Templates are used all along the process, they are present in the variables values, in the cast properties (e.g. `to`, `from`, etc.). Templates are rendered using the [Go text/template engine](https://pkg.go.dev/text/template#pkg-overview).

:::caution

**Caveat:** Despite of other template engines such as handlebars, jinja or mustache, the Go template engine need to refer variables with a dot prefix. For example, to refer to the `project_name` variable, you need to use `.project_name`.

:::

Golang template engine allows to use if-else statements and loops. See more about control structures in the [text/template](https://pkg.go.dev/text/template#hdr-Actions) documentation. In addition, helper functions can be used in the templates to transform and manipulate variables. See more about helper functions in the [Helper Functions](./helpers.md) section.

For further information about the Go template engine, please refer to the [official documentation](https://pkg.go.dev/text/template).


## Example

```yaml
casts:
  - from: "{{ .project_name }}.cast"
    to: "{{ .project_name }}.cast"
    variables:
      - name: project_name
        value: My Project
```
