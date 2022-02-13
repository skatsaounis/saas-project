# AskMeAnythingUI

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 13.2.3.

Inspired by: <https://www.section.io/engineering-education/angular12-material-table/>

## Change node IP in settings

Update `src/environments/environment.prod.ts` with your node IP in `questionsURL` field.

## Build Image

```bash
docker build -t webui .
docker save webui > myimage.tar
multipass transfer myimage.tar microk8s-vm:
microk8s ctr image import myimage.tar
rm myimage.tar
microk8s ctr images ls | grep webui
```
