package builder

import (
	"fmt"
	"log"
	"os"
	"prism/internal/model"
)

type StructureBuilder interface {
	CustomBlock(name string, block []map[string]interface{}) model.TemplateBlock

	Artifact(block model.ConfigBlock) model.TemplateBlock
	Affinity(block model.ConfigBlock) model.TemplateBlock
	ChangeScript(block model.ConfigBlock) model.TemplateBlock
	Check(block model.ConfigBlock) model.TemplateBlock
	CheckRestart(block model.ConfigBlock) model.TemplateBlock
	Connect(block model.ConfigBlock) model.TemplateBlock
	Constraint(block model.ConfigBlock) model.TemplateBlock
	CSIPlugin(block model.ConfigBlock) model.TemplateBlock
	Device(block model.ConfigBlock) model.TemplateBlock
	DispatchPayload(block model.ConfigBlock) model.TemplateBlock
	Env(block model.ConfigBlock) model.TemplateBlock
	EphemeralDisk(block model.ConfigBlock) model.TemplateBlock

	Expose(block model.ConfigBlock) model.TemplateBlock
	ExposePath(block model.ConfigBlock) model.TemplateBlock

	Gateway(model.ConfigBlock) model.TemplateBlock
	GatewayProxy(block model.ConfigBlock) model.TemplateBlock
	GatewayProxyAddress(block model.ConfigBlock) model.TemplateBlock
	GatewayIngress(model.ConfigBlock) model.TemplateBlock
	GatewayIngressTLS(block model.ConfigBlock) model.TemplateBlock
	GatewayIngressListener(block model.ConfigBlock) model.TemplateBlock
	GatewayIngressListenerService(block model.ConfigBlock) model.TemplateBlock
	GatewayTerminating(model.ConfigBlock) model.TemplateBlock
	GatewayTerminatingService(block model.ConfigBlock) model.TemplateBlock
	GatewayMesh() model.TemplateBlock

	Group(block model.ConfigBlock) model.TemplateBlock
	GroupConsul(block model.ConfigBlock) model.TemplateBlock

	Identity(block model.ConfigBlock) model.TemplateBlock
	Job(block model.ConfigBlock) model.TemplateBlock
	Lifecycle(block model.ConfigBlock) model.TemplateBlock
	Logs(block model.ConfigBlock) model.TemplateBlock
	Meta(block model.ConfigBlock) model.TemplateBlock
	Migrate(block model.ConfigBlock) model.TemplateBlock

	Multiregion(model.ConfigBlock) model.TemplateBlock
	MultiregionStrategy(block model.ConfigBlock) model.TemplateBlock
	MultiregionRegion(block model.ConfigBlock) model.TemplateBlock

	Network(group model.ConfigBlock) model.TemplateBlock
	NetworkPort(block model.ConfigBlock) model.TemplateBlock
	NetworkDNS(block model.ConfigBlock) model.TemplateBlock

	Parameterized(block model.ConfigBlock) model.TemplateBlock
	Periodic(block model.ConfigBlock) model.TemplateBlock
	Proxy(block model.ConfigBlock) model.TemplateBlock
	Reschedule(block model.ConfigBlock) model.TemplateBlock
	Resources(block model.ConfigBlock) model.TemplateBlock
	Restart(block model.ConfigBlock) model.TemplateBlock
	Scaling(block model.ConfigBlock) model.TemplateBlock
	Service(block model.ConfigBlock) model.TemplateBlock
	SidecarService(block model.ConfigBlock) model.TemplateBlock
	SidecarTask(block model.ConfigBlock) model.TemplateBlock

	Spread(block model.ConfigBlock) model.TemplateBlock
	SpreadTarget(block model.ConfigBlock) model.TemplateBlock

	Task(block model.ConfigBlock) model.TemplateBlock
	Template(block model.ConfigBlock, projectPath string) model.TemplateBlock
	Update(block model.ConfigBlock) model.TemplateBlock

	Upstream(block model.ConfigBlock) model.TemplateBlock
	UpstreamMeshGateway(block model.ConfigBlock) model.TemplateBlock

	Vault(block model.ConfigBlock) model.TemplateBlock

	Volume(block model.ConfigBlock) model.TemplateBlock
	VolumeMountOptions(block model.ConfigBlock) model.TemplateBlock

	VolumeMount(block model.ConfigBlock) model.TemplateBlock
}

