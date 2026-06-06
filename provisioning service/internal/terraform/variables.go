package terraform

import (
	"os"
)

type Variables struct {
	Namespace   string
	Postgres    string
	Redis       string
}

func WriteTFVars(
	vars Variables,
) error {

	content := []byte(
		`namespace="` + vars.Namespace + `"
postgres_name="` + vars.Postgres + `"
redis_name="` + vars.Redis + `"
`,
	)

	return os.WriteFile(
		"./terraform/terraform.tfvars",
		content,
		0644,
	)
}
