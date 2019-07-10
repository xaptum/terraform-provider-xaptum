package enf

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceEnfFirewallRule() *schema.Resource {
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
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
			"dest_ip": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dest_port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  0,
			},
		},
	}
}

func resourceEnfFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)

	params := &FirewallRuleRequest{
		Priority:   Int(d.Get("priority").(int)),
		Action:     String(d.Get("action").(string)),
		Direction:  String(d.Get("direction").(string)),
		IPFamily:   String(d.Get("ip_family").(string)),
		Protocol:   String(d.Get("protocol").(string)),
		SourceIP:   String(d.Get("source_ip").(string)),
		SourcePort: Int(d.Get("source_port").(int)),
		DestIP:     String(d.Get("dest_ip").(string)),
		DestPort:   Int(d.Get("dest_port").(int)),
	}

	client := meta.(*EnfClient).Client
	rule, _, err := client.Firewall.CreateRule(context.Background(), network, params)
	if err != nil {
		return err
	}

	d.SetId(*rule.ID)

	return nil
}

func resourceEnfFirewallRuleRead(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)
	id := d.Id()

	client := meta.(*EnfClient).Client
	rule, _, err := client.Firewall.GetRule(context.Background(), network, id)
	if err != nil {
		// TODO: If just not found, run `d.SetId(""); return nil`
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
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Rule %s: %s", d.Id(), err)
	}

	return nil
}
