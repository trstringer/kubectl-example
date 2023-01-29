# kubectl example

Quickly and easily create example manifests so you don't have to remember Kubernetes resource manifest syntax.

## Installation

```
$ make install
```

*Note: This installs the plugin in ~/.local/bin. If you prefer to use a different directory in your `$PATH` then modify the `Makefile` prior to running `make install`.*

Verify installation:

```
$ kubectl example --help
Show example manifests for resources

Usage:
  kubectl-example [flags]

Flags:
  -h, --help   help for kubectl-example
  -l, --list   list available resources
```

## Usage

List out all available example manifests:

```
$ kubectl example --list
deployment
gateway [gateway.networking.k8s.io/v1beta1 networking.istio.io/v1beta1]
namespace
pod
service
virtualservice
```

*Note: This list will continue to grow as I add more example manifests. PRs welcome and appreciated!*

Dump out a deployment example manifest:

```
$ kubectl example deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: STRING
  namespace: STRING
spec:
  replicas: 1234
  selector:
    matchLabels:
      STRING: STRING
  template:
    metadata:
      labels:
        STRING: STRING
    spec:
      containers:
        - name: STRING
          image: STRING
          imagePullPolicy: STRING  # Always, IfNotPresent, Never
          ports:
            - containerPort: 1234
```

You can also dump multiple manifests at the same time:

```
$ kubectl example namespace pod
apiVersion: v1
kind: Namespace
metadata:
  name: STRING
---
apiVersion: v1
kind: Pod
metadata:
  name: STRING
  namespace: STRING
spec:
  containers:
    - name: STRING
      image: STRING
      imagePullPolicy: STRING  # Always, IfNotPresent, Never
      ports:
        - containerPort: 1234
```

Then, just modify the placeholders to your desired values, delete unwanted sections, add whatever you want and create these resources in your Kubernetes clusters!
