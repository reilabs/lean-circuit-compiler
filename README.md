<a href="https://reilabs.io">
<picture>
  <source width="150" media="(prefers-color-scheme: dark)" srcset="https://github.com/reilabs/gnark-lean-extractor/assets/35899/c04bdb7f-4c31-4264-acb6-a96f32c6cc29">
  <source width="150" media="(prefers-color-scheme: light)" srcset="https://github.com/reilabs/gnark-lean-extractor/assets/35899/fc11280b-e3e5-4a6f-83da-788884083c36">
  <img width="150" src="https://github.com/reilabs/gnark-lean-extractor/assets/35899/fc11280b-e3e5-4a6f-83da-788884083c36">
</picture>
</a>

# Lean Circuit Compiler

This repository contains a Go library that transpiles
[zero-knowledge](https://en.wikipedia.org/wiki/Zero-knowledge_proof) (ZK)
circuits from [Go](https://go.dev) to [Lean](https://leanprover.github.io). In
particular, it deals with circuits constructed as part of the
[gnark](https://github.com/ConsenSys/gnark) proof system.

This makes it possible to take existing gnark circuits and export them to Lean
for later formal verification.

This library contains the core of the extractor to be used in conjunction with [gnark-lean-extractor](https://github.com/reilabs/gnark-lean-extractor)
which is the interface layer with [gnark](https://github.com/ConsenSys/gnark).

For an overview of how to use this library, see the documentation in [gnark-lean-extractor](https://github.com/reilabs/gnark-lean-extractor).
If you are interested in contributing, or are new to Go, please see our
[contributing guidelines](./CONTRIBUTING.md) for more information.
