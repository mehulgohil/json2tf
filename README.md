# json2tf

A command line tool which helps to transform resource json to terraform data.

# Installation & Usage
```bash
go install github.com/mehulgohil/json2tf@latest
```

# Usage
```bash
json2tf azure -j '{\"id\":\"/subscriptions/aa9d0c7d-9baa-49ad-954e-34089fa7d2a9/resourceGroups/rgrp-test001\",\"name\":\"rgrp-test001\",\"type\":\"Microsoft.Resources/resourceGroups\",\"location\":\"westindia\",\"tags\":{\"test\":\"test\"},\"properties\":{\"provisioningState\":\"Succeeded\"}}'
```

