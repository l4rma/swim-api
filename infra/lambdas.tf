locals {
  create_lambda_src_path = "../cmd/api/v2/swimmers/create/"
  create_lambda_filename = "create.zip"
  list_lambda_src_path = "../cmd/api/v2/swimmers/list/"
  list_lambda_filename = "list.zip"
  find_lambda_src_path = "../cmd/api/v2/swimmers/find/"
  find_lambda_filename = "find.zip"
  update_lambda_src_path = "../cmd/api/v2/swimmers/update/"
  update_lambda_filename = "update.zip"
  delete_lambda_src_path = "../cmd/api/v2/swimmers/delete/"
  delete_lambda_filename = "delete.zip"
  create_session_lambda_src_path = "../cmd/api/v2/sessions/create/"
  create_session_lambda_filename = "create_session.zip"
  lambdas_building_path = "../bin"
}

##### Create Swimmer Lambda #####
data "archive_file" "create_swimmer" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/create/bootstrap"
  output_path = "${local.create_lambda_filename}"
}

resource "aws_lambda_function" "create_swimmer" {
  function_name     = "CreateSwimmer"
  description       = "Create a swimmer and add it to the database"
  handler           = "bootstrap"
  filename          = "${local.create_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.create_swimmer.output_base64sha256
  role              = aws_iam_role.create_swimmer_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_create_swimmer" {
    triggers = {
        resource_name = "aws_lambda_function.create_swimmer"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.create_lambda_src_path}"
        built_output_path = "./${local.create_lambda_filename}"
    }
}

##### List Swimmers Lambda #####
data "archive_file" "list_swimmers" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/list/bootstrap"
  output_path = "${local.list_lambda_filename}"
}

resource "aws_lambda_function" "list_swimmers" {
  function_name     = "ListSwimmers"
  description       = "List all swimmers in the database"
  handler           = "bootstrap"
  filename          = "${local.list_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.list_swimmers.output_base64sha256
  role              = aws_iam_role.list_swimmers_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_list_swimmers" {
    triggers = {
        resource_name = "aws_lambda_function.list_swimmers"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.list_lambda_src_path}"
        built_output_path = "./${local.list_lambda_filename}"
    }
}

##### Find Swimmer Lambda #####
data "archive_file" "find_swimmer" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/find/bootstrap"
  output_path = "${local.find_lambda_filename}"
}

resource "aws_lambda_function" "find_swimmer" {
  function_name     = "FindSwimmer"
  description       = "Find a swimmer in the database"
  handler           = "bootstrap"
  filename          = "${local.find_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.find_swimmer.output_base64sha256
  role              = aws_iam_role.list_swimmers_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_find_swimmer" {
    triggers = {
        resource_name = "aws_lambda_function.find_swimmer"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.find_lambda_src_path}"
        built_output_path = "./${local.find_lambda_filename}"
    }
}

##### Update Swimmer Lambda #####
data "archive_file" "update_swimmer" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/update/bootstrap"
  output_path = "${local.update_lambda_filename}"
}

resource "aws_lambda_function" "update_swimmer" {
  function_name     = "UpdateSwimmer"
  description       = "Update a swimmer in the database"
  handler           = "bootstrap"
  filename          = "${local.update_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.update_swimmer.output_base64sha256
  role              = aws_iam_role.update_swimmer_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_update_swimmer" {
    triggers = {
        resource_name = "aws_lambda_function.update_swimmer"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.update_lambda_src_path}"
        built_output_path = "./${local.update_lambda_filename}"
    }
}

##### Delete Swimmer Lambda #####
data "archive_file" "delete_swimmer" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/delete/bootstrap"
  output_path = "${local.delete_lambda_filename}"
}

resource "aws_lambda_function" "delete_swimmer" {
  function_name     = "DeleteSwimmer"
  description       = "Delete a swimmer from the database"
  handler           = "bootstrap"
  filename          = "${local.delete_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.delete_swimmer.output_base64sha256
  role              = aws_iam_role.delete_swimmer_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_delete_swimmer" {
    triggers = {
        resource_name = "aws_lambda_function.delete_swimmer"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.delete_lambda_src_path}"
        built_output_path = "./${local.delete_lambda_filename}"
    }
}

##### Create Session Lambda #####
data "archive_file" "create_session" {
  type        = "zip"
  source_file = "${local.lambdas_building_path}/sessions/create/bootstrap"
  output_path = "${local.create_session_lambda_filename}"
}

resource "aws_lambda_function" "create_session" {
  function_name     = "CreateSession"
  description       = "Create a session and add it to the database"
  handler           = "bootstrap"
  filename          = "${local.create_session_lambda_filename}"
  runtime           = "provided.al2023"
  source_code_hash  = data.archive_file.create_session.output_base64sha256
  role              = aws_iam_role.create_swimmer_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_create_session" {
    triggers = {
        resource_name = "aws_lambda_function.create_session"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.create_session_lambda_src_path}"
        built_output_path = "./${local.create_session_lambda_filename}"
    }
}

