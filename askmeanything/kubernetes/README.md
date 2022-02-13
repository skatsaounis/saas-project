# Ask Me Anything in Kubernetes

## Install MongoDB

Inspired by: <https://github.com/allyjunio/microk8s-mongodb-demo>

```bash
TMPFILE=$(mktemp)
openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE
rm $TMPFILE
kubectl apply -f mongo.yaml
```

### Initialize replica-set

```bash
kubectl exec mongo-0 -c mongod-container -- mongo --eval 'rs.initiate({_id: "MainRepSet", version: 1, members: [ {_id: 0, host: "mongo-0.mongodb-service.default.svc.cluster.local:27017"}, {_id: 1, host: "mongo-1.mongodb-service.default.svc.cluster.local:27017"}, {_id: 2, host: "mongo-2.mongodb-service.default.svc.cluster.local:27017"} ]});'
kubectl exec mongo-0 -c mongod-container -- mongo --eval 'rs.status();'
```

### Create admin user

```bash
PASSWORD="123abc"
kubectl exec mongo-0 -c mongod-container -- mongo --eval 'db.getSiblingDB("admin").createUser({user:"main_admin",pwd:"'"${PASSWORD}"'",roles:[{role:"root",db:"admin"}]});'
```

### Create restapi user and database

```bash
# Login to mongo container
kubectl exec -it mongo-0 -c mongod-container bash
# Login to mongo shell with admin
mongo mongodb://main_admin:123abc@mongodb-service/admin
# Create database
db.getSiblingDB("golangAPI").createUser({user:"golangAPI",pwd:"123abc",roles:[{role:"readWrite",db:"golangAPI"}]});
```

## Install web-ui and rest-api

```bash
kubectl apply -f restapi.yaml
kubectl apply -f webui.yaml
```

### Enable cros-site origin from your browser

For Example in Safari: `Develop -> Disable Cros-Origin Restrictions`.
