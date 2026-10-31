data "google_client_config" "default" {}

provider "kubernetes" {
  host = google_container_cluster.gke.endpoint
  token = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(
    google_container_cluster.gke.master_auth[0].cluster_ca_certificate
  )
}

resource "kubernetes_secret" "app" {
  metadata {
    name = "kuber-practice-app-secrets"
    namespace = "default"
  }

  data = {
    postgres_password = random_password.db.result
    redis_password = google_redis_instance.redis.auth_string
    api_key = var.api_key
  }
}