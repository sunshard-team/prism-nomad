job "example" {
  datacenters = ["dc1"]
  type        = "service"
  service {
    name = "example"
    port = "http"
    check {
      name     = "alive"
      type     = "tcp"
      interval = "10s"
      timeout  = "2s"
    }
  }
  restart {
    attempts = 2
    interval = "30m"
    delay    = "15s"
    mode     = "fail"
  }
  ephemeral_disk {
    sticky  = true
    migrate = true
    size    = 300
  }

  group "group1" {
    count = 1
    task "task1" {
      driver = "docker"
      config {
        image = "redis:7"
        args = ["-bind","${NOMAD_PORT_http}","${nomad.datacenter}","${MY_ENV}","${meta.foo}"]
        resources {
          cpu    = 500
          memory = 256
        }
      }
      command = "my-command"
      ports = "db"
      volumes = ["/path/on/host:/path/in/container","relative/to/task:/also/in/container"]
      auth_soft_fail = true
    }
  }
}