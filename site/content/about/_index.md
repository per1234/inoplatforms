---
title: About
breadcrumb: about
layout: basic
page-type: static
---

**inoplatforms** is a catalog of all known Arduino boards platforms.

## About Arduino

The term "[**Arduino**](https://www.arduino.cc/)" is used when referring to various things; a company, a community, development hardware, a firmware framework. At the heart of all these things is the idea of making [embedded systems](https://wikipedia.org/wiki/Embedded_system) accessible; offering everyone the opportunity to create things with [microcontrollers](https://wikipedia.org/wiki/Microcontroller).

## About Arduino Hardware

Arduino projects may created from a completely custom circuit built from base components, or built around a manufactured [development board](/glossary/#development-board). A diverse array of microcontrollers and development boards may be used.

The microcontrollers range from the relatively limited (e.g., [ATtiny13](https://www.microchip.com/product/attiny13)) to high performance cutting edge technology (e.g., [STM32H7](https://wikipedia.org/wiki/STM32#STM32_H7)).

In addition to [the products](https://docs.arduino.cc/) manufactured by the Arduino company, many more development boards have been created by other companies and the community. These boards typically provide the basic support circuitry for the primary microcontroller, convenient electrical interface, hardware programming interface, and sometimes include components (e.g., radio communication module, sensors) that add supplemental capabilities.

## About Arduino Tooling

The diversity of target hardware comes with a similar diversity of [toolchains](https://wikipedia.org/wiki/Toolchain). Arduino development software such as [**Arduino IDE**](https://docs.arduino.cc/software/ide-v2) provide a convenient high level user interface that allows the user to easily perform common operations such as compiling and uploading programs without needing to work with the individual toolchains directly.

## About Arduino Firmware

Each target microcontroller has its own complex and obscure low level firmware programming interface. One of the cornerstones of the Arduino initiative is a standardized firmware code [API](https://wikipedia.org/wiki/API) that provides an [abstraction layer](https://wikipedia.org/wiki/Abstraction_layer) over the architecture-specific low level interface. Ideally, a program that uses this API can run on any target.

### Arduino Cores

The code that implements the [fundamental Arduino firmware API](https://www.arduino.cc/reference/en/) for a given target is referred to as a "[core](https://arduino.github.io/arduino-cli/latest/platform-specification/#cores)".

### Arduino Libraries

Reusable firmware code may be packaged in an [Arduino library](https://arduino.github.io/arduino-cli/latest/library-specification/).

## About Arduino Platforms

An [Arduino platform](https://arduino.github.io/arduino-cli/latest/platform-specification/) contains everything needed to add support for a given target to the Arduino development software. This includes the configuration definitions that allow the development software to generate the appropriate toolchain commands, the Arduino core(s), and [supplemental libraries](https://arduino.github.io/arduino-cli/latest/platform-specification/#platform-bundled-libraries).

---

**ⓘ** Platforms are often [incorrectly](https://arduino.github.io/arduino-cli/latest/platform-specification/#platform-terminology) referred to as "cores". Although a platform may include a core, the core is only one component of a platform.

---

## About the **inoplatforms** Project

The Arduino hardware ecosystem is bounded by the array of available Arduino boards platforms. Although the platforms by prominent entities and for popular targets are well known, hundreds more lesser known platforms are also available. The goal of the **inoplatforms** project is to make it easier for the Arduino community to discover valuable and interesting platforms. This is done by providing a free, open source, comprehensive catalog of Arduino platforms.

More information about **inoplatforms** is available from the [**project repository readme**](https://github.com/per1234/inoplatforms#readme).

---

⚠ **inoplatforms** is a catalog of _every_ unique or significant Arduino platform in existence. Presence in the catalog does not in any way imply a platform is adequate or even functional.

The suitability and safety of platforms should be carefully evaluated by the user before installation and usage.

Source data provided by the catalog will assist in such evaluations.

---

## More Information

- [**Glossary**](/glossary/)
- [**Project documentation**](https://github.com/per1234/inoplatforms#readme)
- [**Arduino Platform Specification**](https://arduino.github.io/arduino-cli/latest/platform-specification/)
- [**Arduino Package Index Specification**](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/)
- [**Arduino IDE Boards Manager tutorial**](https://docs.arduino.cc/software/ide-v2/tutorials/ide-v2-board-manager)
- [**Arduino CLI `core` command documentation**](https://arduino.github.io/arduino-cli/latest/commands/arduino-cli_core/)
- [**Discuss Arduino boards platforms on Arduino Forum**](https://forum.arduino.cc/)
