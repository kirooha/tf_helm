terraform {
  backend "gcs" {
    bucket = "my-new-project-467616-tfstate"
    prefix = "terraform/state"
  }
}