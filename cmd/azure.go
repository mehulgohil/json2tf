package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"

	"github.com/rodaine/hclencoder"
	"github.com/spf13/cobra"
)

type AzureRMResourceGroup struct {
	Name string `hcl:"name", json:"name"`
}

var dataAzuremResourceGroupTemplate = `data "azurerm_resource_group" "example" {
  {{ .HclCode }}
}
`

// azureCmd represents the azure command
var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "azure command helps to transform json to azurem provider data",
	Long:  `azure command helps to transform json to azurem provider data`,
	Run: func(cmd *cobra.Command, args []string) {
		getJson, err := cmd.Flags().GetString("json")
		if err != nil {
			log.Fatal(err)
		}

		var temp AzureRMResourceGroup
		err = json.Unmarshal([]byte(getJson), &temp)
		if err != nil {
			log.Fatal(err)
		}

		hcl, err := hclencoder.Encode(temp)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err := template.New("data_resource_group").Parse(dataAzuremResourceGroupTemplate)
		if err != nil {
			log.Fatal(err)
		}
		var buff = bytes.NewBufferString("")

		err = tmpl.Execute(buff, map[string]string{
			"HclCode": string(bytes.TrimSpace(hcl)),
		})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buff)
	},
}

func init() {
	rootCmd.AddCommand(azureCmd)
	azureCmd.Flags().StringP("json", "j", "", "provide azure resource json in the string format")
}
