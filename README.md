<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/rts-core/versioner">
    <img src="images/logo.png" alt="Logo" width="300">
  </a>

  <p align="center">
    Helm Repository Values for on prem
    <br />
    <a href="https://github.com/rts-core/versioner/issues">Report Bug</a> |
    <a href="https://github.com/rts-core/versioner/issues">Request Feature</a>
  </p>
</p>

### Built With

* [Kubernetes](https://kubernetes.io/)
* [Docker](https://www.docker.com/)
* [Helm](https://helm.sh/)
* [versioner](https://versioner.com/)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

Requires a functional K8s cluster with admin access.

### Installation

1. Create the namespace
   ```sh
   kubectl create namespace versioner
   ```
2. Add Helm Repo
   ```sh
   helm repo add bitnami https://charts.bitnami.com/bitnami
   ```
2. Install the chart
   ```sh
   helm install versioner bitnami/versioner -f values.yaml -n versioner \
        --set auth.rootPassword=<Pass>
   ```


### Usage

Adding Repository to helm:
```sh
helm repo add rts http://versioner/
```

List Repository Charts:
```sh
helm search repo rts
```

Push Chart:

Use the Helm Push Plugin found [Here](https://github.com/chartmuseum/helm-push)
```sh
helm push mychart/ rts
```

<!-- CONTACT -->
## Contact

Project Link: [https://github.com/rts-core/versioner](https://github.com/rts-core/versioner)

Copyright [2021] [Real Technology Solutions LLC]