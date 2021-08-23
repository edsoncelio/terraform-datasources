data "external" "external_data" {
  program = ["go", "run", "${path.module}/external_datasource.go"]

  query = {
    sound = "honk"
  }
}

output "response" {
    value = data.external.external_data.result
}

