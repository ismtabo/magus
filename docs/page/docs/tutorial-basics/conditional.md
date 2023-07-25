---
sidebar_position: 4
description: "Conditionally render a cast"
---

# Conditional Rendering

You can conditionally render a cast using the `if` and `unless` properties. The `is` property is a [Go template](https://pkg.go.dev/text/template) that is evaluated with the cast variables. If the result is `true`, the cast is rendered, otherwise it is skipped. On the other hand the `unless` property is the opposite of `if`, if the result is `true`, the cast is **not** rendered.

## Example

```yaml
casts:
  - from: ./hello-world
    to: ./hello-world
    if: "{{ .project.enabled }}"
    variables:
      - name: project_name
        value: My Project
```
