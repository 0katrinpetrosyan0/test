terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.19.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
  access_key = "AKIAYHLZZFMQEBGYZI7R"
  secret_key = "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"
}
