---
title: Glossary
breadcrumb: glossary
layout: basic
page-type: static
---

## Arduino development software

Application that facilitates the development and deployment of Arduino firmware programs.

More information [**here** (Arduino CLI)](https://arduino.github.io/arduino-cli/latest/) and [**here** (Arduino IDE)](https://docs.arduino.cc/software/ide-v2).

## board

A target embedded system. The target could be a [development board](#development-board) or [microcontroller](#microcontroller) alone.

## Boards Manager

A component of Arduino development software for managing installation and updates of platforms.

More information [**here** (Arduino CLI)](https://arduino.github.io/arduino-cli/latest/commands/arduino-cli_core/) and [**here** (Arduino IDE)](https://docs.arduino.cc/software/ide-v2/tutorials/ide-v2-board-manager).

## core

Code that implements the [fundamental Arduino firmware API](https://www.arduino.cc/reference/en/) for a given target.

More information [**here**](https://arduino.github.io/arduino-cli/latest/platform-specification/#cores).

## development board

Hardware used as a target during development and prototyping of an embedded system device. This is a [PCB](https://wikipedia.org/wiki/Printed_circuit_board) with a [microcontroller](#microcontroller), essential support circuitry, and additional electronic components. Development boards are manufactured by the Arduino company as well as 3rd parties, or created DIY by users.

More information [**here**](https://docs.arduino.cc/).

## Git

Version control system.

More information [**here**](https://git-scm.com/).

## GitHub

Popular website that provides [Git](#git) [repository](#repository) hosting, enhanced with additional collaboration and social features.

More information [**here**](https://docs.github.com/en/get-started).

## host

The environment in which a [tool](#tool) will run.

## library

A collection of reusable firmware code.

More information [**here**](https://arduino.github.io/arduino-cli/latest/library-specification/).

## Library Manager

A component of Arduino development software for managing installation and updates of standalone Arduino libraries. Not directly relevant to Arduino platforms or this project, but sometimes confused with [**Boards Manager**](#boards-manager).

More information [**here** (Arduino CLI)](https://arduino.github.io/arduino-cli/latest/commands/arduino-cli_lib/) and [**here** (Arduino IDE)](https://docs.arduino.cc/software/ide-v2/tutorials/ide-v2-installing-a-library#installing-a-library).

## microcontroller

A small, inexpensive, power efficient computer packaged in an [integrated circuit](https://wikipedia.org/wiki/Integrated_circuit).

More information [**here**](https://wikipedia.org/wiki/Microcontroller).

## package

A collection of Arduino [platforms](#platform) and/or [tools](#tool).

## package index

Provides the data that allows a platform to be installed via [**Boards Manager**](#boards-manager).

More information [**here**](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/).

## package provider

Term used in the **inoplatforms** project to refer to any source of Arduino [packages](#package). This will ideally be a [package index](#package-index), but also applies to cases where a platform creator did not provide a package index.

## packager

The maintainer of an Arduino [package](#package).

---

**ⓘ** The term "package" is used exclusively in the **inoplatforms** project.

---

## platform

Provides everything needed to add support for a given target to the Arduino development software. This includes the configuration definitions that allow the [development software](#arduino-development-software) to generate the appropriate [tool](#tool) invocation commands, [cores](#core), and [libraries](#platform-bundled-library).

---

**ⓘ** Platforms are often [incorrectly](https://arduino.github.io/arduino-cli/latest/platform-specification/#platform-terminology) referred to as "cores". Although a platform may ([optionally](https://arduino.github.io/arduino-cli/latest/platform-specification/#core-reference)) include a [core](#core), the core is only one of the components of a platform.

---

More information [**here**](https://arduino.github.io/arduino-cli/latest/platform-specification/).

## platform bundled library

An [Arduino library](#library) that is included in the installation of a platform. A platform bundled library is only accessible when compiling for a board of its platform.

More information [**here**](https://arduino.github.io/arduino-cli/latest/platform-specification/#platform-bundled-libraries).

## pluggable discovery

A specific type of [tool](#tool) that provides Arduino development software with a list of ports of a given protocol that might be used for communication between the PC and an Arduino board.

More information [**here**](https://arduino.github.io/arduino-cli/latest/pluggable-discovery-specification/).

## pluggable monitor

A specific type of [tool](#tool) that provides the low level interface between Arduino development software and an Arduino board for arbitrary communications by the user. The user interface that wraps this communication in **Arduino IDE** is named [**Serial Monitor**](https://docs.arduino.cc/software/ide-v2/tutorials/ide-v2-serial-monitor) (which, despite the "serial" in the name, can be used for communication via any protocol thanks to the pluggable monitor system).

More information [**here**](https://arduino.github.io/arduino-cli/latest/pluggable-monitor-specification/).

## release

A versioned offering of a [platform](#platform) or [tool](#tool) at a specific point in its development history.

## repository

A project under version control.

More information [**here**](https://wikipedia.org/wiki/Repository_%28version_control%29).

## tool

An application used by the Arduino development software to perform a given process (e.g., compilation, upload) for a given target board.

More information [**here**](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/#tools-definitions).

## tool dependency

The specification of a platform's dependency on a [tool](#tool).

More information [**here**](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/#platforms-definitions).

## vendor

An alternative term for "[packager](#packager)".

---

**ⓘ** The term "package" is used exclusively in the **inoplatforms** project.

---
