resource "opnsense_haproxy_server" "example" {
  enabled = true
  description = "K8s API Server"

  name = "test"
}
