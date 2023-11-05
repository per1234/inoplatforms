# inoplatforms Registry

[`registry.yml`](registry.yml) is the **inoplatforms** registry file.

The **inoplatforms** catalog is built around this data file. It provides the following types of information:

- Pointers to the external sources of information (e.g., package index, repository) from which to automatically collect catalog data from
- Supplemental information not available for automated gathering from the external sources

[`registry.yml`](registry.yml) is written in the [YAML](https://www.yaml.info/learn/index.html) language.

The registry is a list of registration entries of "package providers". A "package provider" will typically be an [Arduino package index](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/). However, the registration of packages for which no package index has been created is also supported.

Instructions for adding platforms to the **inoplatforms** catalog are available in the [**Contributor Guide**](https://www.inoplatforms.info/contributor-guide/).

## Fields

Each registry entry is a mapping (AKA "object") that contains a set of standardized fields.

---

**ⓘ** Fields are optional unless marked "**Required**"

---

### `id`

**Required**

Each package provider is assigned a permanent unique numeric identifier.

You should use the next available integer (e.g., if the previous registry entry has `id: 42`, then the next one should have `id: 43`).

### `package-index`

**Required unless `packages` is present**

This is a mapping that contains data about the [Arduino package index](https://arduino.github.io/arduino-cli/latest/package_index_json-specification/) that provides the packages.

### `package-index.url`

**Required**

The publication URL of the package index.

This is the URL suitable for use in the [**Additional Boards Manager URLs** field of **Arduino IDE** preferences](https://docs.arduino.cc/learn/starting-guide/cores#how-to-install-a-third-party-core), the `board_manager.additional_urls` field of [the **Arduino CLI** configuration](https://arduino.github.io/arduino-cli/latest/configuration/#configuration-keys), etc.

### `package-index.source`

This is a mapping that contains data about the source code of the package index file.

### `package-index.source.repository`

This is a mapping that contains data about the version control repository of the package index source code.

### `package-index.source.repository.url`

**Required**

The URL of the package index source code repository.

### `package-index.source.repository.ref`

**Required**

The [Git reference](https://git-scm.com/book/en/v2/Git-Internals-Git-References) of the package index source code in the repository.

This will typically be a branch name (e.g., `main`). However, it could also be a commit hash or tag name in the case where the package index source is not available at the tip of a branch.

### `package-index.source.repository.path`

**Required**

The path of the package index source file in the repository.

This is the path to the source file itself.

The path should start with a `/` (indicating the root of the repository).

### `package-index.source.repository.notes`

Any supplemental notes about the repository for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `package-index.source.notes`

Any supplemental notes about the package index source code for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `package-index.notes`

Any supplemental notes about the package index for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[]`

**Required unless `package-index` is present**

This is a list of the packages from this provider.

An Arduino package is a container for two distinct components:

- Platforms
- Tools

### `packages[*].name`

**Required**

This is the machine identifier for the package (sometimes referred to as the "vendor" or "packager" name).

### `packages[*].platforms[]`

This is a list of the platforms from this package.

### `packages[*].platforms[*].architecture`

**Required**

This is the machine identifier for the platform.

### `packages[*].platforms[*].installation-references[]`

This is a list of references for user installation of the platform.

### `packages[*].platforms[*].installation-references[*].url`

**Required**

The URL of the reference.

### `packages[*].platforms[*].installation-references[*].label`

A text label to use for the link to the reference.

### `packages[*].platforms[*].source.repository`

This is a mapping that contains data about the version control repository of the platform source code.

### `packages[*].platforms[*].source.repository.url`

**Required**

The URL of the platform source code repository.

### `packages[*].platforms[*].source.repository.ref`

**Required**

The [Git reference](https://git-scm.com/book/en/v2/Git-Internals-Git-References) of the platform source code in the repository.

This will typically be a branch name (e.g., `main`). However, it could also be a commit hash or tag name in the case where the platform source is not available at the tip of a branch.

### `packages[*].platforms[*].source.repository.path`

**Required**

The path of the platform source code in the repository.

This is the path to the folder that contains the platform's `boards.txt` and/or `platform.txt` files.

The path should start with a `/` (indicating the root of the repository).

### `packages[*].platforms[*].source.repository.notes`

Any supplemental notes about the repository for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].platforms[*].source.notes`

Any supplemental notes about the platform source code for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].platforms[*].tools-dependencies[]`

This is a list of data about the platform's tool dependencies.

---

❗ It is not necessary to add tool dependency data for tools that come from the same package as the platform.

---

### `packages[*].platforms[*].tools-dependencies[*].packager`

**Required**

The name of the package of tool dependencies of the platform.

### `packages[*].platforms[*].tools-dependencies[*].package-provider-id`

**Required**

The identifier of the package provider of the tool dependencies package.

---

**ⓘ** This is required because multiple package providers might provide a package of the same `name` and the catalog must be able to identify the exact tool that is being referenced as a dependency.

---

### `packages[*].platforms[*].notes`

Any supplemental notes about the platform for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].tools[]`

This is a list of the tools from this package.

### `packages[*].tools[*].name`

**Required**

The machine identifier of the tool.

### `packages[*].tools[*].source.repository`

This is a mapping that contains data about the version control repository of the tool's source code.

### `packages[*].tools[*].source.repository.url`

**Required**

The URL of the tool's source code repository.

### `packages[*].tools[*].source.repository.ref`

**Required**

The [Git reference](https://git-scm.com/book/en/v2/Git-Internals-Git-References) of the tool's source code in the repository.

This will typically be a branch name (e.g., `main`). However, it could also be a commit hash or tag name in the case where the package index source is not available at the tip of a branch.

### `packages[*].tools[*].source.repository.path`

**Required**

The path of the tool's source code in the repository.

The path should start with a `/` (indicating the root of the repository).

### `packages[*].tools[*].source.repository.notes`

Any supplemental notes about the repository for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].tools[*].source.notes`

Any supplemental notes about the tool's source code for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].tools[*].notes`

Any supplemental notes about the tool for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `packages[*].notes`

Any supplemental notes about the package for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

### `notes`

Any supplemental notes about the package provider for inclusion in the catalog.

---

**ⓘ** For internal notes about the registry data which are not of interest to the catalog users, use [YAML comments](https://yaml.org/refcard.html) instead of the `notes` field.

---

## `github.com/per1234/inoplatforms/registry/assets/go-registry` Module

A [**Go**](https://go.dev/) module is provided to make it easy for community members to work with the registry data in their own projects.

See [**the module's documentation**](assets/go-registry/README.md) for more information.
