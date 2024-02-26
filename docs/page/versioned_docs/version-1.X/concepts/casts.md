---
sidebar_position: 2
description: "Give instructions on how to generate a file(s)"
---

# Cast

Cast is a template that should be generated. It consists of the following sections:

- `to` - path where the cast should be generated
- `from` - source of the cast
- `variables` - variables that can be used in the cast

Additionally, cast can have the following sections:

- `if`, `unless` - condition(s) that should (or not) be met to render the cast
- `each` - collection that should be used to render the cast multiple times

And other sections that will be later mentioned in [Conditional Rendering](/docs/concepts/conditional) and [Collection Casts](/docs/concepts/collection) sections.

## Cast destination (to)

Destination is a path where the cast should be generated. It **must** be relative to the `root` property. The destination may be templated using the [Variables](/docs/concepts/variables) section at root level and from the cast context.

## Cast source (from)

Source is a template that should be generated. The source will be rendered using the [Go templates](https://pkg.go.dev/text/template) package. The source may be templated using the [Variables](/docs/concepts/variables) section at root level and from the cast context.

## Cast variables

Variables are used to provide data for cast generation. They can be used in the `to` and `from` fields. Variables are defined as a list of variable objects. The value can be either a string or a list of strings. Further information about variables can be found in the [Variables](/docs/concepts/variables) section.

## Example

```yaml
version: 1
name: My Project
root: ./generated
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
    variables:
      - name: project_name
        value: My Project
```

