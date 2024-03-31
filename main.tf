provider "aws" {
  region = "eu-west-1"
}
resource "aws_sqs_queue" "LinksQueue" {
  name                        = "LinksQueue.fifo"
  message_retention_seconds   = 3600
  fifo_queue                  = true
  content_based_deduplication = true
}


resource "aws_dynamodb_table" "Pages" {
  name         = "Pages"
  billing_mode = "PAY_PER_REQUEST" # On-demand mode
  hash_key     = "PageURL"

  attribute {
    name = "PageURL"
    type = "S"
  }
}
resource "aws_dynamodb_table" "Website" {
  name         = "Websites"
  billing_mode = "PAY_PER_REQUEST" # On-demand mode
  hash_key     = "BaseURL"

  attribute {
    name = "BaseURL"
    type = "S"
  }
}