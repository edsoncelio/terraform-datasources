data "http" "myip" {
  url = "https://api.ipify.org"

  # Opcional: header da requisição
  request_headers = {
    Accept = "application/text"
  }
}
n

output "myip" {
    description = "Endereço IP"
    value = data.http.myip.body
}