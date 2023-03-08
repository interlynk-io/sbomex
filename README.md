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

# `sbomex`: Finding and fetching sbom's

[![Go Reference](https://pkg.go.dev/badge/github.com/interlynk-io/sbomex.svg)](https://pkg.go.dev/github.com/interlynk-io/sbomex)
[![Go Report Card](https://goreportcard.com/badge/github.com/interlynk-io/sbomex)](https://goreportcard.com/report/github.com/interlynk-io/sbomex)

`sbomex` is your primary tool that allows finding and downloading of SBOM's. 

## `sbomex` : Search sbom's
  You can sarch sbom through sbomlc.db, for searching sbom used below command

    ./sbomex search --format json --spec cdx --tool trivy
      ID   TARGET                   QUALITY  TYPE      CREATOR       
  5    centos:latest            7.39     cdx-json  trivy-0.36.1  
  14   centos:centos7.9.2009    7.38     cdx-json  trivy-0.36.1  
  23   centos:centos7           7.38     cdx-json  trivy-0.36.1  
  32   centos:7.9.2009          7.38     cdx-json  trivy-0.36.1  
  41   centos:7                 7.38     cdx-json  trivy-0.36.1  
  50   centos:centos8.4.2105    7.39     cdx-json  trivy-0.36.1  
  59   centos:centos8           7.39     cdx-json  trivy-0.36.1  
  68   centos:centos6.10        7.38     cdx-json  trivy-0.36.1  
  77   centos:centos6           7.38     cdx-json  trivy-0.36.1  
  86   centos:8.4.2105          7.39     cdx-json  trivy-0.36.1  
  95   busybox:latest           3.25     cdx-json  trivy-0.36.1  
  104  busybox:uclibc           3.25     cdx-json  trivy-0.36.1  
  113  busybox:musl             3.25     cdx-json  trivy-0.36.1  
  122  busybox:glibc            3.25     cdx-json  trivy-0.36.1  
  131  busybox:1-uclibc         3.25     cdx-json  trivy-0.36.1  
  140  busybox:1-musl           3.25     cdx-json  trivy-0.36.1  
  149  busybox:1-glibc          3.25     cdx-json  trivy-0.36.1  
  158  busybox:1                3.25     cdx-json  trivy-0.36.1  
  167  busybox:unstable-uclibc  3.25     cdx-json  trivy-0.36.1  
  176  busybox:unstable-glibc   3.25     cdx-json  trivy-0.36.1  
  185  ubuntu:latest            7.47     cdx-json  trivy-0.36.1  
  194  ubuntu:rolling           7.45     cdx-json  trivy-0.36.1  
  203  ubuntu:lunar-20230128    7.40     cdx-json  trivy-0.36.1  
  212  ubuntu:lunar             7.40     cdx-json  trivy-0.36.1  
  221  ubuntu:kinetic-20230126  7.45     cdx-json  trivy-0.36.1 
