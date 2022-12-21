resource "aws_apigatewayv2_api" "poc_web_socket_api" {
  name                       = "poc_web_soacket_api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}
