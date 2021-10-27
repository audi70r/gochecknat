# gochecknat
A simple Goland tool, that 
- gets the public IP address
- checks if the network is behind an asymmetric NAT (useful information for UDP hole punching).

```
go get github.com/audi70r/gochecknat
```


```gochecknat.GetNATInfo()``` will return information about public ip, a port, that is assigned by the NAT and whether the called is behind a symmetric NAT.