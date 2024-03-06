provider "aws" {
  region = local.PROVIDE_AWS_REGION
}

#####################################################################
## LAMBDA FUNCTION
#####################################################################

resource "aws_lambda_function" "lambda" {
  function_name = local.APP_NAME
  timeout = local.LAMBDA_TIMEOUT
  memory_size = local.LAMBDA_MEMORY_SIZE
  package_type = "Image"
  image_uri = "${local.PROVIDE_AWS_ACCOUNT}.dkr.ecr.${local.PROVIDE_AWS_REGION}.amazonaws.com/${local.APP_NAME}:latest"
  role          = aws_iam_role.role.arn

  environment {
    variables = {
      APP_VERSION=local.APP_VERSION
      APP_LOG_LEVEL=local.APP_LOG_LEVEL
      PROVIDE_AWS_REGION=local.PROVIDE_AWS_REGION
    }
  }
}

#####################################################################
## ASSUME ROLE
#####################################################################

resource "aws_iam_role" "role" {
  name = "${local.APP_NAME}-AssumeRole"
  assume_role_policy = file("role-assume-role.json")
}

#####################################################################
## CLOUD WATCH
#####################################################################

resource "aws_iam_policy" "cloud-watch-policy" {
  name = "${local.APP_NAME}-CloudWatchLogsPolicy"
  policy = file("policy-cloud-watch.json")
}
resource "aws_iam_policy_attachment" "cloud-watch-attachment" {
  name = "${local.APP_NAME}-CloudWatchLogsAttachment"
  roles = [aws_iam_role.role.name]
  policy_arn = aws_iam_policy.cloud-watch-policy.arn
}

#####################################################################
## DATABASE
#####################################################################

resource "aws_iam_policy" "dynamodb-policy" {
  name = "${local.APP_NAME}-DynamoDBPolicy"
  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "dynamodb:BatchGetItem",
          "dynamodb:BatchWriteItem",
          "dynamodb:PutItem",
          "dynamodb:GetItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem"
        ],
        "Resource": aws_dynamodb_table.tb-companies.arn
      }
    ]
  })
}
resource "aws_iam_policy_attachment" "dynamodb-attachment" {
  name = "dynamodb-lambda-Attachment"
  roles = [aws_iam_role.role.name]
  policy_arn = aws_iam_policy.dynamodb-policy.arn
}

#############
### TABLES
#############

resource "aws_dynamodb_table" "tb-companies" {
  name = "${local.APP_NAME}-${local.TABLE_COMPANIES}"
  billing_mode = "PROVISIONED"
  read_capacity = 1
  write_capacity = 1
  hash_key = "ID"

  server_side_encryption {
    enabled = false
  }

  attribute {
    name = "ID"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "tb-companies-items" {
  table_name = "${local.APP_NAME}-${local.TABLE_COMPANIES}"
  hash_key = aws_dynamodb_table.tb-companies.hash_key
  item = jsonencode(
    {
      "ID" : { "S" : "1" },
      "Name" : { "S" : "Tata Consultancy Service" },
      "State" : { "S" : "SP" }
    }
  )
}

#####################################################################
## API GATEWAY
#####################################################################

resource "aws_api_gateway_rest_api" "api" {
  name = "${local.APP_NAME}-API"
}

resource "aws_api_gateway_resource" "api" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  parent_id = aws_api_gateway_rest_api.api.root_resource_id
  path_part = "{proxy+}"
}

resource "aws_api_gateway_method" "api" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_resource.api.id
  http_method = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "api" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_method.api.resource_id
  http_method = aws_api_gateway_method.api.http_method

  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = aws_lambda_function.lambda.invoke_arn
}

resource "aws_api_gateway_deployment" "v1" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  depends_on = [aws_api_gateway_integration.api]
  stage_name = local.API_VERSION
}

# Lambda permission
resource "aws_lambda_permission" "api-lambda" {
  function_name = aws_lambda_function.lambda.arn
  statement_id = "allow-${aws_lambda_function.lambda.function_name}-invoke"
  action = "lambda:InvokeFunction"
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_deployment.v1.execution_arn}/*/*"
}
