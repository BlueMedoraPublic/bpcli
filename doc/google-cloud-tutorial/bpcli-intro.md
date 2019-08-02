# BPcli Tutorial

## Introduction

Welcome to a tutorial on using the `bpcli` tool.  Some prerequisties

* Signed up for a Bindplane account at [https://bindplane.bluemedora.com](https://bindplane.bluemedora.com)
* Make sure it's attached to a Stackdriver instance
* Some resources to manage

Let's get started!

## Install bpcli

Ok, let's install the latest version of bpcli in your cloudshell.  Just run the following commands

```bash
curl -LO https://github.com/BlueMedoraPublic/bpcli/releases/latest/download/bpcli_linux_amd64.zip && unzip bpcli_linux_amd64.zip && sudo mv bpcli /usr/local/bin/ && sudo chmod +x /usr/local/bin/bpcli
```

And a quick check to make sure it's working as it sould

```bash
bpcli version
```

## Setup Environment

* Go to your [Bindplane Profile Page](https://bindplane.bluemedora.com/profile) 
* Click the Generate Button
* With that key, run the following commands, pasting in the key after the `=` sign

```bash
export BINDPLANE_API_KEY=
```

Verify it works by listing your current collectors

```bash
bpcli collector list
```

## Make It Permanent

To make your changes more permanent, run the following command (inserting your API key)

```bash
echo export BINDPLANE_API_KEY=api-key-here >> ~/.bashrc
```

## Congratulations

You are up and running with bpcli!  Keep an eye on [the bpcli Github page](https://github.com/BlueMedoraPublic/bpcli) for updates and useful information.
