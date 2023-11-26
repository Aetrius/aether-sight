## SETUP
## You'd need to specify context for helm before deploying
## export HELM_KUBECONTEXT=dev


## Create Namespace
```bash
kubectl create namespace aethersight
```

## Installation
## ```helm install aethersight-chart aethersight-chart/ --namespace aethersight -f aethersight-chart/env/values.dev.yaml```


### Upgrade

## ```helm upgrade aethersight-chart aethersight-chart/ --namespace aethersight -f aethersight-chart/env/values.dev.yaml```


### Uninstall

##  ```helm uninstall aethersight-chart aethersight-chart/ --namespace aethersight```
