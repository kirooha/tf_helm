resource "helm_release" "argocd" {
  name = "argocd"
  namespace = "argocd"
  create_namespace = true

  repository = "https://argoproj.github.io/argo-helm"
  chart = "argo-cd"
  version = "7.7.0"

  values = [yamlencode({
    dex = { enabled = false }
    notifications = { enabled = false }
    applicationSet = { replicas = 0 }
  })]
}

resource "helm_release" "argocd_apps" {
  name  = "argocd-apps"
  namespace = "argocd"

  repository = "https://argoproj.github.io/argo-helm"
  chart = "argocd-apps"
  version = "2.0.2"

  values = [yamlencode({
    applications = {
      kuber-practice-app = {
        namespace = "argocd"
        project = "default"
        source = {
          repoURL = "https://github.com/kirooha/tf_helm.git"
          targetRevision = "main"
          path = "chart"
          helm = {
            valuesObject = {
              postgres = { host = google_sql_database_instance.postgres.private_ip_address }
              redis = { host = "${google_redis_instance.redis.host}:${google_redis_instance.redis.port}" }
            }
          }
        }
        destination = {
          server = "https://kubernetes.default.svc"
          namespace = "default"
        }
        syncPolicy = {
          automated = { prune = true, selfHeal = true }
        }
      }
    }
  })]

  depends_on = [helm_release.argocd]
}