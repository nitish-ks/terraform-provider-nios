// Create Record AAAA with Basic Fields
resource "nios_dns_record_aaaa" "record1" {
  name     = "example_record.example.com"
  ipv6addr = "2002:1111::1401"
  view     = "default"
}

// Create Record AAAA with additional fields
resource "nios_dns_record_aaaa" "record2" {
  name     = "example_record_with_ttl.example.com"
  ipv6addr = "2002:1111::1401"
  view     = "default"
  use_ttl  = true
  ttl      = 10
  comment  = "Example AAAA record"
  extattrs = {
    Site = "Siteblr"
  }
}

// Create Record AAAA using function call to retrieve ipv6addr
resource "nios_dns_record_aaaa" "record3" {
  name = "example_record_with_func_call.example.com"
  func_call = {
    attribute_name  = "ipv6addr"
    object_function = "next_available_ip"
    result_field    = "ips"
    object          = "ipv6network"
    object_parameters = {
      network      = "2001:db8:abcd:12::/64"
      network_view = "default"
    }
  }
  view = "default"
}
