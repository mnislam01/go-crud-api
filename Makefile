build:
	docker build -t movieapis .
start:
	docker run -p 8080:8000 --name newmovieapis movieapis
