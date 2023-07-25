---
sidebar_position: 7
description: "Helper functions"
---

# Helper Functions

Magus provides a set of helper functions that can be used in the YAML manifest. Helper functions are used in the render of paths, templates, and conditions.

## to_json

The `to_json` function converts a string to a JSON object. For example, the following cast renders a list of users. For example the following template renders a list of users:

```
{{ to_json .It }}
```

With the following input:

```
- name: John
  age: 30
- name: Jane
  age: 25
```

The output will be:

```
[
  {
    "name": "John",
    "age": 30
  },
  {
    "name": "Jane",
    "age": 25
  }
]
```

## String functions

The following functions are available to manipulate strings:

- `lower`: converts a string to lowercase
- `upper`: converts a string to uppercase
- `snake`: converts a string to snake case
- `constant`: converts a string to constant case
- `pascal`: converts a string to pascal case
- `camel`: converts a string to camel case
- `kebab`: converts a string to kebab case
- `title`: converts a string to title case
