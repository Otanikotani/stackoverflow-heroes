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
  stack_exchange_access_token = "SET-FROM-ENV"
  stack_exchange_key = "SET-FROM-ENV"
}