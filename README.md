# Zero Touch Provisioning (ZTP) for Whitebox fabric 
- It uses DHCP information to see if a switch bootsup and matches it's config based on the MAC in the DHCP request (using DHCP harvest)
- Then it calls the Go program wich configures the switch
- Configuration includes, copying the appropriate OFDPA image, installing ofdpa, restarting ofdpa, hostname, make sure the configuration is persisted across reboots, etc.