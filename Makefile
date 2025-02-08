image_build_backend:
	docker build -t yvv4docker/userinfo-backend:latest -f Dockerfile.backend .

image_build_frontend:
	docker build -t yvv4docker/userinfo-frontend:latest -f Dockerfile.frontend .


image_push_backend:
	docker push yvv4docker/userinfo-backend:latest

image_push_frontend:
	docker push yvv4docker/userinfo-frontend:latest