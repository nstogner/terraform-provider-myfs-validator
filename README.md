# Example Terraform Provider (with Validations)

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
