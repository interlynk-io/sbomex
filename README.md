<!--
 Copyright 2023 Interlynk.io
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# `sbomex`: SBOM Explorer

[![Go Reference](https://pkg.go.dev/badge/github.com/interlynk-io/sbomex.svg)](https://pkg.go.dev/github.com/interlynk-io/sbomex)
[![Go Report Card](https://goreportcard.com/badge/github.com/interlynk-io/sbomex)](https://goreportcard.com/report/github.com/interlynk-io/sbomex)

`sbomex` is a command line utility to help query and pull from Interlynk's public SBOM repository. The tool is intended to help familiarize with the specifications and formats of common SBOM standards and the quality of produced SBOMs (See [sbomqs](https://github.com/interlynk-io/sbomqs/) - SBOM Quality Score for how the score is computed). 

The underlying repository is updated periodically with SBOMs from a variety of sources built with many tools.


## `sbomex search` : Search repository for matching SBOMs
search commands finds SBOMs in the repository that matches given filtering criteria (specification, format or tool name)

```sh
sbomex search --format json --spec cdx --tool trivy --target '%centos%7' --limit 3
```
```
  ID  TARGET                 QUALITY  TYPE      CREATOR
  14  centos:centos7.9.2009  7.38     cdx-json  trivy-0.36.1
  23  centos:centos7         7.38     cdx-json  trivy-0.36.1
  32  centos:7.9.2009        7.38     cdx-json  trivy-0.36.1
```

## `sbomex pull` : Downloads specified SBOM from the repository and prints to the screen
```sh
sbomex pull --id 23
 ```
 ```
 {
	"SPDXID": "SPDXRef-DOCUMENT",
	"creationInfo": {
		"created": "2023-03-01T01:32:02.939561Z",
		"creators": [
			"Tool: trivy",
			"Organization: aquasecurity"
		]
	},
	"dataLicense": "CC0-1.0",
 ...
 ```
 
# Installation 

## Using Prebuilt binaries 

```console
https://github.com/interlynk-io/sbomex/releases
```

## Using Homebrew
```console
brew tap interlynk-io/interlynk
brew install sbomex
```

## Using Go install

```console
go install github.com/interlynk-io/sbomex@latest
```

## Using repo

This approach invovles cloning the repo and building it. 

1. Clone the repo `git clone git@github.com:interlynk-io/sbomex.git`
2. `cd` into `sbomex` folder 
3. make build
4. To test if the build was successful run the following command `./build/sbomex version`


# Contributions
We look forward to your contributions, below are a few guidelines on how to submit them 

- Fork the repo
- Create your feature/bug branch (`git checkout -b feature/new-feature`)
- Commit your changes (`git commit -am "awesome new feature"`)
- Push your changes (`git push origin feature/new-feature`)
- Create a new pull-request
