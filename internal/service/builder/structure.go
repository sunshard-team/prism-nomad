package builder

import (
	"fmt"
	"prism/internal/model"
	"strconv"
)

type StructureBuilder interface {
	Job(block model.ConfigBlock) model.TemplateBlock
	Group(block interface{}) model.TemplateBlock
}

type Structure struct{}

func (s *Structure) Job(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]string, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "all_at_once":
				i := fmt.Sprintf(`all_at_once: %v`, v.(bool))
				parameters = append(parameters, i)
			case "datacenters":
				var datacenter []string

				for i, item := range v.([]interface{}) {
					var value string

					if i+1 == len(v.([]interface{})) {
						value = fmt.Sprintf(`"%v"`, item.(string))
					} else {
						value = fmt.Sprintf(`"%v",`, item.(string))
					}

					datacenter = append(datacenter, value)
				}

				i := fmt.Sprintf(`datacenters: %v`, datacenter)
				parameters = append(parameters, i)
			case "node_pool":
				i := fmt.Sprintf(`node_pool: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "namespace":
				i := fmt.Sprintf(`namespace: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "parameterized":
				i := fmt.Sprintf(`parameterized: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "priority":
				i := fmt.Sprintf(`priopity: %v`, strconv.Itoa(v.(int)))
				parameters = append(parameters, i)
			case "region":
				i := fmt.Sprintf(`region: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "type":
				i := fmt.Sprintf(`type: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "vault_token":
				i := fmt.Sprintf(`vault_token: "%s"`, v.(string))
				parameters = append(parameters, i)
			case "consul_token":
				i := fmt.Sprintf(`consul_token: "%s"`, v.(string))
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "job",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Group(block interface{}) model.TemplateBlock {
	var name string
	parameters := make([]string, 0)

	for k, v := range block.(map[string]interface{}) {
		switch k {
		case "name":
			name = v.(string)
		case "count":
			i := fmt.Sprintf(`count: %v`, v.(int))
			parameters = append(parameters, i)
		case "shutdown_delay":
			i := fmt.Sprintf(`shutdown_delay: "%s"`, v.(string))
			parameters = append(parameters, i)
		case "stop_after_client_disconnect":
			i := fmt.Sprintf(`stop_after_client_disconnect: "%s"`, v.(string))
			parameters = append(parameters, i)
		case "max_client_disconnect":
			i := fmt.Sprintf(`max_client_disconnect: "%s"`, v.(string))
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "group",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}
