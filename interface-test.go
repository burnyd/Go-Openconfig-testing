package main

import (
	"fmt"

	oc "github.com/burnyd/Go-Openconfig-testing/pkg/ocdemo"
	"github.com/openconfig/ygot/ygot"
)

func main() {
	d := &oc.Device{}

	i, err := d.NewInterface("eth0")
	i.AdminStatus = oc.OpenconfigInterfaces_Interface_AdminStatus_UP
	i.Mtu = ygot.Uint16(1500)
	i.Description = ygot.String("An Interface")

	if err := d.Interface["eth0"].Validate(); err != nil {
		panic(fmt.Sprintf("Interface validation failed: %v", err))
	}

	d.Interface["eth1"] = &oc.Interface{
		Name:        ygot.String("eth1"),
		Description: ygot.String("Another Interface"),
		Enabled:     ygot.Bool(false),
		Type:        oc.IETFInterfaces_InterfaceType_ethernetCsmacd,
	}


	json, err := ygot.EmitJSON(d, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("JSON demo error: %v", err))
	}

	//fmt.Println(i)
	//fmt.Println(d)
	//fmt.Println(j)
	//fmt.Println(s)
	fmt.Println(json)
	//fmt.Println(a)
}
