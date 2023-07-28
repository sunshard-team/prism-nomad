job "grafana" {
  datacenters = ["SUNSHARD"]
  namespace = "monitoring"

  group "grafana" {
    count = 1

    volume "grafana" {
      type            = "csi"
      source          = "grafana"
      read_only       = false
      attachment_mode = "file-system"
      access_mode     = "multi-node-multi-writer"
    }

    network {
      port "grafana_ui" {}
    }

    task "grafana" {
      driver = "docker"
      config {
        image = "grafana/grafana:9.3.1"
        ports = ["grafana_ui"]
        
        volumes = [
          "local/datasources:/etc/grafana/provisioning/datasources",
          "local/dashboards:/etc/grafana/provisioning/dashboards",
          // "/home/cluster/storage/grafana/dashboards:/var/lib/grafana/dashboards",
          // "/home/cluster/storage/grafana:/var/lib/grafana",
        ]
      }

      volume_mount {
        volume      = "grafana"
        destination = "/var/lib/grafana"
      }  

      volume_mount {
        volume      = "grafana"
        destination = "/var/lib/grafana/dashboards"
      } 

      env {
        GF_SECURITY_ADMIN_USER     = "admin"
        GF_SECURITY_ADMIN_PASSWORD = "Lqrs15d7"
        GF_SERVER_HTTP_PORT        = "${NOMAD_PORT_grafana_ui}"
        GF_INSTALL_PLUGINS         = "natel-discrete-panel"
      }

      template {
        data = <<EOH
apiVersion: 1
datasources:
- name: Prometheus
  type: prometheus
  access: proxy
  url: http://{{ range $i, $s := service "prometheus" }}{{ if eq $i 0 }}{{.Address}}:{{.Port}}{{end}}{{end}}
  isDefault: true
  version: 1
  editable: false
EOH

        destination = "local/datasources/prometheus.yaml"

      }

      template {
        data = <<EOH
apiVersion: 1
datasources:
- name: Loki
  type: loki
  access: proxy
  url: http://{{ range $i, $s := nomadService "loki" }}{{ if eq $i 0 }}{{.Address}}:{{.Port}}{{end}}{{end}}
  isDefault: false
  version: 1
  editable: false
EOH

        destination = "local/datasources/loki.yaml"
      }

      template {
        data = <<EOH
apiVersion: 1
providers:
- name: PSQL
  folder: PSQL
  folderUid: PSQL
  type: file
  disableDeletion: true
  editable: false
  allowUiUpdates: false
  options:
    path: /var/lib/grafana/dashboards
EOH

        destination = "local/dashboards/nomad-autoscaler.yaml"
      }

      resources {
        cpu    = 1000
        memory = 1000
      }

      service {
        name     = "grafana"
        // provider = "nomad"
        port     = "grafana_ui"
        

        check {
          type     = "http"
          path     = "/api/health"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }
}
