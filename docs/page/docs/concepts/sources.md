---
sidebar_position: 4
description: "Sources of casts"
---

# Sources

Sources are used to provide the cast content to be rendered. In [v1](../../versioned_docs/version-1.0.0/concepts/casts.md#cast-source-from), can be only defined by using templated string for the `from` property of casts. In this version, we include other magic manifests as sources.

## Template sources

When the `from` property of a cast contains a templated string, the template is rendered using the [Go template engine](https://pkg.go.dev/text/template). See [Templates](./templates.md) for further information.

Example:

```yaml
casts:
  - from: |
      Hello {{ .project_name }}!
    to: ./hello-world
    variables:
      - name: project_name
        value: My Project
```

## Magic sources

Magic sources are manifests that are included in the current manifest. They are defined by providing a `magic` property. The value of the property is the path to the manifest to include. The path is relative to the current manifest.

All the variables defined in the current manifest are available in the magic source.

Example:

```yaml
---
# manifest.yaml
version: 1
name: My Project
root: ./generated
casts:
  - from:
        magic: ./hello-world.yaml
    to: ./hello-world
    variables:
      - name: project_name
        value: My Project
---
# hello-world.yaml
version: 1
name: Hello World
root: ./hello-world
casts:
  - from: |
      Hello {{ .project_name }}!
    to: ./hello-world
```

:::caution

**Caveat:** The magic source is rendered in the same context as the current manifest. This means that the `root` property of the magic source is overwritten by the `to` property of the magic cast.

:::

:::danger

**Danger:** Including a manifest in itself or referencing a parent manifest in any of its descendants -referenced magic sources- will result in an infinite loop. To avoid this, Magus will throw an error if it detects a loop.

:::
