sqlc:
	sqlc generate

docker:
	docker buildx build --platform linux/amd64 -t eu.gcr.io/my-new-project-467616/kuber_practice_app:latest .
	docker push eu.gcr.io/my-new-project-467616/kuber_practice_app:latest
	docker buildx build --platform linux/amd64 -f Dockerfile.redis -t eu.gcr.io/my-new-project-467616/kuber_practice_app_redis:latest .
	docker push eu.gcr.io/my-new-project-467616/kuber_practice_app_redis:latest
