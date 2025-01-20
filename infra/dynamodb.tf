resource "aws_dynamodb_table" "swimmers_and_sessions" {
  name           = "SwimmersAndSessions"
  billing_mode   = "PROVISIONED"
  hash_key       = "PK"
  range_key      = "SK"
  read_capacity  = 5
  write_capacity = 5

  attribute {
    name = "PK"
    type = "S" # String type
  }

  attribute {
    name = "SK"
    type = "S" # String type
  }

  global_secondary_index {
    name               = "GSI1"
    hash_key           = "PK"
    range_key          = "Date"
    projection_type    = "ALL"
    read_capacity      = 5
    write_capacity     = 5
  }

  attribute {
    name = "Date"
    type = "S" # String type for ISO 8601 date format
  }

  tags = {
    Environment = "dev"
    Project     = "swim-api"
  }
}

