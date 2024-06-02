
rest_with_postgres:
	rm -rf cmd/.env
	echo "DATABASE=postgres" > cmd/.env
	echo "SERVICE=rest" >> cmd/.env
	docker-compose up -d --build postgres_db rest_server