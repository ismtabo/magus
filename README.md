# MAGUS (MAgic Generation Utility for Source)

MAGUS is a tool for automatically generating magic source files based on a manifest.

## Installation

### Using Go
```bash
go install github.com/ismtabo/magus
```

## Usage

### Generate a new magic source file
```bash
magus generate [manifest] --dir output
```

**Arguments:**

- `manifest`: path to the manifest file.

**Flags:**

- `--dir`: directory where the magic source file will be generated. Defaults to the current directory.
- `--dry-run`: whether to perform a dry run or not. Defaults to `false`.
- `--overwrite,-w`: whether to overwrite existing files or not. Defaults to `false`.
- `--clean`: whether to clean the output directory before generating the magic source file or not. Defaults to `false`.
- `--var`: list of variables to be used in the manifest. This flag can be used multiple times. E.g: `--var name=foo --var version=1.0.0`.
- `--var-files`: list of files containing variables to be used in the manifest. This flag can be used multiple times. E.g: `--var-files ./vars.yaml --var-files ./vars.json`.

## Manifest

The manifest is a YAML file that describes the magic source file to be generated. It contains the following fields:

- `name`: Name of the magic source file to be generated.
- `version`: Version of the magic source file to be generated.
- `root`: Root directory of the magic source file to be generated.
- `variables`: List of variables to be used in the magic source file to be generated.
- `casts`: List of casts to be used in the magic source file to be generated.

### Variables

Variables are used to define the values of the magic source file to be generated. They are defined as a list of key-value pairs, where the key is the name of the variable and the value is the value of the variable.

Variables can be of three types:

- `literal`: The value of the variable is a literal value.
- `env`: The value of the variable is the value of an environment variable.
- `template`: The value of the variable is a template that will be rendered using the variables defined in the manifest.

### Casts

Casts are used to define the casts of the magic source file to be generated. They are defined as a record with the following fields:

- `to`: output directory relative to `root` where the cast will be generated.
- `from`: input source where the cast will be generated from.
- `variables`: list of variables to be used in the cast.

Additionally, the following fields can be used to generate casts conditionally:

- `if/unless`: condition to be evaluated to determine if the cast should be generated.

Furthermore, the following fields can be used to generate casts for each element of a list defined in a variable:

- `each`: list of casts to be generated for each element of the list defined in the variable. This field must render into a valid JSON array. To do so, the helper `to_json` can be used.
- `as`: alias of the `each` element in the cast.
- `include/omit`: condition to be evaluated to determine if the `each` element should be included in the cast.

### Sources

Sources are used to define the sources of the magic source file to be generated. Currently, only the `template` source is supported. It is defined as template string to be render using `text/template` with the variables defined in the manifest.

### Helpers

Helpers are used to define functions that can be used in the render process. These functions are provided by magus. Currently, the following helpers are available:

- `to_json`: converts a value into a JSON string.
- `lower`: converts a string into lowercase.
- `upper`: converts a string into uppercase.
- `title`: converts a string into title case.
- `snake`: converts a string into snake case.
- `kebab`: converts a string into kebab case.
- `camel`: converts a string into camel case.
- `pascal`: converts a string into pascal case.
- `constant`: converts a string into constant case.

### Example of manifest for a golang repository

```yaml
version: "1"
name: "go"
root: "."
casts:
    README.md:
        to: "./README.md"
        from: |
            # {{ .name | title }}
            {{ .description }}
            {{ .usage }}
    main.go:
        to: "./main.go"
        from: |
            package main

            import (
                "fmt"
            )

            func main() {
                fmt.Println("Hello, world!")
            }
    .gitignore:
        to: "./.gitignore"
        from: |
            {{ range .gitignore }}
            {{ . }}
            {{ end }}
```

### Example of React component with tests

Given the following variables, this manifest will generate a React component with tests and stories.

- `name`: Name of the component.
- `stories`: Whether to generate stories or not.


```yaml
version: "1"
name: "react"
root: "./{{ .name | kebab }}"
casts:
    index.js:
        to: "./index.jsx"
        from: |
            import React from 'react';
            import PropTypes from 'prop-types';

            const {{ .name | pascal }} = (props) => {
                return (
                    <div>
                        {{ .name | pascal }}
                    </div>
                );
            };

            {{ .name | pascal }}.propTypes = {
                name: PropTypes.string,
            };

            {{ .name | pascal }}.defaultProps = {
                name: '{{ .name | pascal }}',
            };

            export default {{ .name | pascal }};
    index.test.js:
        to: "./index.test.js"
        from: |
            import React from 'react';
            import { render } from '@testing-library/react';
            import {{ .name | pascal }} from './index';

            describe('{{ .name | pascal }}', () => {
                it('should render the component', () => {
                    const { container } = render(<{{ .name | pascal }} />);
                    expect(container).toMatchSnapshot();
                });
            });
    index.stories.js:
        to: "./index.stories.jsx"
        from: |
            import React from 'react';
            import {{ .name | pascal }} from './index';

            export default {
                title: '{{ .name | pascal }}',
                component: {{ .name | pascal }},
            };

            const Template = (args) => <{{ .name | pascal }} {...args} />;

            export const Default = Template.bind({});
            Default.args = {
                name: '{{ .name | pascal }}',
            };
        if: "{{ .stories }}"
```

## License

See [LICENSE](LICENSE) for more information.

## Future work

- [ ] Add support for more sources (e.g.: files or other magic manifest).
- [ ] Add support for more helpers (e.g.: shell scripts defined in the manifest).
- [ ] Add verification of manifest templates.
- [ ] Add support for more casts (e.g.: `copy`).
- [ ] Add support for spells (aka actions) (e.g.: `append`, `format`, `compile`, ...).

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for more information.

## Changelog

See [CHANGELOG](CHANGELOG.md) for more information.

## Versioning

This project uses [SemVer](https://semver.org/) for versioning. For the versions available, see [tags](https://github.com/ismtabo/magus/tags).

## Contributors

- [Ismael Taboada](github.com/ismtabo)
