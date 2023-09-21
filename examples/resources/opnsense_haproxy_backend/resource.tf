resource "opnsense_haproxy_backend" "example" {
  enabled = true
  description = "K8s API Backend"

  name = "test"
}
