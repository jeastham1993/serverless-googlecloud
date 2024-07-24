terraform {
  backend "gcs" {
    bucket = "serverless-gym-tf-state-prod"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = "serverless-sandbox-429409"
  region  = "europe-west2"
}
