---
sidebar_position: 1
description: "Describe what cast should be generated and how"
---

# Manifest

Manifest represents a collection of source files that should be generated for either a project, tool configuration, project component, or other similar use cases. Manifest is a YAML file that describes the cast generation process.

## Manifest Structure

Manifest is a YAML file that describes the cast generation process. It consists of the following sections:

- `version` - version of the manifest
- `name` - name of the manifest
- `root` - root where the cast should be generated
- `variables` - variables that can be used in the manifest
- `casts` - list of casts that should be generated

Both `version` and `name` are required fields, for debugging purposes it is recommended to provide a name for the manifest.

### Root

Root is a path where the cast should be generated. It can be either absolute or relative to the manifest file. The root maybe later overridden by the user using the `--dir` flag.


:::caution

**Caveat:** this property is **not** a template.

:::

### Variables

Variables are used to provide data for cast generation. They can be used in the casts. Variables are defined as a list of variable objects. The value can be either a string or a list of strings. Further information about variables can be found in the [Variables](./variables.md) section.

### Casts

Casts are used to describe what should be generated. Casts are defined as a map of cast objects, they represent the source of the template to generate and where to generate the cast. They have additional feature such as conditional rendering and rendering over collections. Further information about casts can be found in the [Cast](./casts.md) section.

:::caution

**Caveat:** Magus will not apply the manifest if there is a conflict between destinations of multiples casts. For example, if you have two casts that generate the same file, Magus will show an error warning about the conflict. In case of existing files, Magus can override them using the `--clean` or `--overwrite` flags. See more at the [Generate command](/docs/generate) page.

:::

## Example

```yaml
version: 1
name: My Project
root: ./generated
variables:
  - name: project_name
    value: My Project
casts:
  hello-world:
    to: ./hello-world
    from: |
      Hello {{ .project_name }}!
    variables:
      - name: project_name
        value: My Project
```

### Example of manifest for a golang repository

Given the following variables, this manifest will generate a React component with tests and stories.

- `name`: Name of the component.
- `description`: Description of the component.
- `usage`: Usage of the component.
- `gitignore`: List of gitignore rules.

```yaml
version: "2"
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
      {{ .gitignore }}
```

### Example of React component with tests

Given the following variables, this manifest will generate a React component with tests and stories.

- `name`: Name of the component.
- `stories`: Whether to generate stories or not.

```yaml
version: "2"
name: "react"
root: "."
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
    if: "{{ .stories }}"
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
```
