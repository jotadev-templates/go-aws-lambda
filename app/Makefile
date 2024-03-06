# Environment
#TF_DIR=/path # create an environment variable in your terminal

#PROVIDE_AWS_ACCOUNT=123456 # create an environment variable in your terminal
PROVIDE_AWS_REGION=sa-east-1

APP_NAME=app-crm-lambda
APP_VERSION=v1.0.0
APP_LOG_LEVEL=DEBUG

#########################################################################################

go-build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ${APP_NAME}

docker-build:
	docker build -t ${APP_NAME}:${APP_VERSION} .

docker-tag:
	docker tag ${APP_NAME}:${APP_VERSION} ${PROVIDE_AWS_ACCOUNT}.dkr.ecr.${PROVIDE_AWS_REGION}.amazonaws.com/${APP_NAME}:latest

docker-push:
	docker push ${PROVIDE_AWS_ACCOUNT}.dkr.ecr.${PROVIDE_AWS_REGION}.amazonaws.com/${APP_NAME}:latest

aws-login-ecr:
	aws ecr get-login-password \
	--region ${PROVIDE_AWS_REGION} \
	| docker login \
 	--username AWS \
 	--password-stdin ${PROVIDE_AWS_ACCOUNT}.dkr.ecr.${PROVIDE_AWS_REGION}.amazonaws.com

aws-create-repo:
	aws ecr create-repository \
	--repository-name ${APP_NAME} \
	--image-scanning-configuration scanOnPush=true \
	--image-tag-mutability MUTABLE \
	--region ${PROVIDE_AWS_REGION}

aws-delete-repo:
	aws ecr delete-repository \
	--repository-name ${APP_NAME} \
	--force \
	--region ${PROVIDE_AWS_REGION}

aws-delete-log-group:
	aws logs delete-log-group \
 	--log-group-name /aws/lambda/${APP_NAME} \
 	--region ${PROVIDE_AWS_REGION}

aws-update-function:
	aws lambda update-function-code \
	--function-name ${APP_NAME} \
    --image-uri=${PROVIDE_AWS_ACCOUNT}.dkr.ecr.${PROVIDE_AWS_REGION}.amazonaws.com/${APP_NAME}:latest \
    --region ${PROVIDE_AWS_REGION}

aws-invoke-function:
	aws lambda invoke --function-name ${APP_NAME} output-lambda.json
	@echo
	@cat output-lambda.json
	@echo
	@echo

terraform-apply:
	cd ${TF_DIR} && terraform apply --auto-approve

terraform-destroy:
	cd ${TF_DIR} && terraform destroy --auto-approve

update-function: docker-build docker-tag docker-push aws-update-function aws-invoke-function

provide-infra: aws-login-ecr aws-create-repo docker-build docker-tag docker-push terraform-apply aws-invoke-function

destroy-infra: aws-login-ecr aws-delete-repo aws-delete-log-group terraform-destroy
