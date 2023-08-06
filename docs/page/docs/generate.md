---
sidebar_position: 3
description: "Generate command"
---

# Generate command

## Overview

The `generate` command is used to generate a new project from a manifest.

## Usage

```bash
magus generate [manifest] [flags]
```

**Arguments:**

- `manifest`: path to the manifest file.

**Flags:**

- `--dir <output_dir>`: directory where the magic source file will be generated. Defaults to the current directory (`.`).
- `--dry-run`: whether to perform a dry run or not. Defaults to `false`.
- `--overwrite,-w`: whether to overwrite existing files or not. Defaults to `false`.
- `--clean`: whether to clean the output directory before generating the magic source file or not. Defaults to `false`.
- `--var`: list of variables to be used in the manifest. This flag can be used multiple times. E.g: `--var name=foo --var version=1.0.0,build=bcdd1a0`.
- `--var-file`: list of files containing variables to be used in the manifest. This flag can be used multiple times. E.g: `--var-file vars.yaml --var-file ./foo.yaml,./bar.json`.
