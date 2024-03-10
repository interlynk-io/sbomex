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

#### Using containerized sbomex

```sh
docker run ghcr.io/interlynk-io/sbomex [search|pull] [options]
```
Example
```sh
docker run ghcr.io/interlynk-io/sbomex:latest search --format json --spec cdx --tool trivy --target '%box%' --limit 3
```
```
Unable to find image 'ghcr.io/interlynk-io/sbomex:latest' locally
latest: Pulling from interlynk-io/sbomex
bc89d6624a71: Already exists
bacb9c1935ff: Already exists
Digest: sha256:a00682b085fd21b7f071245a4d62d4825a07d9e783a8dfcda6b1f30f6a49514c
Status: Downloaded newer image for ghcr.io/interlynk-io/sbomex:latest
downloading db 100% |███████████████████████████| (89/89 MB, 5.4 MB/s)

A new version of sbomex is available v0.0.6.

  ID   TARGET          QUALITY  TYPE      CREATOR
  95   busybox:latest  3.25     cdx-json  trivy-0.36.1
  104  busybox:uclibc  3.25     cdx-json  trivy-0.36.1
  113  busybox:musl    3.25     cdx-json  trivy-0.36.1
```

# SBOM Card 
[![SBOMCard](https://api.interlynk.io/api/v1/badges?type=hcard&project_group_id=d6fbe787-51e6-44bc-a691-f792fb581f63
)](https://app.interlynk.io/customer/products?id=d6fbe787-51e6-44bc-a691-f792fb581f63&signed_url_params=eyJfcmFpbHMiOnsibWVzc2FnZSI6IkltVTJObUl4T0RFNUxUSXpaR1l0TkdFM09DMDRZVEptTFRkbE1EYzJZak13TTJJMk5pST0iLCJleHAiOm51bGwsInB1ciI6InNoYXJlX2x5bmsvc2hhcmVfbHluayJ9fQ==--9ab55c63454b3144b175f0439119cb442b5eae1bbfc5f18a9639a69d89487396)

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

# Other SBOM Open Source tools
- [SBOM Assembler](https://github.com/interlynk-io/sbomasm) - A tool to compose a single SBOM by combining other (part) SBOMs
- [SBOM Quality Score](https://github.com/interlynk-io/sbomqs) - A tool for evaluating the quality and completeness of SBOMs
- [SBOM Search Tool](https://github.com/interlynk-io/sbomagr) - A tool to grep style semantic search in SBOMs
- [SBOM Explorer](https://github.com/interlynk-io/sbomex) - A tool for discovering and downloading SBOM from a public repository

# Contact 
We appreciate all feedback. The best ways to get in touch with us:
- :phone: [Live Chat](https://www.interlynk.io/#hs-chat-open)
- 📫 [Email Us](mailto:hello@interlynk.io)
- 🐛 [Report a bug or enhancement](https://github.com/interlynk-io/sbomex/issues) 
- :x: [Follow us on X](https://twitter.com/InterlynkIo)

# Stargazers

If you like this project, please support us by starring it. 

[![Stargazers](https://starchart.cc/interlynk-io/sbomex.svg)](https://starchart.cc/interlynk-io/sbomex)
