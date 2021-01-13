resource "aws_apigatewayv2_api" "api" {
  name          = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
  protocol_type = "HTTP"

  description = "This is my API for demonstration purposes"
}

resource "aws_apigatewayv2_stage" "stage" {
  depends_on = [aws_apigatewayv2_integration.integration]
  name   = "$default"
  api_id = aws_apigatewayv2_api.api.id

  auto_deploy = true
}

resource "aws_apigatewayv2_deployment" "deployment" {
  api_id      = aws_apigatewayv2_api.api.id
  description = "deployment"

  triggers = {
    redeployment = sha1(
      join(
        ",",
        list(
          jsonencode(aws_apigatewayv2_integration.integration),
          jsonencode(aws_apigatewayv2_route.route),
        )
      )
    )
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_apigatewayv2_integration" "integration" {
  api_id             = aws_apigatewayv2_api.api.id
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
  integration_uri    = aws_lambda_function.lambda.invoke_arn

  connection_type    = "INTERNET"
  payload_format_version = "2.0"

  description        = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
}

resource "aws_apigatewayv2_api_mapping" "api_mapping" {
  api_id      = aws_apigatewayv2_api.api.id
  domain_name = aws_apigatewayv2_domain_name.domain.domain_name
  stage       = aws_apigatewayv2_stage.stage.id
}

resource "aws_apigatewayv2_domain_name" "domain" {
  domain_name = var.record.name

  domain_name_configuration {
    certificate_arn = var.record.acm_arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }
}

resource "aws_apigatewayv2_route" "route" {
  api_id             = aws_apigatewayv2_api.api.id
  authorization_type = "NONE"
  route_key          = "ANY /${local.proxy}"
  operation_name     = "APIServer"
  target             = "integrations/${aws_apigatewayv2_integration.integration.id}"
}

resource "aws_route53_record" "record" {
  name    = var.record.name
  type    = "A"
  zone_id = var.record.zone_id
  alias {
    evaluate_target_health = true
    name                   = aws_apigatewayv2_domain_name.domain.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.domain.domain_name_configuration[0].hosted_zone_id
  }
}

locals {
  proxy = "{proxy+}"
}

output "endpoint" {
  value = "https://${aws_route53_record.record.name}"
}

