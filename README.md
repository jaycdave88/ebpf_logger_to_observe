# EBPF User Agent

This application serves as a Minimum Viable Product (MVP) to capture logs using eBPF (Extended Berkeley Packet Filter) and ship them over HTTP to Observe.

## Getting Started

These instructions will cover usage information and for the docker container 

## Prerequisites

In order to run this container you'll need docker installed.

* [Get Docker](https://docs.docker.com/get-docker/)
* Docker running on Apple Silicon if you are using a Mac with the M1 chip or newer.

## Usage

### Container Parameters

To build and run the Docker container on Apple Silicon, use the following command:

```docker build --platform linux/arm64 -t ebpf-user-agent .```

To execute the MVP, run the following command:

```docker run --privileged ebpf-user-agent```

Note: The --privileged flag gives the container full access to the host, which can be a security risk. Only use it if you trust the application and understand the implications.

### Environment Variables
```OBSERVE_HTTP_ENDPOINT``` - The HTTP endpoint to which the logs will be shipped.
```BEARER_TOKEN``` - The bearer token used for authentication to the HTTP endpoint.

These variables are to be set in the ```src/.env``` file.
