# SecOps CLI
The SecOps CLI allows users to interact with Google SecOps APIs without the need for service account keys using `gcloud` application default credentials.

## Prerequisites
The gcloud CLI tool is required for authenticating with Google APIs.

https://cloud.google.com/sdk/docs/install

## Usage
Refresh authentication with:
```
gcloud auth login --update-adc
```
Commands follow the following structure:
```
secops [resource] [method]
```
