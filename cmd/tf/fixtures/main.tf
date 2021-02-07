terraform {
  required_version = "= 0.14.4"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" new_server {
  instance_type = "t2.micro"
}
