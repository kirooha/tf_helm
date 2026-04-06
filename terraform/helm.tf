resource "helm_release" "ingress_nginx" {
  name = "ingress-nginx"
  namespace = "ingress-nginx"
  create_namespace = true

  repository = "https://kubernetes.github.io/ingress-nginx"
  chart = "ingress-nginx"
  version = "4.12.3"

  set {
    name = "controller.service.type"
    value = "LoadBalancer"
  }
}

resource "helm_release" "kuber_practice_app" {
  name = "kuber-practice-app"
  chart = "../chart"

  values = [
    file("../chart/values.yaml")
  ]

  set {
    name = "postgres.host"
    value = google_sql_database_instance.postgres.private_ip_address
  }

  set {
    name = "postgres.user"
    value = var.db_user
  }

  set_sensitive {
    name = "postgres.password"
    value = var.db_password
  }

  set {
    name = "postgres.db"
    value = var.db_name
  }

  set {
    name = "redis.host"
    value = google_redis_instance.redis.host
  }

  set_sensitive {
    name = "redis.password"
    value = google_redis_instance.redis.auth_string
  }

  set_sensitive {
    name = "api_key"
    value = var.api_key
  }

  set {
    name = "image.digest"
    value = var.app_image_digest
  }

  set {
    name = "cronJobImage.digest"
    value = var.cronjob_image_digest
  }
}