variable "controller_count" {
  type    = number
  default = 3
  validation {
    condition     = var.controller_count >= 1
    error_message = "Must be 1 or more."
  }
}

locals {
  controller_nodes = [
    for i in range(var.controller_count) : {
      name    = "k8controller-${1 + i}"
      description = "This is for the K8s controller ${1 + i}"
      ip = "192.168.24.${11 + i}"
    }
  ]
}

resource "opnsense_haproxy_server" "controller" {
  count   = var.controller_count
  enabled = true
  name = local.controller_nodes[count.index].name
  description = local.controller_nodes[count.index].description
  address = local.controller_nodes[count.index].ip
  port = 6443
  //resolver_opts = "allow-dup-ip"
}

resource "opnsense_haproxy_backend" "controller" {
  enabled = true
  description = "K8s API Backend"
  linked_servers = [for i in resource.opnsense_haproxy_server.controller : i.id]
  name = "k8-controller-backend"
  //resolver_opts = ["allow-dup-ip", "ignore-weight"]
  depends_on = [opnsense_haproxy_server.controller]
}

resource "opnsense_haproxy_frontend" "controller" {
  enabled = true
  description = "K8s API Frontend"
  bind = [
    "192.168.24.1:6443",
    "192.168.2.1:6443"
  ]
  name = "k8-controller-frontend"
  default_backend = resource.opnsense_haproxy_backend.controller.id

  depends_on = [opnsense_haproxy_backend.controller]
}
