##### LAMBDA #####
data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }

}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "iam_for_lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "lambda_policies" {
  statement {
    effect = "Allow"
    actions = [
       "logs:CreateLogGroup",
       "logs:CreateLogStream",
       "logs:PutLogEvents"
    ]
    resources = ["*"]
  }

  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:GetItem",
      "dynamodb:UpdateItem",
      "dynamodb:Scan"
    ]
    resources = [
      aws_dynamodb_table.swimmers_and_sessions.arn
    ]
  }

  depends_on = [
    aws_dynamodb_table.swimmers_and_sessions
  ]
}

resource "aws_iam_policy" "log_policy" {
  name        = "Log-policy"
  description = "Policy to log with cloudwatch"
  policy      = data.aws_iam_policy_document.lambda_policies.json
}

resource "aws_iam_role_policy_attachment" "policy_attachment" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.log_policy.arn
}

##### API GATEWAY #####
resource "aws_lambda_permission" "api_gateway" {
  for_each = {
    create = aws_lambda_function.create_swimmer.arn
    read   = aws_lambda_function.list_swimmers.arn
    update   = aws_lambda_function.update_swimmer.arn
  }

  statement_id  = "AllowAPIGatewayInvoke-${each.key}"
  action        = "lambda:InvokeFunction"
  function_name = each.value
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api_gw.execution_arn}/*/*/*"
}

