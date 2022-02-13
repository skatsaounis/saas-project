# SaaS Project

## Install microk8s

Follow getting started [guide](https://microk8s.io/docs/getting-started)

1. (Optional): Add autocompletion
1. Enable DNS and storage:

    ```bash
    microk8s enable dns storage
    microk8s status --wait-ready
    ```

## Install Demo WordPress blog

1. Install mySQL and WordPress, based on [example](https://kubernetes.io/docs/tutorials/stateful-application/mysql-wordpress-persistent-volume/):

    ```bash
    kubectl apply -f mysql.yaml
    kubectl apply -f wordpress.yaml
    ```

1. Expose WordPress to browser:

    ```bash
    kubectl port-forward service/wordpress 1080:80
    ```

1. Access from browser: <http://127.0.0.1:1080>


## Install AskMeAnything SaaS

### Get your node IP

```bash
# Mark IPv4
kubectl get node microk8s-vm -o wide
```

1. Build rest-api: Follow instructions in [README.md](./askmeanything/rest-api/README.md)
1. Build web-ui: Follow instructions in [README.md](./askmeanything/web-ui/README.md)
1. Install Kubernetes resources: Follow instructions in [README.md](./askmeanything/kubernetes/README.md)
1. Access from browser:

    * Web UI: `http://<node_IP>:30300`
    * Rest API: `http://<node_IP>:30301/questions`
