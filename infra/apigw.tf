# API Gateway
resource "aws_api_gateway_rest_api" "api_gw" {
  name        = "swim-api"
  description = "An API to record swimming sessions"
}

# Path: /swimmers 
# Description: List all swimmers in database
resource "aws_api_gateway_resource" "swimmers" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_rest_api.api_gw.root_resource_id
  path_part   = "swimmers"
}

resource "aws_api_gateway_method" "swimmers" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.swimmers.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "swimmers" {
  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
  resource_id             = aws_api_gateway_resource.swimmers.id
  http_method             = aws_api_gateway_method.swimmers.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.list_swimmers.invoke_arn
}


# Path: /swimmers/add
# Description: Add a swimmer to the database
resource "aws_api_gateway_resource" "add_swimmer" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_resource.swimmers.id
  path_part   = "add"
}

resource "aws_api_gateway_method" "add_swimmer" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.add_swimmer.id
  http_method   = "POST"
  authorization = "NONE"
}
 
resource "aws_api_gateway_integration" "add_swimmer" {
  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
  resource_id             = aws_api_gateway_resource.add_swimmer.id
  http_method             = aws_api_gateway_method.add_swimmer.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.create_swimmer.invoke_arn
}

# Path: /swimmers/update
# Description: Update a swimmer in the database
resource "aws_api_gateway_resource" "update_swimmer" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_resource.swimmers.id
  path_part   = "update"
}

resource "aws_api_gateway_method" "update_swimmer" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.update_swimmer.id
  http_method   = "PUT"
  authorization = "NONE"
}
 
resource "aws_api_gateway_integration" "update_swimmer" {
  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
  resource_id             = aws_api_gateway_resource.update_swimmer.id
  http_method             = aws_api_gateway_method.update_swimmer.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.update_swimmer.invoke_arn
}

# /swimmers/find
#resource "aws_api_gateway_resource" "find_swimmer" {
#  rest_api_id = aws_api_gateway_rest_api.api_gw.id
#  parent_id   = aws_api_gateway_resource.swimmers.id
#  path_part   = "find"
#}
#
#
#resource "aws_api_gateway_method" "find_swimmer" {
#  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
#  resource_id   = aws_api_gateway_resource.find_swimmer.id
#  http_method   = "GET"
#  authorization = "NONE"
#
#  request_parameters = {
#        "method.request.querystring.id" = true
#      }
#}
#
#resource "aws_api_gateway_integration" "find_swimmer" {
#  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
#  resource_id             = aws_api_gateway_resource.find_swimmer.id
#  http_method             = aws_api_gateway_method.find_swimmer.http_method
#  integration_http_method = "POST"
#  type                    = "AWS_PROXY"
#  uri                     = aws_lambda_function.my_lambda.invoke_arn
#}

# /sessions
#resource "aws_api_gateway_resource" "sessions" {
#  rest_api_id = aws_api_gateway_rest_api.api_gw.id
#  parent_id   = aws_api_gateway_rest_api.api_gw.root_resource_id
#  path_part   = "sessions"
#}

# /sessions/add
#resource "aws_api_gateway_resource" "add_session" {
#  rest_api_id = aws_api_gateway_rest_api.api_gw.id
#  parent_id   = aws_api_gateway_resource.sessions.id
#  path_part   = "add"
#}
#
#resource "aws_api_gateway_method" "add_session" {
#  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
#  resource_id   = aws_api_gateway_resource.add_session.id
#  http_method   = "POST"
#  authorization = "NONE"
#}
# 
#resource "aws_api_gateway_integration" "add_session" {
#  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
#  resource_id             = aws_api_gateway_resource.add_session.id
#  http_method             = aws_api_gateway_method.add_session.http_method
#  integration_http_method = "POST"
#  type                    = "AWS_PROXY"
#  uri                     = aws_lambda_function.my_lambda.invoke_arn
#}


# API Gateway Deployment
resource "aws_api_gateway_deployment" "dev" {
  depends_on  = [
    aws_api_gateway_integration.add_swimmer,
    aws_api_gateway_integration.swimmers,
    aws_api_gateway_integration.update_swimmer
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
}

resource "aws_api_gateway_stage" "dev" {
  deployment_id = aws_api_gateway_deployment.dev.id
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  stage_name    = "swim-api"
  depends_on    = [aws_api_gateway_deployment.dev]
}

resource "aws_api_gateway_method_settings" "dev" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  stage_name  = aws_api_gateway_stage.dev.stage_name
  method_path = "*/*"

  settings {
    metrics_enabled = true
  }
}
