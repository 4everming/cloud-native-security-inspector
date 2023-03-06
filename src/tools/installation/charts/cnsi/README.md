# Cloud Native Security Inspector (Project Narrows) Helm Chart

Cloud Native Security Inspector is an open source cloud native runtime security tool. It allows end users to assess
the security posture of Kubernetes clusters at runtime. This project will add dynamic scanning giving Security Auditors
greater awareness and control of running workloads.

## Requirements
- Kubernetes >= 1.24
- Helm >= 2.17.0
- We recommend you to have 8 GiB of memory available for this deployment, or at least 4 GiB for the minimum requirement.
Else, the deployment is expected to fail. 
## Installing
```shell
$ helm install [release-name] src/tools/installation/charts/cnsi
```
## Uninstalling
```shell
$ helm uninstall [release-name]
```
## Configuration
You can specify your own image registry and image tag to install.
```shell
$ helm install [release-name] src/tools/installation/charts/cnsi/ --set image.repository="your-own-repository" --set image.tag="latest"
```
If you'd like to install Cloud Native Security Inspector without OpenSearch:
```shell
$ helm install [release-name] src/tools/installation/charts/cnsi/ --set opensearch.enabled=false
```

| Parameter            | Description                                                                                   | Default                             |
|----------------------|-----------------------------------------------------------------------------------------------|-------------------------------------|
| `image.repository`   | `The repository to pull images`                                                               | `projects.registry.vmware.com/cnsi` |
| `image.tag`          | `The tags of the images`                                                                      | `0.3`                               |
| `image.pullPolicy`   | `The image pull policy in Kubernetes`                                                         | `IfNotPresent`                      |
| `opensearch.enabled` | `Flag to indicate if OpenSearch will be installed along with Cloud Native Security Inspector` | `true`                              |