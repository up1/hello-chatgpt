provider "digitalocean" {
  token = "your_digitalocean_token"
}

resource "digitalocean_droplet" "app" {
  image  = "ubuntu-20-04-x64"
  name   = "app"
  region = "nyc3"
  size   = "s-1vcpu-1gb"

  connection {
    user = "root"
  }

  provisioner "file" {
    source      = "./Dockerfile"
    destination = "/root/Dockerfile"
  }

  provisioner "file" {
    source      = "./docker-compose.yml"
    destination = "/root/docker-compose.yml"
  }

  provisioner "file" {
    source      = "./"
    destination = "/root/app"
  }

  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get install -y docker.io",
      "cd /root",
      "docker-compose build",
      "docker-compose up -d"
    ]
  }
}

resource "digitalocean_droplet" "db" {
  image  = "ubuntu-20-04-x64"
  name   = "db"
  region = "nyc3"
  size   = "s-1vcpu-1gb"

  connection {
    user = "root"
  }

  provisioner "remote-exec" {
    inline = [
      "apt-get update",
      "apt-get install -y docker.io",
      "docker-compose up -d"
    ]
  }