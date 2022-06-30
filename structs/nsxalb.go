package structs

import (
	"fmt"
	"strings"
)

type Cloud struct {
	Id        string `json:"uuid"`
	Name      string `json:"name"`
	TenantRef string `json:"tenant_ref"`
}

type ServiceEngineGroup struct {
	Id           string `json:"uuid"`
	Name         string `json:"name"`
	HaMode       string `json:"ha_mode"`
	Buffer       int    `json:"buffer_se"`
	VsLimitPerSe int    `json:"max_vs_per_se"`
	TenantRef    string `json:"tenant_ref"`
	SE           []ServiceEngine
}

type ServiceEngine struct {
	Primary    bool              `json:"is_primary"`
	Secondary  bool              `json:"is_secondary"`
	Memory     int               `json:"memory"`
	Address    map[string]string `json:"mgmt_ip"`
	SeRef      string            `json:"se_ref"`
	Cpu        int               `json:"vcpus"`
	Interfaces []SeInterface     `json:"vip_intf_list"`
	VipMac     string            `json:"vip_intf_mac"`
	VipMask    int               `json:"vip_intf_mask"`
}

type SeInterface struct {
	Address map[string]string `json:"vip_intf_ip"`
	Mac     string            `json:"vip_intf_mac"`
	Vlan    int               `json:"vlan_id"`
}

type VSResult struct {
	VirtualServices []VirtualService `json:"results"`
}

type VirtualService struct {
	Type        string   `json:"type"`
	UUID        string   `json:"uuid"`
	Tenant      string   `json:"tenant_ref"`
	Name        string   `json:"name"`
	Ports       []VSPort `json:"services"`
	Cloud       string   `json:"cloud_type"`
	CloudRef    string   `json:"cloud_ref"`
	AutoGateway bool     `json:"enable_autogw"`
	SeGroupRef  string   `json:"se_group_ref"`
	Vips        []Vip    `json:"vip"`
}

func (v *VirtualService) Print(cloud string, seg string) {
	ports := ""
	for i, p := range v.Ports {
		summary := p.GetSummary()
		ports += summary
		if i != len(v.Ports)-1 {
			ports += ","
		}
	}
	vips := ""
	for i, vip := range v.Vips {
		vips += vip.Address["addr"]
		if i != len(v.Vips)-1 {
			ports += ","
		}
	}
	fmt.Printf("%-51v  %-16v  %-15v  %-5v  %-12v  %-16v\n", v.UUID, v.Name, vips, ports, cloud, seg)
}

func (v *VirtualService) GetCloudId() string {
	path := strings.Split(v.CloudRef, "/")

	return path[len(path)-1]
}

func (v *VirtualService) GetSegId() string {
	path := strings.Split(v.SeGroupRef, "/")

	return path[len(path)-1]
}

type Vip struct {
	Id      string            `json:"vip_id"`
	Address map[string]string `json:"ip_address"`
}

type VipRuntime struct {
	Se []ServiceEngine `json:"se_list"`
}

type VSPort struct {
	Ssl       bool `json:"enable_ssl"`
	Port      uint `json:"port"`
	PortRange uint `json:"port_range_end"`
}

func (p *VSPort) GetSummary() string {
	var summary string
	if p.Port != p.PortRange {
		summary = fmt.Sprintf("%v-%v", p.Port, p.PortRange)
	} else {
		summary = fmt.Sprintf("%v", p.Port)
	}
	if p.Ssl {
		summary += fmt.Sprintf("(SSL)")
	}

	return summary
}
