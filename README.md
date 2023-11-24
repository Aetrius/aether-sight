# aether-sight
### Project Aether-Sight
```
"Aether" traditionally refers to a hypothetical substance or medium that fills the region of the universe above the terrestrial sphere. In ancient cosmology and physics, it was considered the material that made up the heavens or the upper regions of space. It was thought to be a medium through which light and other electromagnetic waves could travel.

In more contemporary contexts, "aether" can symbolize a concept of the quintessence or the essence of something intangible, often associated with the mystical or the transcendental. It's sometimes used poetically or metaphorically to denote an ethereal or otherworldly quality.

Sight - is well sight. Do what you will. Not a permanent name, just eazy.

Long term vision: vnc/ssh/ station image/ai analysis.
```
## RUN Image Capture
```bash
go run ./cmd/aetherImageCapture 
```

## ROADMAP
### Phase 1: 
Station image upload, api to retrieve images as binary forms.
Language: golang
Cache: likely Redis


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

## Run Aether API

```bash
go run ./cmd/aetherAPI
```


## RUN Aether Image Capture
```bash
go run ./cmd/aetherImageCapture
```