resource "opnsense_haproxy_frontend" "example" {
  enabled = true
  description = "K8s API Frontend"
  bind = "192.168.100.10:8080"
  name = "test"
}
