# Ask Me Anything REST API

Inspired by: <https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh>

## Build API Docker image

```bash
docker build -t restapi .
docker save restapi > myimage.tar
multipass transfer myimage.tar microk8s-vm:
microk8s ctr image import myimage.tar
rm myimage.tar
microk8s ctr images ls | grep restapi
```
