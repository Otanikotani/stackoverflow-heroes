terraform {
  backend "s3" {
    bucket = "otanikotani-tf"
    key = "stackoverflow-heroes.tfstate"
    region = "us-east-1"
    encrypt = true
  }
}

provider "aws" {
  region  = "us-east-1"
}

module "fetch" {
  source = "./fetch"
  stack_exchange_access_token = var.stack_exchange_access_token
  stack_exchange_key = var.stack_exchange_key
}