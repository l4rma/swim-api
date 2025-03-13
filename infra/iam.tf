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

resource "aws_iam_role" "create_swimmer_lambda" {
  name               = "CreateSwimmerLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role" "list_swimmers_lambda" {
  name               = "ListSwimmersLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role" "update_swimmer_lambda" {
  name               = "UpdateSwimmerLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role" "delete_swimmer_lambda" {
  name               = "DeleteSwimmerLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "ddb_write" {
  statement {
    sid = "AllowLambdaToWriteToDynamoDB"
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:UpdateItem"
    ]
    resources = [
      aws_dynamodb_table.swimmers_and_sessions.arn
    ]
  }
}

data "aws_iam_policy_document" "ddb_read" {
  statement {
    sid = "AllowLambdaToReadFromDynamoDB"
    effect = "Allow"
    actions = [
      "dynamodb:GetItem",
      "dynamodb:Scan"
    ]
    resources = [
      aws_dynamodb_table.swimmers_and_sessions.arn
    ]
  }
}

data "aws_iam_policy_document" "ddb_delete" {
  statement {
    sid = "AllowLambdaToDeleteFromDynamoDB"
    effect = "Allow"
    actions = [
      "dynamodb:DeleteItem"
    ]
    resources = [
      aws_dynamodb_table.swimmers_and_sessions.arn
    ]
  }
}

data "aws_iam_policy_document" "logs" {
  statement {
    sid = "AllowLambdaToLogToCloudWatch"
    effect = "Allow"
    actions = [
       "logs:CreateLogGroup",
       "logs:CreateLogStream",
       "logs:PutLogEvents"
    ]
    resources = ["*"]
  }
}

# Policies
resource "aws_iam_policy" "log_policy" {
  name        = "Log-policy"
  description = "Policy to log with cloudwatch"
  policy      = data.aws_iam_policy_document.logs.json
}

resource "aws_iam_policy" "write_policy" {
  name        = "Write-policy"
  description = "Policy to write to DynamoDB"
  policy      = data.aws_iam_policy_document.ddb_write.json
}

resource "aws_iam_policy" "read_policy" {
  name        = "Read-policy"
  description = "Policy to read from DynamoDB"
  policy      = data.aws_iam_policy_document.ddb_read.json
}

resource "aws_iam_policy" "delete_policy" {
  name        = "Delete-policy"
  description = "Policy to delete from DynamoDB"
  policy      = data.aws_iam_policy_document.ddb_delete.json
}

# Add the policy to the role
locals {
  log_policy_arn    = aws_iam_policy.log_policy.arn
  write_policy_arn  = aws_iam_policy.write_policy.arn
  read_policy_arn   = aws_iam_policy.read_policy.arn
  delete_policy_arn = aws_iam_policy.delete_policy.arn
}

# Attach policies to create_swimmer role
resource "aws_iam_role_policy_attachment" "create_swimmer_role" {
  for_each = {
    log   = local.log_policy_arn
    write = local.write_policy_arn
  }
  role       = aws_iam_role.create_swimmer_lambda.name
  policy_arn = each.value
}

# Attach policies to list_swimmers role
resource "aws_iam_role_policy_attachment" "list_swimmers_role" {
  for_each = {
    log  = local.log_policy_arn
    read = local.read_policy_arn
  }
  role       = aws_iam_role.list_swimmers_lambda.name
  policy_arn = each.value
}

# Attach policies to update_swimmer role
resource "aws_iam_role_policy_attachment" "update_swimmer_role" {
  for_each = {
    log   = local.log_policy_arn
    write = local.write_policy_arn
    read  = local.read_policy_arn
  }
  role       = aws_iam_role.update_swimmer_lambda.name
  policy_arn = each.value
}

# Attach policies to delete_swimmer role
resource "aws_iam_role_policy_attachment" "delete_swimmer_role" {
  for_each = {
    log    = local.log_policy_arn
    delete = local.delete_policy_arn
  }
  role       = aws_iam_role.delete_swimmer_lambda.name
  policy_arn = each.value
}

##### API GATEWAY #####
resource "aws_lambda_permission" "api_gateway" {
  for_each = {
    create = aws_lambda_function.create_swimmer.arn
    read   = aws_lambda_function.list_swimmers.arn
    update   = aws_lambda_function.update_swimmer.arn
    delete   = aws_lambda_function.delete_swimmer.arn
  }

  statement_id  = "AllowAPIGatewayInvoke-${each.key}"
  action        = "lambda:InvokeFunction"
  function_name = each.value
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.api_gw.execution_arn}/*/*/*"
}
