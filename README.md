# aether-sight
### Project Aether-Sight
```
"Aether" traditionally refers to a hypothetical substance or medium that fills the region of the universe above the terrestrial sphere. In ancient cosmology and physics, it was considered the material that made up the heavens or the upper regions of space. It was thought to be a medium through which light and other electromagnetic waves could travel.

In more contemporary contexts, "aether" can symbolize a concept of the quintessence or the essence of something intangible, often associated with the mystical or the transcendental. It's sometimes used poetically or metaphorically to denote an ethereal or otherworldly quality.

Sight - is well sight. Do what you will. Not a permanent name, just eazy.

Long term vision: vnc/ssh/ station image/ai analysis.
```

## ROADMAP
### Phase 1: Complete
Station image upload, api to retrieve images as binary forms.
Language: golang
Cache: Redis

### Phase 1.5: Create Helm Charts & Docker Images
Docker Images: Through github workflow
Helm Charts: Build into the k8s helm directory


### Phase 2: 
UI to view images
Lanugage: React + Vite + ??
Database: likely Postgres (because it's freeemium)

### Phase 3:
RBAC applied to API
RBAC applied to UI
Database: ??? whatever fills the need


### Phase 4
AI: based on cached binary images. 
-- Some form of AI to review the images, find patterns like user behavior.
-- Review anything that looks like an error or error states

## Run Aether API - interfaces to redis

```bash
go run ./cmd/aetherAPI
```

## RUN Aether Image Capture - client code
```bash
go run ./cmd/aetherImageCapture
```

## RUN Aether Redis Consumer - server consumer
```bash
go run ./cmd/aetherImageCapture
```


### K8s Redis Pre-req
Install the base yaml configurations
- redis-commander
- redis-master-pvc
- redis-replica-1-pvc
- redis-replica-2-pvc
- redis-svc

### K8s Redis installation
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install my-release bitnami/redis
helm upgrade my-release bitnami/redis --set auth.enabled=false
```

### K8s Redis uninstall
```bash
helm uninstall my-release bitnami/redis
```



### Linux/Ubuntu Hosts
Install Gnome-Screenshot
```bash
sudo apt install gnome-screenshot
```

## Run Client Computer Docker Image
install wget if you don't have the package...

```bash
wget https://github.com/Aetrius/aether-sight/blob/main/docker-compose.yaml
docker-compose up -d
```

