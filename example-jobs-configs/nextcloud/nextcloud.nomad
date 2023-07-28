job "nextcloud" {
  datacenters = ["SUNSHARD"]
  type = "service"

  group "mariadb" {
    count = 1

    network {
      mode = "bridge"
    }

    update {
      min_healthy_time  = "10s"
      healthy_deadline  = "5m"
      progress_deadline = "10m"
      auto_revert       = true
    }

    // volume "mariadb" {
    //   type = "host"
    //   read_only = false
    //   source = [[ .wordpress.mariadb_group_volume | quote ]]
    // }

    service {
      name = "mariadb-nextcloud"
      port = "3306"
      tags = [ "database" ]

      connect {
        sidecar_service {}
      }
   
    //   check {
    //     name     = "mariadb"
    //     type     = "tcp"
    //     // port     = 3306
    //     interval = "10s"
    //     timeout  = "2s"
    //   }

    }
 

    restart {
      attempts = "2"
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "mariadb" {
      driver = "docker"

      config {
        image = "mariadb:10.6.4-focal"
        // command = "--transaction-isolation=READ-COMMITTED --binlog-format=ROW"
      }
      
    //   volume_mount {
    //     volume      = "mariadb"
    //     destination = "/var/lib/mysql"
    //     read_only   = false
    //   }

      env {
        MYSQL_ROOT_PASSWORD="lqrs15d7"
        MYSQL_DATABASE="nextcloud"
        MYSQL_USER="nextcloud"
        MYSQL_PASSWORD="lqrs15d7"
      }
 

      resources {
        cpu    = 1000
        memory = 1000
      }
    }
  }

  group "nextcloud" {
    count = 1

    network {
      mode = "bridge"
      port "http" {
        to = 80
      }
    }

    update {
        min_healthy_time  = "10s"
        healthy_deadline  = "5m"
        progress_deadline = "10m"
        auto_revert       = true
    }


    service {
      name = "nextcloud"
      port = "http"
      tags = ["app"]

      connect {
        sidecar_service {
          proxy {
            upstreams {
              destination_name = "mariadb-nextcloud"
              local_bind_port  = 3306
            }
          }
        }
      }

    //   check {
    //     name     = "wordpress"
    //     path     = "/wp-admin/install.php"
    //     type     = "http"
    //     port     = "http"
    //     interval = "10s"
    //     timeout  = "2s"
    //   }

    }

    restart {
      attempts = "2"
      interval = "30m"
      delay = "15s"
      mode = "fail"
    }

    task "nextcloud" {
      driver = "docker"

      config {
        image = "nextcloud"
      }

    # For more env visit https://github.com/docker-library/wordpress/pull/142
      env {
        MYSQL_HOST = "${NOMAD_UPSTREAM_ADDR_mariadb-nextcloud}"
        MYSQL_PASSWORD="lqrs15d7"
        MYSQL_DATABASE="nextcloud"
        MYSQL_USER="nextcloud"
        NEXTCLOUD_TRUSTED_DOMAINS="next.sunshard.ru"
      }


      resources {
        cpu    = 3000
        memory = 3000
      }
    }
  }
}
