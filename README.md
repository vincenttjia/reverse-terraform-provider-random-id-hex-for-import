# reverse-terraform-provider-random-id-hex-for-import

If you are using [https://registry.terraform.io/modules/traveloka/resource-naming/aws/latest](https://registry.terraform.io/modules/traveloka/resource-naming/aws/latest) and needed to re-create/import the resources using the same random id

You can run the script above or use this link [https://go.dev/play/p/zoksTAHECp5](https://go.dev/play/p/zoksTAHECp5)

Then do 

```
terraform import module.this.random_id.this TheGolangOutput
```

Example for input abcdata-app-971c842539f7ec43
```
terraform import module.this.random_id.this lxyEJTn37EM
```
