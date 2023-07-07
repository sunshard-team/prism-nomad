job "nginx" {
  datacenters = ["dc1"]
  type        = "смотреть prism.yaml"
	namespace   = "default"
  meta {
    run_uuid = "${uuidv4()}" //подставить deploy_version
  }
  group "nginx" {
    count = 1

    network {
      port "http" {
        static = 8080
      }
    }

    scaling {
			enabled = true
			min     = 1
			max     = 6
		}

    service {
      name = "nginx"
      port = "http"
      tags = ["nginx"]
      // CUSTOM
      // connect {
			// 	sidecar_service {
			// 		proxy {
			// 			upstreams {
			// 				destination_name = "mysql-server"
			// 				local_bind_port  = 3306
			// 			}
			// 		}
			// 	}
			// }
    }

    restart {
			attempts = 10
			interval = "5m"
			delay    = "25s"
			mode     = "delay"
		}

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx"
        force_pull = true 

        ports = ["http"]

        volumes = [
          "local:/etc/nginx/conf.d",
          // OTHER VOLUMES
          // "local/test.conf:/etc/nginx/conf.d/test.conf",
          // "custom_path:/custom_path_in_container",
        ]
      }

      template {
        data = <<EOF
upstream backend {
{{ range service "$ENV_NAME_SERVICE" }}
  server {{ .Address }}:{{ .Port }};
{{ else }}server 127.0.0.1:65535; # force a 502
{{ end }}
}

server {
   listen 8080;

   location / {
      proxy_pass http://backend;
   }
}
EOF

        destination   = "local/load-balancer.conf"
        change_mode   = "signal"
        change_signal = "SIGHUP"
      }
//      CUSTOM
//       template {
//         data = <<EOF
// test
// EOF

//         destination   = "local/test.conf"
//         change_mode   = "signal"
//         change_signal = "SIGHUP"
//       }

      // CUSTOM    
      // env = {
			// 	WORDPRESS_DB_HOST = "${NOMAD_UPSTREAM_IP_mysql-server}"
			// 	WORDPRESS_CONFIG_EXTRA = "define('WP_HOME','wp.sunshard.ru:443')"
			// 	WORDPRESS_CONFIG_EXTRA = "define('WP_ALLOW_REPAIR', true );\ndefine('WP_HOME','https://wp.sunshard.space');\ndefine('WP_SITEURL','https://wp.sunshard.space');"
			// }

      resources {
					cpu    = 1000
					memory = 1000

			}
    }
  }
}
