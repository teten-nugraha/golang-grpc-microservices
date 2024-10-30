# Nama Docker images
USER_IMAGE=user-service
PRODUCT_IMAGE=product-service
CHECKOUT_IMAGE=checkout-service

# Nama file build image
USER_DOCKERFILE=user-service/Dockerfile
PRODUCT_DOCKERFILE=product-service/Dockerfile
CHECKOUT_DOCKERFILE=checkout-service/Dockerfile

# Nama file deployment YAML untuk Kubernetes
USER_DEPLOYMENT=user-service/user-deployment.yaml
PRODUCT_DEPLOYMENT=product-service/product-deployment.yaml
CHECKOUT_DEPLOYMENT=checkout-service/checkout-deployment.yaml
NGINX_DEPLOYMENT=nginx/nginx-deployment.yaml
ENVOY_DEPLOYMENT=envoy/envoy-deployment.yaml

# Build Docker images
build: build-user build-product build-checkout

build-user:
	docker build -t $(USER_IMAGE) -f $(USER_DOCKERFILE) .

build-product:
	docker build -t $(PRODUCT_IMAGE) -f $(PRODUCT_DOCKERFILE) .

build-checkout:
	docker build -t $(CHECKOUT_IMAGE) -f $(CHECKOUT_DOCKERFILE) .

# Deploy ke Kubernetes
deploy: deploy-user deploy-product deploy-checkout deploy-envoy

deploy-user:
	kubectl apply -f $(USER_DEPLOYMENT)

deploy-product:
	kubectl apply -f $(PRODUCT_DEPLOYMENT)

deploy-checkout:
	kubectl apply -f $(CHECKOUT_DEPLOYMENT)

deploy-envoy:
	kubectl apply -f $(ENVOY_DEPLOYMENT)

# Jalankan semua perintah
all: build deploy
