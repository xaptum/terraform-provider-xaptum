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
			},
		},
	}
}

func createInt(x int) *int {
	return &x
}

func createString(x string) *string {
	return &x
}

func resourceEnfFirewallRuleCreate(d *schema.ResourceData, meta interface{}) error {
	network := d.Get("network").(string)

	params := &FirewallRuleRequest{
		Priority:   createInt(d.Get("priority").(int)),
		Action:     createString(d.Get("action").(string)),
		Direction:  createString(d.Get("direction").(string)),
		IPFamily:   createString(d.Get("ip_family").(string)),
		Protocol:   createString(d.Get("protocol").(string)),
		SourceIP:   createString(d.Get("source_ip").(string)),
		SourcePort: createInt(d.Get("source_port").(int)),
		DestIP:     createString(d.Get("dest_ip").(string)),
		DestPort:   createInt(d.Get("dest_port").(int)),
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
	d.Set("source_ip", rule.DestIP)
	d.Set("source_port", rule.DestPort)

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
