# kubefirst-api
Kubefirst API that serves console frontend

## Running the API

1. Locally: `KUBEFIRST_VERSION=1.9.9 HOSTED_ZONE_NAME=gh.mgmt.kubefirst.com go run .`
2. Docker: build `docker build -f build/Dockerfile -t api .` and run `docker run -d -p 3000:3000 -t -e KUBEFIRST_VERSION=1.9.9 -e HOSTED_ZONE_NAME=gh.mgmt.kubefirst.com api` 
3. Docker compose: `docker-compose up -d`

## Build the Docker Image

1. Build the docker image: `docker build -f build/Dockerfile --tag public.ecr.aws/kubefirst/api:{version} .`
2. Log in `aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/kubefirst`
3. Push Image: `docker push public.ecr.aws/kubefirst/api:0.0.1`
