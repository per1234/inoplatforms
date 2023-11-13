---
title: Platform Submission Guide
breadcrumb: submission
layout: basic
---

---

❗ Do you need help or have a question about using this project? Support requests should be made to the [Arduino Forum](https://forum.arduino.cc).

---

Since there is no way to automatically collect the list of available <a href="/glossary/#platform" target="_blank">Arduino boards platforms</a>, the **inoplatforms** catalog is built from a manually curated registry. Submissions of platforms not already in the catalog are very welcome and a valuable contribution to the project.

We want to catalog **every** unique boards platform. This includes platforms that are in an unfinished state, since one of the goals of the project is to encourage community collaboration on platforms. This also includes platforms that are obsolete, since they may be of historical interest. The only cause for exclusion is if a platform is purely a copy of another, without any meaningful differences (as is common in the case where a fork is created by someone solely for the sake of having a backup available if the parent were to be lost).

## Prerequisites

In order for the registry maintainers to efficiently track them, platform submissions are made via issues in the **inoplatforms** project's **GitHub** repository. For this reason, you will need a [**GitHub**](https://github.com/) account to make a submission.

If you don't have one already, it is free and easy to sign up. Since **GitHub** is the most popular site for hosting and collaborating on open source software projects, you will likely find it useful to have an account for other purposes as well over time. You can get started by clicking the "**Sign up**" button on the **GitHub** home page:

{{< github-url href="https://github.com/" >}}

## Procedure

1. Search the catalog to see if the platform was already registered:<br />
   <a href="https://github.com/per1234/inoplatforms/blob/main/ino-hardware-package-list.tsv" target="_blank">https://github.com/per1234/inoplatforms/blob/main/ino-hardware-package-list.tsv</a>
1. Click the link below to begin composing the submission<br />
   <a href="https://github.com/per1234/library-registry/issues/new?assignees=&labels=topic%3A+rename&projects=&template=rename.yml&title=Library+name+change+request" target="_blank">https://github.com/per1234/library-registry/issues/new?assignees=&labels=topic%3A+rename&projects=&template=rename.yml&title=Library+name+change+request</a>
1. You will see a form with fields to fill in with information about the submission. If you don't know what to enter in a field, or if the information is not available, it is fine to leave it empty. The registry maintainers will find the missing information.
1. Once you have finished filling in the fields, click the "**Submit new issue**" button at the bottom of the page.<br />
   Thanks for your submission!

A maintainer will process your submission and register the platform. They will close the submission issue you created at that time and you will receive a notification of this event from **GitHub**. Soon after that (within less than a day's time), the **inoplatforms** catalog will be updated and the submitted platform will be listed there.

## Direct Registration

Submission via **GitHub** issues as described [above](#procedure) is recommended due to being the most simple option for the contributor. However, if you would like to submit a pull request to add a platform to the registry directly, that is also welcome. See the [**Pull Request Guide**](/contributor-guide/pull-requests/) and [**Development Guide**](/contributor-guide/pull-requests/development-guide/) for information.
