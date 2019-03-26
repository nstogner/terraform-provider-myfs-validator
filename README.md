# Example Terraform Provider (with Validations)

This provider acts as a shim for the [terraform-provider-myfs](https://github.com/nstogner/terraform-provider-myfs) provider and adds a validation step before calling create/update on a resource.

```sh
go build -o terraform-provider-myfs .
terraform init

# Apply should fail due to a validation error.
terraform apply

# Update text = "Hello!"
vi ./main.tf

# Apply should now succeed.
terraform apply
```
