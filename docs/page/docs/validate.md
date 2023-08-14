---
sidebar_position: 4
description: "Validate command"
---

# Validate command

## Overview

The `validate` command is used to validate a manifest. The validation does following checks:

- Checks if the manifest has valid schema.
- Checks if the manifest version is supported.
- Checks if the manifest does not contains cycles in dependencies. See more at [Magic Sources](./concepts/sources.md#magic-sources).

## Usage

```bash
magus validate [manifest]
```

**Arguments:**

- `manifest`: path to the manifest file.
