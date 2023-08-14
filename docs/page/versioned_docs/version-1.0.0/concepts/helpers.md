---
sidebar_position: 7
description: "Helper functions"
---

# Helper Functions

Magus extends the basic golang text/template engine by including [`slim-sprig`](https://go-task.github.io/slim-sprig). This package provides a set of helper functions that can be used in the YAML manifest. Helper functions are used in the render of paths, templates, and conditions.

## Additional functions

In addition to the functions provided by `slim-sprig`, the following functions are available:

### String functions

The following functions are available to manipulate strings:

- `snake`: converts a string to snake case
- `constant`: converts a string to constant case
- `pascal`: converts a string to pascal case
- `camel`: converts a string to camel case
- `kebab`: converts a string to kebab case
