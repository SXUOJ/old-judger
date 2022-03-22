build: Dockerfile
	docker build . -t judger

run: 
	docker run -itd --name judger -p 8080:8080 judger:latest

delete: 
	docker rm -f judger

