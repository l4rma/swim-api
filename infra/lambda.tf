# lambda.tf
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

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "../bootstrap"
  output_path = "lambda_function_payload.zip"
}

resource "aws_lambda_function" "my_lambda" {
  filename      = "lambda_function_payload.zip"
  function_name = "swim-api"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "bootstrap"

  source_code_hash = data.archive_file.lambda.output_base64sha256

  runtime = "provided.al2023"

  environment {
    variables = {
      foo = "bar"
    }
  }
}

locals {
  lambda_src_path = "../cmd/"
  building_path = "./"
  lambda_code_filename = "lambda_function_payload.zip"
}

resource "null_resource" "sam_metadata_aws_lambda_function_my_lambda" {
    triggers = {
        resource_name = "aws_lambda_function.my_lambda"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.lambda_src_path}"
        built_output_path = "./lambda_function_payload.zip"
    }
}

