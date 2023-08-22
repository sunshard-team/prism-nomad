package builder

import (
	"prism/internal/model"
)

type StructureBuilder interface {
	Artifact(block []map[string]interface{}) model.TemplateBlock
	Affinity(block []map[string]interface{}) model.TemplateBlock
	ChangeScript(block []map[string]interface{}) model.TemplateBlock
	Check(block []map[string]interface{}) model.TemplateBlock
	CheckRestart(block []map[string]interface{}) model.TemplateBlock
	Connect(block []map[string]interface{}) model.TemplateBlock
	Constraint(block []map[string]interface{}) model.TemplateBlock
	CSIPlugin(block []map[string]interface{}) model.TemplateBlock
	Device(block []map[string]interface{}) model.TemplateBlock
	DispatchPayload(block []map[string]interface{}) model.TemplateBlock
	Env(block []map[string]interface{}) model.TemplateBlock
	EphemeralDisk(block []map[string]interface{}) model.TemplateBlock
	Expose(block []map[string]interface{}) model.TemplateBlock
	Gateway(block []map[string]interface{}) model.TemplateBlock
	Group(block map[string]interface{}) model.TemplateBlock
	Identity(block []map[string]interface{}) model.TemplateBlock
	Job(block model.ConfigBlock) model.TemplateBlock
}

type Structure struct{}

// Returns a simple block with key-value parameters.
func (s *Structure) customBlock(
	name string,
	block map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == name {
			parameters = append(parameters, i)
			break
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Artifact(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	var internalBlock []model.TemplateBlock

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "destination", "mode", "source":
				parameters = append(parameters, i)
			case "options", "headers":
				block := s.customBlock(k, i)
				internalBlock = append(internalBlock, block)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "artifact",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Affinity(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "attribute", "operator", "value", "weight":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "affinity",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) ChangeScript(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "command", "args", "timeout", "fail_on_error":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "change_script",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Check(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	var internalBlock []model.TemplateBlock

	parameterName := []string{
		"address_mode",
		"args",
		"command",
		"grcp_service",
		"grpc_use_tls",
		"initial_status",
		"success_before_passing",
		"failures_before_critical",
		"interval",
		"method",
		"body",
		"name",
		"path",
		"expose",
		"port",
		"protocol",
		"task",
		"timeout",
		"type",
		"tls_server_name",
		"tls_skip_verify",
		"on_update",
	}

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "header" {
				header := s.customBlock(k, i)
				internalBlock = append(internalBlock, header)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, i)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "check",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) CheckRestart(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "limit", "grace", "ignore_warnings":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "check_restart",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Connect(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "native" {
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "connect",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Constraint(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "attribute", "operator", "value":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "constraint",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) CSIPlugin(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	parameterName := []string{
		"id",
		"type",
		"mount_dir",
		"stage_publish_base_dir",
		"health_timeout",
	}

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, i)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "csi_plugin",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Device(block []map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "name":
				name = v.(string)
			case "count":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		Name:      name,
		BlockName: "device",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) DispatchPayload(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "file" {
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "dispatch_payload",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Env(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "env",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) EphemeralDisk(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "migrate", "size", "sticky":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "ephemeral_disk",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Expose(block []map[string]interface{}) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "path" {
				block := s.exposePath(i)
				internalBlock = append(internalBlock, block)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "expose",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) exposePath(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "path", "protocol", "local_path_port":
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "path",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Gateway(block []map[string]interface{}) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "proxy":
				block := s.gatewayProxy(i)
				internalBlock = append(internalBlock, block)
			case "ingress":
				block := s.gatewayIngress(i)
				internalBlock = append(internalBlock, block)
			case "terminating":
				block := s.gatewayTerminating(i)
				internalBlock = append(internalBlock, block)
			case "mesh":
				mesh := model.TemplateBlock{
					BlockName: "mesh",
				}
				internalBlock = append(internalBlock, mesh)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "gateway",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) gatewayProxy(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	var internalBlock []model.TemplateBlock

	parameterName := []string{
		"connect_timeout",
		"envoy_gateway_bind_tagged_addresses",
		"envoy_gateway_no_default_bind",
		"envoy_dns_discovery_type",
	}

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == "envoy_gateway_bind_addresses" {
			address := s.gatewayProxyAddress(i)
			internalBlock = append(internalBlock, address)
		}

		if k == "config" {
			config := s.customBlock(k, i)
			internalBlock = append(internalBlock, config)
		}

		for _, p := range parameterName {
			switch k {
			case p:
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "proxy",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) gatewayProxyAddress(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "address", "port":
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "envoy_gateway_bind_addresses",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) gatewayIngress(block map[string]interface{}) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "tls":
			tls := s.gatewayIngressTLS(i)
			internalBlock = append(internalBlock, tls)
		case "listener":
			listener := s.gatewayIngressListener(i)
			internalBlock = append(internalBlock, listener)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "ingress",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) gatewayIngressTLS(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"enabled",
		"tls_min_version",
		"tls_max_version",
		"cipher_suites",
	}

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		for _, p := range parameterName {
			switch k {
			case p:
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "tls",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) gatewayIngressListener(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	var internalBlock []model.TemplateBlock

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "port", "protocol":
			parameters = append(parameters, i)
		case "service":
			service := s.gatewayIngressListenerService(i)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "listener",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) gatewayIngressListenerService(
	block map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "name", "hosts":
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) gatewayTerminating(block map[string]interface{}) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == "service" {
			service := s.gatewayTerminatingService(i)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "terminating",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) gatewayTerminatingService(
	block map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		switch k {
		case "name", "ca_file", "cert_file", "key_file", "sni":
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Group(block map[string]interface{}) model.TemplateBlock {
	var name string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"count",
		"shutdown_delay",
		"stop_after_client_disconnect",
		"max_client_disconnect",
	}

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == "name" {
			name = v.(string)
		}

		if k == "consul" {
			consul := s.groupConsul(i)
			internalBlock = append(internalBlock, consul)
		}

		for _, p := range parameterName {
			switch k {
			case p:
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "group",
		Name:      name,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) groupConsul(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		i := make(map[string]interface{})
		i[k] = v

		if k == "namespace" {
			parameters = append(parameters, i)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "consul",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Identity(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			switch k {
			case "env", "file":
				parameters = append(parameters, i)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "identity",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Job(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"all_at_once",
		"datacenters",
		"node_pool",
		"namespace",
		"priority",
		"region",
		"type",
		"vault_token",
		"consul_token",
	}

	for _, item := range block.Parameter {
		for k, v := range item {
			i := make(map[string]interface{})
			i[k] = v

			if k == "name" {
				name = v.(string)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, i)
				}
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
