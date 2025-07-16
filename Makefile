secret-generic:
	kubectl create secret generic rarible-secret \
	--from-env-file=.env

helm-install:
	helm install rarible-app ./helm/rarible-app --namespace default

helm-upgrade:
	helm upgrade rarible-app ./helm/rarible-app --namespace default
