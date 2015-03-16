## Minimal squeleton of a DNS server

This is nothing, just a **minimal** DNS sniffer, I had an idea but realized it wasn't worth it. However this migh be reused as a starting point.

Interesting DNS server can be found in Consul and Skydns (both used the same package).

To work with it, add an entry in `resolver` (mac os x) or in `resolv.conf` (linux) for a particular domain (i.e: dev) and hit something.dev.

That's it.
