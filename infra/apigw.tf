# API Gateway
resource "aws_api_gateway_rest_api" "api_gw" {
  name        = "swim-api"
  description = "An API to record swimming sessions"
}

# /swimmers
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
  uri                     = aws_lambda_function.my_lambda.invoke_arn
}


# /swimmers/add
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
  uri                     = aws_lambda_function.my_lambda.invoke_arn
}

# /swimmers/find
resource "aws_api_gateway_resource" "find_swimmer" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_resource.swimmers.id
  path_part   = "find"
}


resource "aws_api_gateway_method" "find_swimmer" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.find_swimmer.id
  http_method   = "GET"
  authorization = "NONE"

  request_parameters = {
        "method.request.querystring.id" = true
      }
}

resource "aws_api_gateway_integration" "find_swimmer" {
  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
  resource_id             = aws_api_gateway_resource.find_swimmer.id
  http_method             = aws_api_gateway_method.find_swimmer.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.my_lambda.invoke_arn
}

# /sessions
resource "aws_api_gateway_resource" "sessions" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_rest_api.api_gw.root_resource_id
  path_part   = "sessions"
}

# /sessions/add
resource "aws_api_gateway_resource" "add_session" {
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
  parent_id   = aws_api_gateway_resource.sessions.id
  path_part   = "add"
}

resource "aws_api_gateway_method" "add_session" {
  rest_api_id   = aws_api_gateway_rest_api.api_gw.id
  resource_id   = aws_api_gateway_resource.add_session.id
  http_method   = "POST"
  authorization = "NONE"
}
 
resource "aws_api_gateway_integration" "add_session" {
  rest_api_id             = aws_api_gateway_rest_api.api_gw.id
  resource_id             = aws_api_gateway_resource.add_session.id
  http_method             = aws_api_gateway_method.add_session.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.my_lambda.invoke_arn
}


# Lambda permission
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.my_lambda.arn}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api_gw.execution_arn}/*/*/*"
}

resource "aws_api_gateway_deployment" "example_deployment" {
  depends_on  = [aws_api_gateway_integration.add_swimmer, aws_api_gateway_integration.find_swimmer]
  rest_api_id = aws_api_gateway_rest_api.api_gw.id
}

output "api_endpoint" {
  value = aws_api_gateway_deployment.example_deployment.invoke_url
}

