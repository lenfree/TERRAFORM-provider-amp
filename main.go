package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return Provider()
        },
    })
}

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
                "amp_vpc_peering_accept_all": resourceServer(),
        },
    }
}

func resourceServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceServerCreate,
        Read:   resourceServerRead,
        Update: resourceServerUpdate,
        Delete: resourceServerDelete,

        Schema: map[string]*schema.Schema{
            "owner_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "aws_account_ids": &schema.Schema{
                Type:     schema.TypeList,
                Required: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "aws_region": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        owner_id := d.Get("owner_id").(string)
        aws_account_ids := d.Get("aws_account_ids").([]interface{})
        aws_region := d.Get("aws_region").(string)
        err := query(owner_id, aws_account_ids, aws_region)
        if err != nil {
                return err
        }
        d.SetId(owner_id)
        return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}

func query(owner_id string, aws_account_ids []interface{}, aws_region string, ) error {
         svc := ec2.New(session.New(), &aws.Config{ Region: aws.String(aws_region) })

         params := &ec2.DescribeVpcPeeringConnectionsInput{}
         resp, err := svc.DescribeVpcPeeringConnections(params)
         if err != nil {
                 fmt.Println(err.Error())
                 return err
         }
         parseResponse(resp, svc, owner_id, aws_account_ids)
         return nil
}

func parseResponse(resp *ec2.DescribeVpcPeeringConnectionsOutput, svc *ec2.EC2, owner_id string, aws_account_ids []interface{}, ) {
        for _, v := range resp.VpcPeeringConnections {
                if isValidAccount(v.RequesterVpcInfo.OwnerId, aws_account_ids) == true && isOwner(v.AccepterVpcInfo.OwnerId, owner_id) == true {
                        err := acceptPeeringRequest(v, svc)
                        if err != nil {
                                fmt.Println(err.Error())
                        }
                }
        }
}

func isValidAccount(vpc *string, account_ids []interface{}) bool {

        for _, v := range account_ids {
                if v == *vpc {
                        return true
                }
        }
        return false
}

func isOwner(vpc *string, owner_id string) bool {
         if owner_id == *vpc {
                 return true
         }
         return false
}

func acceptPeeringRequest(v *ec2.VpcPeeringConnection, svc *ec2.EC2) error {
        if *v.Status.Code == "pending-acceptance" {
                params := &ec2.AcceptVpcPeeringConnectionInput{
                        VpcPeeringConnectionId: aws.String(*v.VpcPeeringConnectionId),
                }
                resp, err := svc.AcceptVpcPeeringConnection(params)
                if err != nil {
                        return err
                }
                fmt.Printf("VPC peering accepted")
                fmt.Printf("%s\n", resp)
        }
        return nil
}
