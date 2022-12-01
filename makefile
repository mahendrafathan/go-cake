.PHONY: all run

start:
	docker-compose up -d
	
stop:
	docker-compose down

remove-volume:
	docker-compose down -v