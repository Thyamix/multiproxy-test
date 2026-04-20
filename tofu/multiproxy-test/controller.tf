# controller.tf

resource "kubernetes_deployment_v1" "controller" {
	metadata {
		name = "controller"
		namespace = "multiproxy-test"
	}
	spec {
		selector {
			match_labels = {
				app = "controller"
			}
		}
		replicas = 5
		template {
			metadata {
				labels = {
					app = "controller"
				}
			}
			spec {
				container {
					name = "controller"
					image = ""
					image_pull_policy = "Always"
					port {
						container_port = 8080
					}
				}
			}
		}
	}
}
