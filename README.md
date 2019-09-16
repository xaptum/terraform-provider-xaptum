# Terraform Provider for the Xaptum ENF #

[![Release](https://img.shields.io/github/release/xaptum/terraform-provider-enf.svg)](https://github.com/xaptum/terraform-provider-enf/releases)
[![Build Status](https://travis-ci.com/xaptum/terraform-provider-enf.svg?branch=master)](https://travis-ci.com/xaptum/terraform-provider-enf)

This is a Terraform provider for managing ENF resources on the
[Xaptum](https://www.xaptum.com) Edge Network Fabric (ENF), a secure
overlay network for IoT.

## Installation ##

1. Download the latest compiled binary from [GitHub releases](https://github.com/xaptum/terraform-provider-enf/releases).

1. Unzip/untar the archive.

1. Move it into `$HOME/.terraform.d/plugins`:

   ```sh
   $ mkdir -p $HOME/.terraform.d/plugins
   $ mv terraform-provider-enf $HOME/.terraform.d/plugins/terraform-provider-enf
   ```

1. Create your Terraform configurations as normal, and run `terraform init`:

   ```sh
   $ terraform init
   ```

   This will find the plugin locally.

## Development ##

1. Install the plugin to the Terraform plugins directory:

    ```sh
    $ make install
    ```

1. Build the plugin from source:

    ```sh
    $ make build
    ```

1. Create your Terraform configurations as normal, and run `terraform init`:

    ```sh
    $ terraform init
    ```

1. Repeat steps 2 and 3 after making any changes to the provider code.


## License ##
Copyright 2019 Xaptum, Inc.

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
