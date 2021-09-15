# GoFilex is a TCP file transfer by golang

## Why and What it does

In recent pentest engagement, the target host was protected by Zscalar proxy and Carbon Black / McAfee EDP/EDR.
The Zscalar only allow the host to visit trusted websites, no direct IP via HTTP/HTTPS, domain name category check to prevent file transfer via HTTP/HTTPS to the new created shellbox.
Most of network traffic been blocked to go out, such as SSH, FTP, etc.
During enumeration, identified that the host were allowed to access port 80, 443 on Internet-based shellbox. However, the classic well-known NetCat swisse knife is blocked by EDP :/

