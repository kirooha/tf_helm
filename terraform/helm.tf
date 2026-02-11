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
}