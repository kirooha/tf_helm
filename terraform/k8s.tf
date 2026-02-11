data "google_client_config" "default" {}

provider "kubernetes" {
  host = google_contrainer_cluster.gke.endpoint
  token = data.google_client_config.default.access_token
  cluster_ca_certificate = base64decode(
    google_contrainer_cluster.gke.master_auth[0].cluster_ca_certificate
  )
}