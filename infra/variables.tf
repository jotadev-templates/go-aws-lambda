locals {
  APP_NAME = "app-crm-lambda"
  APP_VERSION = "v1.0.0"
  APP_LOG_LEVEL= "DEBUG"
}

locals{
  PROVIDE_AWS_ACCOUNT = ""
  PROVIDE_AWS_REGION = "sa-east-1"
}

locals {
  LAMBDA_TIMEOUT = 3
  LAMBDA_MEMORY_SIZE = 128
}

locals {
  TABLE_COMPANIES = "tb_companies"
}

locals{
  API_VERSION = "v1"
}
