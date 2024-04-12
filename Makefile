docker:
	docker build . -t http-user-service

maiden_deploy: docker
	kind create cluster --config=kind.config.yaml
	kind load docker-image http-user-service:latest
	helm package http-user-service
	helm install http-dev http-user-service-0.1.0.tgz
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

teardown:
	helm uninstall http-dev
	kind delete cluster
