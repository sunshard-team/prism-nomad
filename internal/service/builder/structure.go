package builder

import (
	"prism/internal/model"
)

type StructureBuilder interface {
	CustomBlock(name string, block []map[string]interface{}) model.TemplateBlock

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

	Expose() model.TemplateBlock
	ExposePath(block map[string]interface{}) model.TemplateBlock

	Gateway() model.TemplateBlock
	GatewayProxy(block []map[string]interface{}) model.TemplateBlock
	GatewayProxyAddress(block []map[string]interface{}) model.TemplateBlock
	GatewayIngress() model.TemplateBlock
	GatewayIngressTLS(block []map[string]interface{}) model.TemplateBlock
	GatewayIngressListener(block []map[string]interface{}) model.TemplateBlock
	GatewayIngressListenerService(block []map[string]interface{}) model.TemplateBlock
	GatewayTerminating() model.TemplateBlock
	GatewayTerminatingService(block []map[string]interface{}) model.TemplateBlock
	GatewayMesh() model.TemplateBlock

	Group(block []map[string]interface{}) model.TemplateBlock
	GroupConsul(block map[string]interface{}) model.TemplateBlock

	Identity(block []map[string]interface{}) model.TemplateBlock
	Job(block model.ConfigBlock) model.TemplateBlock
	Lifecycle(block []map[string]interface{}) model.TemplateBlock
	Logs(block []map[string]interface{}) model.TemplateBlock
	Meta(block []map[string]interface{}) model.TemplateBlock
	Migrate(block []map[string]interface{}) model.TemplateBlock

	Multiregion() model.TemplateBlock
	MultiregionStrategy(block map[string]interface{}) model.TemplateBlock
	MultiregionRegion(block map[string]interface{}) model.TemplateBlock

	Network(group []map[string]interface{}) model.TemplateBlock
	NetworkPort(block []map[string]interface{}) model.TemplateBlock
	NetworkDNS(block []map[string]interface{}) model.TemplateBlock

	Parameterized(block []map[string]interface{}) model.TemplateBlock
	Periodic(block []map[string]interface{}) model.TemplateBlock
	Proxy(block []map[string]interface{}) model.TemplateBlock
	Reschedule(block []map[string]interface{}) model.TemplateBlock
	Resources(block []map[string]interface{}) model.TemplateBlock
	Restart(block []map[string]interface{}) model.TemplateBlock
	Scaling(block []map[string]interface{}) model.TemplateBlock
	Service(block []map[string]interface{}) model.TemplateBlock
	SidecarService(block []map[string]interface{}) model.TemplateBlock
	SidecarTask(block []map[string]interface{}) model.TemplateBlock

	Spread(block []map[string]interface{}) model.TemplateBlock
	SpreadTarget(block map[string]interface{}) model.TemplateBlock

	Task(block []map[string]interface{}) model.TemplateBlock
	Template(block []map[string]interface{}) model.TemplateBlock
	Update(block []map[string]interface{}) model.TemplateBlock

	Upstream(block []map[string]interface{}) model.TemplateBlock
	MeshGateway(block []map[string]interface{}) model.TemplateBlock

	Vault(block []map[string]interface{}) model.TemplateBlock

	Volume(block []map[string]interface{}) model.TemplateBlock
	MountOptions(block map[string]interface{}) model.TemplateBlock

	VolumeMount(block []map[string]interface{}) model.TemplateBlock
}

type Structure struct{}

// Returns a block with any key-value parameters.
func (s *Structure) CustomBlock(
	name string,
	block []map[string]interface{},
) model.TemplateBlock {
	// parameters := make([]map[string]interface{}, 0)

	// for _, item := range block {
	// 	for k := range item {
	// 		if k == name {
	// 			parameters = append(parameters, item)
	// 			break
	// 		}
	// 	}
	// }

	templateBlock := model.TemplateBlock{
		BlockName: name,
		Parameter: block,
	}

	return templateBlock
}

