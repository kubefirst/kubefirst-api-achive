# kubefirst-api

Kubefirst API that serves console frontend

## Running the API

1. Locally: `KUBEFIRST_VERSION=1.9.9 HOSTED_ZONE_NAME=gh.mgmt.kubefirst.com go run .`
2. Docker: build `docker build -f build/Dockerfile -t api .` and run `docker run -d -p 3000:3000 -t -e KUBEFIRST_VERSION={kubefirstVersion} -e HOSTED_ZONE_NAME={hostedZoneName} api`
3. Docker compose: `docker-compose up -d`

## Build the Docker Image

1. Build the docker image: `docker build -f build/Dockerfile --tag public.ecr.aws/kubefirst/api:{version} .`
2. Test the docker image `docker run -d -p 3000:3000 -t public.ecr.aws/kubefirst/console:{version}`
3. Go to `http://localhost:3000/healthz` and check the response

### Push Docker Image to ECR

Note: You might need an AWS Management Account

1. Retrieve an authentication token and authenticate your Docker client to your registry.
   `aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/kubefirst`
2. Run the following command to push this image to your ECR repository:
   `docker push public.ecr.aws/kubefirst/api:{version}`
3. Go to the [API Chart](https://github.com/kubefirst/charts/tree/gh-pages/charts/api) and bump up the docker version and chart version. Follow the [README](https://github.com/kubefirst/charts) instructions.

### Chart

1. The Api chart is stored in the [Charts repository](https://github.com/kubefirst/charts)
