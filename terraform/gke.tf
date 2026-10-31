resource "google_container_cluster" "gke" {
  name = var.cluster_name
  location = var.zone

  remove_default_node_pool = true
  initial_node_count = 1

  deletion_protection = false
}

resource "google_container_node_pool" "primary" {
  name = "primary-node-pool"
  cluster = google_container_cluster.gke.name
  location = var.zone
  node_count = 1

  node_config {
    machine_type = "e2-medium"

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}

data "google_project" "project" {}

resource "google_project_iam_member" "gcr_pull" {
  project = data.google_project.project.project_id
  role    = "roles/storage.objectViewer"
  member  = "serviceAccount:${data.google_project.project.number}-compute@developer.gserviceaccount.com"
}
