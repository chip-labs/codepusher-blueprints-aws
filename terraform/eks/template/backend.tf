provider "aws" {
  region = "{{ .region }}"
  assume_role {
    role_arn = "{{ .roleArn }}"
  }
}

terraform {
  backend "s3" {
    bucket         = "{{ .stateBucket }}"
    key            = "aws/environment/{{ .environmentName }}/eks/{{ .clusterName }}/terraform.tfstate"
    region         = "{{ .region }}"
    encrypt        = true
    dynamodb_table = "{{ .stateTable }}"
    role_arn       = "{{ .roleArn }}"
  }
}