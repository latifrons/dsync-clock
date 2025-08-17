job "dsync-clock" {
  datacenters = ["aws-main"]

  constraint {
    attribute = "${attr.unique.network.ip-address}"
    value     = "10.2.3.12"
  }

  group "dsync-clock" {
    count = 1

    spread {
      attribute = "${unique.hostname}"
      weight    = 100
    }

    ephemeral_disk {
      migrate = false
      size    = 50
      sticky  = false
    }

    network {
      port "tcp-port" {
        static = 8000
        to     = 8080
      }
    }

    restart {
      attempts = 20
      interval = "5m"
      delay    = "15s"
      mode     = "delay"
    }
    shutdown_delay = "15s"

    task "dsync-clock" {
      logs {
        max_files     = 4
        max_file_size = 10
      }

      driver = "docker"

      config {
        network_mode = "host"
        image = " 10.2.3.13:5000/dsync-clock:1.0"
        ports = ["tcp-port"]
      }

      resources {
        cpu        = 100
        memory     = 200
        memory_max = 200
      }
      kill_timeout = "0s"
    }
  }
}




