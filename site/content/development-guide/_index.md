---
title: Development Guide
breadcrumb: development
layout: page
---

<!-- Source: https://github.com/arduino/tooling-project-assets/blob/main/documentation-templates/contributor-guide/task/development.md -->

**inoplatforms** is a free open source project. The source code and registry is hosted in a public GitHub repository:
{{< github-url href="https://github.com/per1234/inoplatforms" >}}

## Prerequisites

The following development tools must be available in your local environment:

- [**Go**](https://go.dev/dl/) version 1.21.x - programming language, dependency manager
  - [**gvm**](https://github.com/moovweb/gvm#installing) is recommended if you want to manage multiple installations of **Go** on your system.
- [**Node.js**](https://nodejs.org/en/download) version 18.x - JavaScript runtime environment and **npm** package manager
  - [**nvm**](https://github.com/nvm-sh/nvm#installing-and-updating) is recommended if you want to manage multiple installations of **Node.js** on your system.
- [**Poetry**](https://python-poetry.org/docs/#installation) - Python package manager
- **POSIX-compliant shell**
  - Windows users should note that the **Git** for Windows installation includes the [**Git BASH**](https://gitforwindows.org/#bash) shell.
- [**Python**](https://www.python.org/downloads/) version 3.11.x - programming language

## Development Tasks

The project has a comprehensive infrastructure to assist with common development tasks. These operations are performed using the [**Task**](https://taskfile.dev/) task runner tool.

Run the following command to install **Task**:

```text
go install github.com/go-task/task/v3/cmd/task
```

Then this command for a list of the available tasks:

```text
task
```

## Website Preview

You can build and serve a local copy of the **inoplatforms** website from the source code by running these commands from the root folder of the project:

```text
go install github.com/go-task/task/v3/cmd/task
task website:serve
```

Once the build has finished, the URL of the locally served website will be shown in the terminal. Open that URL in your web browser to see the preview version of the site.

## Running Checks

Checks and tests are set up to ensure the project content is functional and compliant with the established standards.

You can run the checks by running this command from the root folder of the project:

```text
go install github.com/go-task/task/v3/cmd/task
task check
```

## Automatic Corrections

Tools are provided to automatically bring the project into compliance with some of the required checks.

You can make these automatic fixes by running this command from the root folder of the project:

```text
go install github.com/go-task/task/v3/cmd/task
task fix
```

## Additional Resources

- [**Contributor Guide**](/contributor-guide/)
- [**Acknowledgements**](https://github.com/per1234/inoplatforms/blob/main/docs/acknowledgments.md#acknowledgments)
