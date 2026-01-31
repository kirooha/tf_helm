resource "helm_release" "kuber_practice_app" {
  name = "kuber-practice-app"
  chart = "../app/chart"

  values = [
    file("../app/chart/values.yaml")
  ]
}