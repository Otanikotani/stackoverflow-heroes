resource "aws_lambda_function" "so_fetch" {
  function_name = "so_fetch"
  filename = "build/fetch.zip"
  handler = "fetch"
  source_code_hash = filebase64sha256("build/fetch.zip")
  role = aws_iam_role.so_fetch_role.arn
  runtime = "go1.x"
  memory_size = 128
  timeout = 15

  tags = {
    Project = "Stackoverflow Heroes"
  }

  environment {
    variables = {
      BUCKET = aws_s3_bucket.fetch_bucket.bucket
      STACK_EXCHANGE_ACCESS_TOKEN = var.stack_exchange_access_token
      STACK_EXCHANGE_KEY = var.stack_exchange_key
    }
  }
}

resource "aws_s3_bucket" "fetch_bucket" {
  bucket = "stackoverflow-heroes-fetch"
  acl = "private"
  force_destroy = true //ensure we can destroy the bucket even if it has content

  tags = {
    Project = "Stackoverflow Heroes"
  }
}

resource "aws_iam_role" "so_fetch_role" {
  name = "so_fetch_role"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY

  tags = {
    Project = "Stackoverflow Heroes"
  }
}

data "aws_iam_policy_document" "so_fetch_policy_document" {
  statement {
    actions = [
      "s3:*",
    ]

    resources = [
      "*"
    ]
  }

  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:CreateLogGroup",
      "logs:PutLogEvents"
    ]

    resources = [
      "arn:aws:logs:*:*:*"
    ]
  }
}

resource "aws_iam_policy" "so_fetch_policy" {
  name = "so_fetch_policy"
  path = "/"
  description = "IAM policy for accessing s3 from a lambda"
  policy = data.aws_iam_policy_document.so_fetch_policy_document.json
}

resource "aws_iam_role_policy_attachment" "attach_policy_to_lambda" {
  role = aws_iam_role.so_fetch_role.name
  policy_arn = aws_iam_policy.so_fetch_policy.arn
}
