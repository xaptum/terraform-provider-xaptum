package xaptum

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceXaptumFirewallRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnfFirewallRuleCreate,
		Read:   resourceEnfFirewallRuleRead,
		Delete: resourceEnfFirewallRuleDelete,

		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"direction": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_family": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_ip": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressDefaultCIDR,
			},
			"source_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
			"dest_ip": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressDefaultCIDR,
			},
			"dest_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
		},
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				id := d.Id()
				s := strings.Split(id, ".")
				d.Set("network", s[0])
				d.SetId(s[1])
				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func resourceEnfFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)

	params := &enf.FirewallRuleRequest{
		Priority:   enf.Int(d.Get("priority").(int)),
		Action:     enf.String(d.Get("action").(string)),
		Direction:  enf.String(d.Get("direction").(string)),
		IPFamily:   enf.String(d.Get("ip_family").(string)),
		Protocol:   enf.String(d.Get("protocol").(string)),
		SourceIP:   enf.String(d.Get("source_ip").(string)),
		SourcePort: enf.Int(d.Get("source_port").(int)),
		DestIP:     enf.String(d.Get("dest_ip").(string)),
		DestPort:   enf.Int(d.Get("dest_port").(int)),
	}

	client := meta.(*EnfClient).Client
	rule, _, err := client.Firewall.CreateRule(context.Background(), network, params)
	if err != nil {
		return err
	}

	d.SetId(*rule.ID)

	return resourceEnfFirewallRuleRead(d, meta)
}

func resourceEnfFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)
	id := d.Id()

	client := meta.(*EnfClient).Client
	rule, resp, err := client.Firewall.GetRule(context.Background(), network, id)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("network", rule.Network)
	d.Set("priority", rule.Priority)
	d.Set("action", rule.Action)
	d.Set("direction", rule.Direction)
	d.Set("ip_family", rule.IPFamily)
	d.Set("protocol", rule.Protocol)
	d.Set("source_ip", rule.SourceIP)
	d.Set("source_port", rule.SourcePort)
	d.Set("dest_ip", rule.DestIP)
	d.Set("dest_port", rule.DestPort)

	return nil
}

func resourceEnfFirewallRuleDelete(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)
	id := d.Id()

	client := meta.(*EnfClient).Client
	_, err := client.Firewall.DeleteRule(context.Background(), network, id)

	return err
}

func suppressDefaultCIDR(k, old, new string, d *schema.ResourceData) bool {
	network := d.Get("network").(string)
	direction := d.Get("direction").(string)

	if ((k == "source_ip" && direction == "EGRESS") ||
		(k == "dest_ip" && direction == "INGRESS")) &&
		(old == network && new == "") {
		return true
	}

	return false
}
