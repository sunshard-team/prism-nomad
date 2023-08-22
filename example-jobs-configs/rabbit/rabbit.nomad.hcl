job "rabbit" {

  datacenters = ["Binta"]
  type = "service"

  group "rabbitmq" {
    count = 1

    network {
      mode = "bridge"

      port "rabbitmq_amqp" {
        static = 5672
        to = 5672
      }

      port "rabbit_ui" {
        to = 15672
      }
    }

    task "rabbit" {
      driver = "docker"

      config {
        image = "rabbitmq:3-management"
      }

      env {
        RABBITMQ_DEFAULT_USER = "test"
        RABBITMQ_DEFAULT_PASS = "test"
      }

      service {
        name = "rabbit-ui"
        port = "rabbit_ui"
      }

      service {
        name = "rabbitmq-amqp"
        port = "rabbitmq_amqp"
      }

      resources {
        cpu = 1000
        memory = 1000
      }
    }
  }
}