type Structure struct{}

// Returns a block with any key-value parameters.
func (s *Structure) CustomBlock(
	name string,
	block []map[string]interface{},
) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: name,
		Parameter: block,
	}

	return templateBlock
}

func (s *Structure) Artifact(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "destination", "mode", "source":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "options" || item.Name == "headers" {
			block := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "artifact",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Affinity(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) ChangeScript(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Check(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"address_mode",
		"args",
		"command",
		"grpc_service",
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

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "header" {
			header := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, header)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "check",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) CheckRestart(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Connect(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "native", "open_sidecar_service":
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

func (s *Structure) Constraint(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) CSIPlugin(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)
	parameterName := []string{
		"id",
		"type",
		"mount_dir",
		"stage_publish_base_dir",
		"health_timeout",
	}

	for _, item := range block.Parameter {
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

func (s *Structure) Device(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) DispatchPayload(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Env(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "env",
		Parameter: block.Parameter,
	}

	return templateBlock
}

func (s *Structure) EphemeralDisk(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Expose(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Name == "path" {
			path := s.ExposePath(item)
			internalBlock = append(internalBlock, path)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "expose",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) ExposePath(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "path", "protocol", "local_path_port", "listener_port":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "path",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Gateway(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "proxy":
			proxy := s.GatewayProxy(item)
			internalBlock = append(internalBlock, proxy)
		case "ingress":
			ingress := s.GatewayIngress(item)
			internalBlock = append(internalBlock, ingress)
		case "terminating":
			terminating := s.GatewayTerminating(item)
			internalBlock = append(internalBlock, terminating)
		case "mesh":
			mesh := s.GatewayMesh()
			internalBlock = append(internalBlock, mesh)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "gateway",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) GatewayProxy(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"connect_timeout",
		"envoy_gateway_bind_tagged_addresses",
		"envoy_gateway_no_default_bind",
		"envoy_dns_discovery_type",
	}

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		switch item.Name {
		case "envoy_gateway_bind_addresses":
			address := s.GatewayProxyAddress(item)
			internalBlock = append(internalBlock, address)
		case "config":
			config := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "proxy",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) GatewayProxyAddress(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) GatewayIngress(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "tls":
			tls := s.GatewayIngressTLS(item)
			internalBlock = append(internalBlock, tls)
		case "listener":
			listener := s.GatewayIngressListener(item)
			internalBlock = append(internalBlock, listener)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "ingress",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) GatewayIngressTLS(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"enabled",
		"tls_min_version",
		"tls_max_version",
		"cipher_suites",
	}

	for _, item := range block.Parameter {
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

func (s *Structure) GatewayIngressListener(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "port", "protocol":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "service" {
			service := s.GatewayIngressListenerService(item)
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

func (s *Structure) GatewayIngressListenerService(
	block model.ConfigBlock,
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) GatewayTerminating(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Name == "service" {
			service := s.GatewayTerminatingService(item)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "terminating",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) GatewayTerminatingService(
	block model.ConfigBlock,
) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Group(block model.ConfigBlock) model.TemplateBlock {
	var name string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"count",
		"shutdown_delay",
		"stop_after_client_disconnect",
		"max_client_disconnect",
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

	for _, item := range block.Block {
		if item.Name == "consul" {
			consul := s.GroupConsul(item)
			internalBlock = append(internalBlock, consul)
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

func (s *Structure) GroupConsul(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			if k == "namespace" {
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "consul",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Identity(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Lifecycle(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Logs(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "max_files", "max_file_size", "disabled":
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

func (s *Structure) Meta(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "meta",
		Parameter: block.Parameter,
	}

	return templateBlock
}

func (s *Structure) Migrate(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"max_parallel",
		"health_check",
		"min_healthy_time",
		"healthy_deadline",
	}

	for _, item := range block.Parameter {
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

func (s *Structure) Multiregion(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "strategy":
			strategy := s.MultiregionStrategy(item)
			internalBlock = append(internalBlock, strategy)
		case "region":
			region := s.MultiregionRegion(item)
			internalBlock = append(internalBlock, region)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "multiregion",
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) MultiregionStrategy(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "max_parallel", "on_failure":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "strategy",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) MultiregionRegion(block model.ConfigBlock) model.TemplateBlock {
	var name string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "count", "datacenters", "node_pool":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "meta" {
			meta := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, meta)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "region",
		Name:      name,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Network(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "mode":
				parameters = append(parameters, item)
			case "hostname":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		switch item.Name {
		case "port":
			port := s.NetworkPort(item)
			internalBlock = append(internalBlock, port)
		case "dns":
			dns := s.NetworkDNS(item)
			internalBlock = append(internalBlock, dns)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "network",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) NetworkPort(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) NetworkDNS(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "servers", "searches", "options":
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

func (s *Structure) Parameterized(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Periodic(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Proxy(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "local_service_address", "local_service_port":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "config" {
			config := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "proxy",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Reschedule(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"attempts",
		"interval",
		"delay",
		"delay_function",
		"max_delay",
		"unlimited",
	}

	for _, item := range block.Parameter {
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

func (s *Structure) Resources(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Restart(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Scaling(block model.ConfigBlock) model.TemplateBlock {
	var name string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "min", "max", "enabled":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "policy" {
			policy := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, policy)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "scaling",
		Name:      name,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Service(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
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

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		switch item.Name {
		case "tagged_addresses", "meta", "canary_meta":
			block := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "service",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) SidecarService(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "disable_default_tcp_check", "port", "tags":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "meta" {
			meta := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, meta)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "sidecar_service",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) SidecarTask(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"name",
		"driver",
		"user",
		"kill_timeout",
		"shutdown_delay",
		"kill_signal",
	}

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		switch item.Name {
		case "config", "env", "meta":
			block := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "sidecar_task",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Spread(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "attribute", "weight":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "target" {
			target := s.SpreadTarget(item)
			internalBlock = append(internalBlock, target)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "spread",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) SpreadTarget(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "value", "percent":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "target",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) Task(block model.ConfigBlock) model.TemplateBlock {
	var name string
	var internalBlock []model.TemplateBlock
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

	for _, item := range block.Parameter {
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

	for _, item := range block.Block {
		if item.Name == "config" {
			config := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "task",
		Name:      name,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Template(
	block model.ConfigBlock,
	projectPath string,
) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
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

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				// read file and add data into "data" parameter.
				filePath := fmt.Sprintf("%s/files/%s", projectPath, v)

				file, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatalf("error read template files - %v", err)
				}

				i := make(map[string]interface{})
				i["data"] = string(file)
				parameters = append(parameters, i)
			}

			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "wait" {
			wait := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, wait)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "template",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) Update(block model.ConfigBlock) model.TemplateBlock {
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

	for _, item := range block.Parameter {
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

func (s *Structure) Upstream(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"destination_name",
		"destination_namespace",
		"datacenter",
		"local_bind_address",
	}

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		switch item.Name {
		case "mesh_gateway":
			meshGateway := s.UpstreamMeshGateway(item)
			internalBlock = append(internalBlock, meshGateway)
		case "config":
			config := s.CustomBlock(item.Name, item.Parameter)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "upstream",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) UpstreamMeshGateway(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
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

func (s *Structure) Vault(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"change_mode",
		"change_signal",
		"env",
		"disable_file",
		"namespace",
		"policies",
	}

	for _, item := range block.Parameter {
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

func (s *Structure) Volume(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
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

	for _, item := range block.Parameter {
		for k := range item {
			for _, p := range parameterName {
				switch k {
				case p:
					parameters = append(parameters, item)
				}
			}
		}
	}

	for _, item := range block.Block {
		if item.Name == "mount_options" {
			mountOptions := s.VolumeMountOptions(item)
			internalBlock = append(internalBlock, mountOptions)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "volume",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (s *Structure) VolumeMountOptions(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k := range item {
			switch k {
			case "fs_type", "mount_flags":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "mount_options",
		Parameter: parameters,
	}

	return templateBlock
}

func (s *Structure) VolumeMount(block model.ConfigBlock) model.TemplateBlock {
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"volume",
		"destination",
		"read_only",
		"propagation_mode",
	}

	for _, item := range block.Parameter {
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
