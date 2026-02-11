resource "helm_release" "kuber_practice_app" {
  name = "kuber-practice-app"
  chart = "../chart"

  values = [
    file("../chart/values.yaml")
  ]
}