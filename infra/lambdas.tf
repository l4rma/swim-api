locals {
  create_lambda_src_path = "../cmd/api/v2/swimmers/create/"
  create_lambda_filename = "create.zip"
  list_lambda_src_path = "../cmd/api/v2/swimmers/list/"
  list_lambda_filename = "list.zip"
  update_lambda_src_path = "../cmd/api/v2/swimmers/update/"
  update_lambda_filename = "update.zip"
  delete_lambda_src_path = "../cmd/api/v2/swimmers/delete/"
  delete_lambda_filename = "delete.zip"
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
  role              = aws_iam_role.iam_for_lambda.arn
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
  role              = aws_iam_role.iam_for_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_list_swimmers" {
    triggers = {
        resource_name = "aws_lambda_function.list_swimmers"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.list_lambda_src_path}"
        built_output_path = "./${local.list_lambda_filename}"
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
  role              = aws_iam_role.iam_for_lambda.arn
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
  role              = aws_iam_role.iam_for_lambda.arn
}

resource "null_resource" "sam_metadata_aws_lambda_function_delete_swimmer" {
    triggers = {
        resource_name = "aws_lambda_function.delete_swimmer"
        resource_type = "ZIP_LAMBDA_FUNCTION"
        original_source_code = "${local.delete_lambda_src_path}"
        built_output_path = "./${local.delete_lambda_filename}"
    }
}
