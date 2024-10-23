docker build -t k8s-test-app .
docker run --rm -p 3000:3000 k8s-test-app



```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:latest
        ports:
        - containerPort: 8080
```



after error with:  Failed to pull image "k8s-test-app": failed to pull and unpack image "docker.io/library/k8s-test-app:latest":
````
kind load docker-image k8s-test-app:1.0.0
```

````
kubectl delete deployment --all
```


```


Service wont work. To access service:

kind delete cluster 
kind create cluster --config=kind.yaml




1. **Template the Namespace in Your Deployment Files**

  ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: {{ include "myapp.fullname" . }}
     namespace: {{ .Release.Namespace }}
   spec:

   ```

2. **Update Helper Templates (if any)**
 to use `{{ .Release.Namespace }}` 

### Step 2: Prepare Namespace-Specific Values
```plaintext
my-chart/
  └─ values/
     ├─ values-stage.yaml
     ├─ values-prod.yaml
```


And for `values-prod.yaml`, you might have:

```yaml
replicaCount: 3
image:
  tag: "prod"
resources:
  limits:
    cpu: 500m
    memory: 512Mi
```

### Step 3: Deploy to Each Namespace

You can now deploy your Helm chart to each namespace by specifying the namespace and the corresponding values file with Helm’s `--namespace` and `-f` options.

1. Create the Namespaces if They Don’t Exist

   ```bash
   kubectl create namespace stage
   kubectl create namespace prod
   ```

2. Deploy to Stage Namespace

   ```bash
   helm install myapp-stage my-chart/ --namespace stage -f my-chart/values/values-stage.yaml
   ```

3. Deploy to Prod Namespace

   ```bash
   helm install myapp-prod my-chart/ --namespace prod -f my-chart/values/values-prod.yaml
   ```

### Step 4: Manage Upgrades and Rollbacks


```bash
helm upgrade myapp-stage my-chart/ --namespace stage -f my-chart/values/values-stage.yaml
helm upgrade myapp-prod my-chart/ --namespace prod -f my-chart/values/values-prod.yaml
```





```
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: myapp:latest
          env:
            - name: HELM_RELEASE_NAME
              value: "{{ .Release.Name }}"
            - name: HELM_RELEASE_NAMESPACE
              value: "{{ .Release.Namespace }}"
              ```