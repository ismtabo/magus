---
sidebar_position: 3
description: "Provide data for cast generation"
---

# Variables

Variables are used to provide data for cast generation. Variables are defined as a list of variable objects. Variables can be of three different types: `literals`, `environment`, and `templates`. Each variable object has a name and extra value property depending on the type of the variable.

```yaml
variables:
  # Literal variable
  - name: project_name
    value: My Project
  # Environment variable
  - name: project_name
    env: PROJECT_NAME
  # Template variable
  - name: project_name
    template: |
      {{ .project_name }}
```

## Literals

Literals are the simplest type of variables. Literals are defined by providing a `value` property. The value can be of any type. The value is used as is in the cast generation process.

```yaml
variables:
  - name: project_name
    value: My Project
```

## Environment

Environment variables are defined by providing an `env` property. The value of the variable is taken from the environment variable with the name provided.

```yaml
variables:
  - name: project_name
    env: PROJECT_NAME
```

## Templates

Template variables are defined by providing a `template` property. The value of the variable is taken from the template provided. The template is rendered using the [Go template engine](https://pkg.go.dev/text/template).

```yaml
variables:
  - name: project_name
    template: |
      {{ .project_name }}
```

## Input variables

Input variables are variables that are provided by the user using the `--var` flag. They are not defined in the manifest. They can be used in the manifest as any other variable. They are not accessible from other casts.

```yaml
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
```

In this example, the `project_name` variable is not defined in the manifest. It is provided by the user using the `--var` flag.

### Input file variables

Input file variables are variables that are provided by the user using the `--var-file` flag. They are not defined in the manifest. They can be used in the manifest as any other variable. They are not accessible from other casts.

```yaml
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
```

In this example, the `project_name` variable is not defined in the manifest. It is provided by the user using the `--var-file` flag.

```
$ cast manifest.yaml --var-file vars.yaml
```

The `vars.yaml` file should contain the following content:

```yaml
project_name: My Project
```

## Variables evaluation

Variables are evaluated in the following order of declaration (from top to bottom). Variables can be overridden by the user using the `--var` flag. Take the following example:

```yaml
variables:
  - name: project_name
    value: My Project
  - name: project_name
    env: PROJECT_NAME
  - name: project_name
    template: |
      {{ .project_name }}
```

In this example, the `project_name` variable is defined three times. The first time it is defined as a literal, the second time as an environment variable, and the third time as a template. The value of the variable will be overwrite at each re-declaration. E.g: later the first declaration as a literal, the environment variable `PROJECT_NAME` will overwrite the value of the variable even if it is empty. Finally, the value of the variable will be taken from the render of the template.

If the user provides a value for the variable using the `--var` flag, the value will be taken from the user input.

## Scopes

Variables can be defined in different scopes (at root level and for each of the casts). Casts variables will overwrite the root variables. And they are not accessible from other casts.

```yaml
variables:
  - name: project_name
    value: My Project
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
  hello-world-2:
    to: ./hello-world-2
    from: |
      Hello {{ .project_name }}!
    variables:
      - name: project_name
        value: My Project 2
```

In this example, the `project_name` variable is defined at root level. It is used in the `hello-world` cast. And it is overwritten in the `hello-world-2` cast.

## Special variables

There are two special variables: `conditions` and `collections`. They are used to provide conditional rendering and rendering over collections. Will be explained in depth later in [Conditional Rendering](./conditional.md) and [Rendering over Collections](./collection.md) sections.

To summarize, `condition` variables will be interpreted as a boolean value. If the value is `true` the cast will be rendered, otherwise it will be considered false. `collections` variables will be interpreted as a list of values and need to be rendered as valid JSON Array strings. So, they can later be used in the cast to render over the collection. The `to_json` function can be used to render a valid JSON Array string from a list of values.

```yaml
variables:
  - name: condition
    value: true
  - name: collection
    value: ["a", "b", "c"]
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
    if: "{{ .condition }}"
    each: "{{ .collection }}"
    omit: "{{ .cast-condition | to_json }}"
    variables:
      - name: cast-condition
        value: true
```
