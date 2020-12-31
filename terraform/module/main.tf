//ENV CGO_ENABLED=0
//RUN

resource "null_resource" "cmd" {
  // command = "${path.module}/appsettings.sh ${azuread_application.rbac-server-principal.application_id} ${azuread_application.rbac-client-principal.application_id}"
}

data "archive_file" "zip" {
  type        = "zip"
  source_file = "bin/aws-lambda-go"
  output_path = "aws-lambda-go.zip"

  provisioner "local-exec" {
    command = "cat ${path.module}/dist"
  }
}