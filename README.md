terraform-provider-amp
======================

Custom provider to allow auto accept VPC Peering connection 
requests from different AWS Account ID. At the moment,
aws_vpc_peering_connection resource only allow auto accept
requests if they belong to a same AWS Account ID.

Moreover, as of this stage, only create resource has been
implemented.

Status: [![Build Status](https://travis-ci.org/lenfree/TERRAFORM-provider-amp.svg?branch=master)](https://travis-ci.org/lenfree/TERRAFORM-provider-amp)


How to use this resource:

```
resource "amp_vpc_peering_accept_all" "mytest" {
  owner_id        = "<ACCOUNT-01"
  aws_account_ids = ["ACCOUNT-01", "ACCOUNT-02", "ACCOUNT-03"]
  aws_region      = "ap-southeast-2"
```

Binary release available:

```
https://github.com/lenfree/TERRAFORM-provider-amp/releases
```

Requirement:

# Make sure binary exists on project directory where you call
`terraform apply`

## Contributing

1. Fork it ( https://github.com/lenfree/vcenter-metrics?fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
