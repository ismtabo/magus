---
sidebar_position: 6
description: "Render a cast multiple times"
---

# Collection Casts

You can render a cast multiple times using the `each` property. The `each` property is a [Go template](https://pkg.go.dev/text/template) that is evaluated with the cast variables. The result of the template **MUST** be a string representing a **JSON array**. Each element in the list is used to render the cast, and its value is passed to the context of the render as the `It` variable named. For example, the following cast renders a list of users:

```yaml
casts:
  users:
    to: ./users-{{ .It.name }}.txt
    from: |
      {{ .It.name }} is {{ .It.age }} years old.
    each: |
      [
        { "name": "John", "age": 30 },
        { "name": "Jane", "age": 25 }
      ]
```

## Item Alias

You can override the default `It` variable name using the `as` property.

```yaml
casts:
  users:
    to: ./users-{{ .user.name }}.txt
    from: |
      {{ .user.name }} is {{ .user.age }} years old.
    each: |
      [
        { "name": "John", "age": 30 },
        { "name": "Jane", "age": 25 }
      ]
    as: user
```

## Iteration properties

In addition to the item value, the following properties are available in the `each` render context:

- `Index`: The index of the item in the list.
- `First`: `true` if the item is the first in the list, `false` otherwise.
- `Last`: `true` if the item is the last in the list, `false` otherwise.

For example, the following cast renders a list of users:

```yaml
casts:
  users:
    to: ./users-{{ .Index }}.txt
    from: |
      {{ .It.name }} is {{ .It.age }} years old.
    each: |
      [
        { "name": "John", "age": 30 },
        { "name": "Jane", "age": 25 }
      ]
```

## Conditional Item Rendering

You can use the `include` or `omit` properties to conditionally render an item. The `include` property is a [Go template](https://pkg.go.dev/text/template) that is evaluated with the cast variables. The result of the template should be a boolean value. If the result is `true`, the item is rendered. If the result is `false`, the item is not rendered. The `omit` property is the opposite of the `include` property. If the result of the `omit` template is `true`, the item is **not** rendered.

:::caution

The `include` and `omit` properties are mutually exclusive. If both properties are specified, the `include` property is used.

:::

:::tip

The iteration properties (`Index`, `First`, and `Last`) are available in the `include` and `omit` render context.

:::

For example, the following cast renders a list of users:

```yaml
casts:
  users:
    to: ./users
    from: |
      {{ .It.name }} is {{ .It.age }} years old.
    each: |
      [
        { "name": "John", "age": 30 },
        { "name": "Jane", "age": 25 }
      ]
    include: |
      {{ .It.age > 25 }}
```

## Combination with `if` and `unless`

You can combine the `each` property with the `if` and `unless` properties. These last properties affects to the whole cast, not to each item. For example, the following cast wont be rendered if the `renderUsers` variable is `false`:

```yaml
variables:
  - name: renderUsers
    value: false
casts:
  users:
    to: ./users
    from: |
      {{ .It.name }} is {{ .It.age }} years old.
    each: |
      [
        { "name": "John", "age": 30 },
        { "name": "Jane", "age": 25 }
      ]
    if: |
      {{ .renderUsers }}
```
