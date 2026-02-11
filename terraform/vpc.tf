data "google_compute_network" "default" {
  name = "default"
}

resource "google_project_service" "servicenetworking" {
  service = "servicenetworking.googleapis.com"
}

resource "google_compute_global_address" "cloudsql_private_range" {
  name = "cloudsql-private-range"
  purpose = "VPC_PEERING"
  address_type = "INTERNAL"
  prefix_length = 16
  network = data.google_compute_network.default.id
}

resource "google_service_networking_connection" "cloudsql_vpc_connection" {
  network = data.google_compute_network.default.id
  service = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [
    google_compute_global_address.cloudsql_private_range.name
  ]

  depends_on = [google_project_service.servicenetworking]
}