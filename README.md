# Go-Openconfig-testing
git clone https://github.com/burnyd/Go-Openconfig-testing
go run interface-test.go
#Pray it works. 
Should come out with some json output of openconfig interfaces. 

{
  "openconfig-interfaces:interfaces": {
    "interface": [
      {
        "config": {
          "description": "An Interface",
          "mtu": 1500,
          "name": "eth0"
        },
        "name": "eth0",
        "state": {
          "admin-status": "UP"
        }
      },
      {
        "config": {
          "description": "Another Interface",
          "enabled": false,
          "name": "eth1",
          "type": "iana-if-type:ethernetCsmacd"
        },
        "name": "eth1"
      }
    ]
  }
}


