---
sidebar_position: 7
description: "Helper functions"
---

# Helper Functions

Magus extends the basic golang text/template engine by including [`slim-sprig`](https://masterminds.github.io/sprig/). This package provides a set of helper functions that can be used in the YAML manifest. Helper functions are used in the render of paths, templates, and conditions.

## Additional functions

In addition to the functions provided by `slim-sprig`, the following functions are available:

### String functions

The following functions are available to manipulate strings:

- `snake`: converts a string to snake case[^1]
- `constant`: converts a string to constant case[^1]
- `pascal`: converts a string to pascal case[^1]
- `camel`: converts a string to camel case[^1]
- `kebab`: converts a string to kebab case[^1]
- `pluralize`: converts a string to plural form[^2]
- `singularize`: converts a string to singular form[^2]

[^1]: Using [gostrcase](https://github.com/iancoleman/strcase)
[^2]: Using [go-pluralize](https://github.com/gertd/go-pluralize)
