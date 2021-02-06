# ino-hardware-package-list

A list of all known [Arduino](http://arduino.cc) hardware packages.

Hardware packages provide the [boards platform](https://arduino.github.io/arduino-cli/latest/platform-specification) (AKA "core") and toolchain needed to add support for a board to the Arduino development software (e.g., Arduino IDE).

### Columns

- **Name**: The platform name from the boards manager JSON file, the `name` property from platform.txt, or an arbitrary name determined from looking at the repository content.
- **Vendor**: The name of the package's vendor folder. The machine-friendly name of the package is `{vendor}:{architecture}`.
- **Architecture**: The name of the package's architecture folder. The machine-friendly name of the package is `{vendor}:{architecture}`.
- **Repository**: The website where the package files are stored.
- **Boards Manager URL**: The URL for the JSON file that provides Boards Manager installation support. This URL must be added to the Arduino IDE's **File > Preferences > Additional Boards Manager URLs**.
- **Repository Data Folder**: The folder in the repository that contains boards.txt.
- **Branch Name**: The branch of the repository that contains the package files.
- **Notes**: Additional information.

### Contributing

Additions/corrections/updates to the list are welcome! Please submit a pull request or issue.
