secret-generic:
	kubectl create secret generic rarible-secret \
	--from-env-file=.env

port-forward:
	kubectl port-forward svc/rarible 8080:8080 

helm-install:
	helm install rarible-app ./helm/rarible-app --namespace default

helm-upgrade:
	helm upgrade rarible-app ./helm/rarible-app --namespace default