func (s *Structure) Artifact(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "destination", "mode", "source":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "artifact",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Affinity(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "attribute", "operator", "value", "weight":
				parameters = append(parameters, item)
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
		for k := range item {
			switch k {
			case "command", "args", "timeout", "fail_on_error":
				parameters = append(parameters, item)
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
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "check",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) CheckRestart(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "limit", "grace", "ignore_warnings":
				parameters = append(parameters, item)
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
		for k := range item {
			if k == "native" {
				parameters = append(parameters, item)
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
		for k := range item {
			switch k {
			case "attribute", "operator", "value":
				parameters = append(parameters, item)
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
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
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
			switch k {
			case "name":
				name = v.(string)
			case "count":
				parameters = append(parameters, item)
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
		for k := range item {
			if k == "file" {
				parameters = append(parameters, item)
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
	templateBlock := model.TemplateBlock{
		BlockName: "env",
		Parameter: block,
	}

	return templateBlock
}

func (s *Structure) EphemeralDisk(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "migrate", "size", "sticky":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "ephemeral_disk",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Expose() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "expose",
	}

	return templateBlock
}

func (s *Structure) ExposePath(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k := range block {
		switch k {
		case "path", "protocol", "local_path_port", "listener_port":
			parameters = append(parameters, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "path",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Gateway() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "gateway",
	}

	return templateBlock
}

func (s *Structure) GatewayProxy(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"connect_timeout",
		"envoy_gateway_bind_tagged_addresses",
		"envoy_gateway_no_default_bind",
		"envoy_dns_discovery_type",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "proxy",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayProxyAddress(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "address", "port":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "envoy_gateway_bind_addresses",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayIngress() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "ingress",
	}

	return templateBlock
}

func (s *Structure) GatewayIngressTLS(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"enabled",
		"tls_min_version",
		"tls_max_version",
		"cipher_suites",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "tls",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayIngressListener(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "port", "protocol":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "listener",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayIngressListenerService(
	block []map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "name", "hosts":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayTerminating() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "terminating",
	}

	return templateBlock
}

func (s *Structure) GatewayTerminatingService(
	block []map[string]interface{},
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "name", "ca_file", "cert_file", "key_file", "sni":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GatewayMesh() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "mesh",
	}

	return templateBlock
}

func (s *Structure) Group(block []map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"count",
		"shutdown_delay",
		"stop_after_client_disconnect",
		"max_client_disconnect",
	}

	for _, item := range block {
		for k, v := range item {
			if k == "name" {
				name = v.(string)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "group",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) GroupConsul(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k := range block {
		if k == "namespace" {
			parameters = append(parameters, block)
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
		for k := range item {
			switch k {
			case "env", "file":
				parameters = append(parameters, item)
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
			if k == "name" {
				name = v.(string)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
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

func (s *Structure) Lifecycle(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "hook", "sidecar":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "lifecycle",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Logs(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "max_files", "max_files_size", "disabled":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "logs",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Meta(block []map[string]interface{}) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "meta",
		Parameter: block,
	}

	return templateBlock
}

func (s *Structure) Migrate(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"max_parallel",
		"health_check",
		"min_healthy_time",
		"healthy_deadline",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "migrate",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Multiregion() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "multiregion",
	}

	return templateBlock
}

func (s *Structure) MultiregionStrategy(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k := range block {
		switch k {
		case "max_parallel", "on_failure":
			parameters = append(parameters, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "strategy",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) MultiregionRegion(block map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for k, v := range block {
		switch k {
		case "name":
			name = v.(string)
		case "count", "datacenters", "node_pool":
			parameters = append(parameters, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "region",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Network(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "mode":
				parameters = append(parameters, item)
			case "hostname":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "network",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) NetworkPort(block []map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "static", "to", "host_network":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "port",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) NetworkDNS(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "server", "searches", "options":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "dns",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Parameterized(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "meta_optional", "meta_required", "payload":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "parameterized",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Periodic(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "cron", "prohibit_overlap", "time_zone", "enabled":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "periodic",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Proxy(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "local_service_address", "local_service_port":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "proxy",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Reschedule(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"attempts",
		"interval",
		"delay",
		"delay_function",
		"max_delay",
		"unlimited",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "reschedule",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Resources(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "cpu", "cores", "memory", "memory_max":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "resources",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Restart(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "attempts", "delay", "interval", "mode":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "restart",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Scaling(block []map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "min", "max", "enabled":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "scaling",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Service(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"provider",
		"name",
		"port",
		"tags",
		"canary_tags",
		"enable_tag_override",
		"address",
		"address_mode",
		"task",
		"on_update",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) SidecarService(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "disable_default_tcp_check", "port", "tags":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "sidecar_service",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) SidecarTask(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"name",
		"driver",
		"user",
		"kill_timeout",
		"shutdown_delay",
		"kill_signal",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "sidecar_task",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Spread(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "attribute", "weight":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "spread",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) SpreadTarget(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k := range block {
		switch k {
		case "value", "percent":
			parameters = append(parameters, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "target",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Task(block []map[string]interface{}) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"driver",
		"kill_timeout",
		"kill_signal",
		"leader",
		"shutdown_delay",
		"user",
		"kind",
	}

	for _, item := range block {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "task",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Template(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"change_mode",
		"change_signal",
		"destination",
		"env",
		"error_on_missing_key",
		"left_delimiter",
		"perms",
		"uid",
		"gid",
		"right_delimiter",
		"source",
		"splay",
		"vault_grace",
	}

	for _, item := range block {
		for k := range item {
			switch k {
			case "data":
				// read file and add data into "data" parameter.
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "template",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Update(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"max_parallel",
		"health_check",
		"min_healthy_time",
		"healthy_deadline",
		"progress_deadline",
		"auto_revert",
		"auto_promote",
		"canary",
		"stagger",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "update",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Upstream(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"destination_name",
		"destination_namespace",
		"datacenters",
		"local_bind_address",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "upstream",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) MeshGateway(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block {
		for k := range item {
			switch k {
			case "mode":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "mesh_gateway",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Vault(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"change_mode",
		"change_signal",
		"env",
		"disable_file",
		"namespace",
		"policies",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "vault",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Volume(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"type",
		"source",
		"read_only",
		"pear_alloc",
		"access_mode",
		"attachment_mode",
		"mount_options",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "volume",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) MountOptions(block map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for k := range block {
		switch k {
		case "fs_type", "mount_flags":
			parameters = append(parameters, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "mount_options",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) VolumeMount(block []map[string]interface{}) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"volume",
		"destination",
		"read_only",
		"propagation_mode",
	}

	for _, item := range block {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "volume_mount",
		Parameter: parameters,
	}

	return templateBlock
}
