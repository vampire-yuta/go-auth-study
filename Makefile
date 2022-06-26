up:
	docker-compose up -d

ps:
	docker-compose ps

down:
	docker-compose down

logs:
	docker-compose logs -f

restart:
	docker-compose restart

login-db:
	docker-compose exec mysql mysql -uadmin -padmin -hlocalhost
