# crybaby

crybaby is a noisy log message generator.
Its sole purpose in life is to generate log messages from a container.

Someday, it may accept parameters to alter its runtime behavior.
At the moment, it's probably not all that useful.

## Usage

### Docker

```sh
docker run --rm -it ghcr.io/djschaap/crybaby:main
```

### Podman

```sh
podman run --rm -it ghcr.io/djschaap/crybaby:main
```

### Kubernetes

```sh
kubectl create job cry-me-a-river --image=ghcr.io/djschaap/crybaby:main
```
