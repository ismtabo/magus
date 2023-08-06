---
sidebar_position: 6
description: "Template using Go templates"
---

# Templates

Templates are used all along the process, they are present in the variables values, in the path properties (e.g. `root` or `to`), and in the cast `from`. Templates are rendered using the [Go template engine](https://pkg.go.dev/text/template).

:::caution

**Caveat:** Despite of other template engines such as handlebars, jinja or mustache, the Go template engine need to refer variables with a dot prefix. For example, to refer to the `project_name` variable, you need to use `.project_name`.

:::

For further information about the Go template engine, please refer to the [official documentation](https://pkg.go.dev/text/template).


## Example

```yaml
root: "{{ .project_name }}"

casts:
  - from: "{{ .project_name }}.cast"
    to: "{{ .project_name }}.cast"
    variables:
      - name: project_name
        value: My Project
```
