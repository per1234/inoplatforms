# ino-hardware-package-list

[![Check General Formatting status](https://github.com/per1234/inoplatforms/actions/workflows/check-general-formatting-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-general-formatting-task.yml)
[![Check Go status](https://github.com/per1234/inoplatforms/actions/workflows/check-go-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-go-task.yml)
[![Check License status](https://github.com/per1234/inoplatforms/actions/workflows/check-license.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-license.yml)
[![Check Markdown status](https://github.com/per1234/inoplatforms/actions/workflows/check-markdown-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-markdown-task.yml)
[![Check npm status](https://github.com/per1234/inoplatforms/actions/workflows/check-npm-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-npm-task.yml)
[![Check Poetry status](https://github.com/per1234/inoplatforms/actions/workflows/check-poetry-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-poetry-task.yml)
[![Check Prettier Formatting status](https://github.com/per1234/inoplatforms/actions/workflows/check-prettier-formatting-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-prettier-formatting-task.yml)
[![Check Workflows status](https://github.com/per1234/inoplatforms/actions/workflows/check-workflows-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-workflows-task.yml)
[![Check YAML status](https://github.com/per1234/inoplatforms/actions/workflows/check-yaml-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/check-yaml-task.yml)
[![Spell Check status](https://github.com/per1234/inoplatforms/actions/workflows/spell-check-task.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/spell-check-task.yml)
[![Sync Labels status](https://github.com/per1234/inoplatforms/actions/workflows/sync-labels-npm.yml/badge.svg)](https://github.com/per1234/inoplatforms/actions/workflows/sync-labels-npm.yml)

A list of all known [Arduino](http://arduino.cc) hardware packages.

Hardware packages provide the [boards platform](https://arduino.github.io/arduino-cli/latest/platform-specification) (AKA "core") and toolchain needed to add support for a board to the Arduino development software (e.g., Arduino IDE).

The list can be viewed online [here](ino-hardware-package-list.tsv).

### Columns

- **Name**: The package index `packages[].platforms[].name` value, the platform.txt `name` property value, or an arbitrary name determined from looking at the repository content.
- **Vendor**: The package index `packages[].name` value or the name of the platform's vendor folder. The machine-friendly name of the package is `{vendor}:{architecture}`.
- **Architecture**: The package index `packages[].platforms[].architecture` value or the name of the platform's architecture folder. The machine-friendly name of the package is `{vendor}:{architecture}`.
- **Repository**: The website where the platform files are stored.
- **Boards Manager URL**: The URL for the [package index](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/) file that provides [Boards Manager](https://docs.arduino.cc/learn/starting-guide/cores) installation support. This URL must be added to the Arduino IDE's **File > Preferences > Additional Boards Manager URLs**.
- **Repository Data Folder**: The folder in the repository that contains [boards.txt](https://arduino.github.io/arduino-cli/latest/platform-specification/#boardstxt).
- **Branch Name**: The branch of the repository that contains the platform files.
- **Notes**: Additional information.

### Related

- https://github.com/arduino/Arduino/wiki/Unofficial-list-of-3rd-party-boards-support-urls
  - The "Unofficial list of 3rd party boards support urls" differs from ino-hardware-package-list in:
    - Less comprehensive due to consisting of submissions from boards platform authors.
    - Is not in a machine-readable form.

### Contributing

Additions/corrections/updates to the list are welcome! Please submit a pull request or issue.
