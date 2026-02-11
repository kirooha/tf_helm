sqlc:
	sqlc generate

docker:
	docker build -t eu.gcr.io/my-new-project-467616/kuber_practice_app:latest .
	docker push eu.gcr.io/my-new-project-467616/kuber_practice_app:latest
