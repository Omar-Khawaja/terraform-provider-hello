output "pings" {
  value = hello_ping.example_ping[*].name
}

output "pongs" {
  value = hello_pong.example_pong.name
}

output "items" {
  value = hello_append_item.example[*].name
}
