<p align="center">
  <a href="https://rudderstack.com/">
    <img src="https://user-images.githubusercontent.com/59817155/121357083-1c571300-c94f-11eb-8cc7-ce6df13855c9.png">
  </a>
</p>

<p align="center"><b>The Customer Data Platform for Developers</b></p>

<p align="center">
  <b>
    <a href="https://rudderstack.com">Website</a>
    ·
    <a href="">Documentation</a>
    ·
    <a href="https://rudderstack.com/join-rudderstack-slack-community">Community Slack</a>
  </b>
</p>

---

# Rudderstack Observability Kit

Common libraries for building services with standard observability 

## Overview

This repo contains code for go, node and python; it developed using [NX](https://nx.dev/).
### Common Labels
We want labels used for observability to be shared across all language runtimes for consistancy so we are using [nx generators](https://nx.dev/extending-nx/recipes/local-generators) to generate the common labels using a single [labels file](./packages/generator-plugin/generator/src/generators/labels/labels.ts), [templates](./packages/generator-plugin/generator/src/generators/labels/files/) and [generator](./packages/generator-plugin/generator/src/generators/labels/generator.ts).

#### How to add new labels
1. Add new labels to this [file](./packages/generator-plugin/generator/src/generators/labels/labels.ts)
1. `npm run generate:labels`
1. Labels will be updated for each supported languages.
    * [Go](./packages/go/labels/common.go)
    * [Node](./packages/node/src/labels/common.ts)
    * [Python](./packages/python/labels/common.py)

### Packages
* [Go](./packages/go)
* [Node](./packages/node)
* [Pythob](./packages/python)

These packages contains common labels and interfaces required for observability.

## Contribute

We would love to see you contribute to RudderStack. Get more information on how to contribute [**here**](CONTRIBUTING.md).

## License

The RudderStack \*\*software name\*\* is released under the [**MIT License**](https://opensource.org/licenses/MIT).
