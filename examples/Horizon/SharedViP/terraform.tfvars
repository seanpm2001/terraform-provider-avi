# Values of the variables required for configuring Avi for Horizon deployment in a Shared VIP with L7 and L4 VS.

ipaddr_placement = "10.10.24.0"
mgmt_net = "Mgmt_Arista"
cloud_name = "VMwareCloud-Horizon"
ip_vip = "10.10.24.204"
domain_name = "demo.horizondemo.com"
avi_controller = "10.10.24.93"
avi_username = "admin"
avi_password = "<your password here>"
avi_api_version = "18.2.7"
pool_server1 = "10.10.24.90"
pool_server2 = "10.10.24.108"
app_profile = "System-Secure-HTTP-VDI"
horizon_cert = "HorizonDemoCert"
horizon_hm = "https-monitor-horizonUAG"
udp_profile = "System-UDP-Fast-Path-VDI"
tcp_profile = "System-TCP-Proxy"
ip_group = "UAG_Servers"
ssl_profile = "System-Standard"
l4_pool = "L4-pool-horizon"
l4_app_profile = "System-L4-Application"
l7_pool = "L7-pool-horizon"
l7_vs = "L7-VS_Horizon"
l4_vs = "L4_VS_Horizon"
ip_vip_l4 = "10.10.24.250"
l4_domain_name = "demo.l4vshorizon.com"