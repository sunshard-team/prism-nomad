package builder

import (
	"log"
	"os"
	"path/filepath"
	"prism/internal/model"
)

type BlockBuilder struct{}

func NewBlockBuilder() *BlockBuilder {
	return &BlockBuilder{}
}

// Returns a block with any key-value parameters.
func (b *BlockBuilder) CustomBlock(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: block.Name,
		Parameter: block.Parameter,
	}

	return templateBlock
}

func (b *BlockBuilder) Artifact(block model.ConfigBlock) model.TemplateBlock {
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
			block := b.CustomBlock(item)
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

func (b *BlockBuilder) Affinity(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) ChangeScript(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Check(block model.ConfigBlock) model.TemplateBlock {
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
			header := b.CustomBlock(item)
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

func (b *BlockBuilder) CheckRestart(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Connect(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Constraint(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) CSIPlugin(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Device(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) DispatchPayload(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Env(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "env",
		Parameter: block.Parameter,
	}

	return templateBlock
}

func (b *BlockBuilder) EphemeralDisk(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Expose(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Name == "path" {
			path := b.ExposePath(item)
			internalBlock = append(internalBlock, path)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "expose",
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) ExposePath(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Gateway(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "proxy":
			proxy := b.GatewayProxy(item)
			internalBlock = append(internalBlock, proxy)
		case "ingress":
			ingress := b.GatewayIngress(item)
			internalBlock = append(internalBlock, ingress)
		case "terminating":
			terminating := b.GatewayTerminating(item)
			internalBlock = append(internalBlock, terminating)
		case "mesh":
			mesh := b.GatewayMesh()
			internalBlock = append(internalBlock, mesh)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "gateway",
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayProxy(block model.ConfigBlock) model.TemplateBlock {
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
			address := b.GatewayProxyAddress(item)
			internalBlock = append(internalBlock, address)
		case "config":
			config := b.CustomBlock(item)
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

func (b *BlockBuilder) GatewayProxyAddress(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				name = v.(string)
			case "address", "port":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "envoy_gateway_bind_addresses",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayIngress(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "tls":
			tls := b.GatewayIngressTLS(item)
			internalBlock = append(internalBlock, tls)
		case "listener":
			listener := b.GatewayIngressListener(item)
			internalBlock = append(internalBlock, listener)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "ingress",
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayIngressTLS(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) GatewayIngressListener(block model.ConfigBlock) model.TemplateBlock {
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
			service := b.GatewayIngressListenerService(item)
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

func (b *BlockBuilder) GatewayIngressListenerService(
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

func (b *BlockBuilder) GatewayTerminating(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Name == "service" {
			service := b.GatewayTerminatingService(item)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "terminating",
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayTerminatingService(
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

func (b *BlockBuilder) GatewayMesh() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "mesh",
	}

	return templateBlock
}

func (b *BlockBuilder) Group(block model.ConfigBlock) model.TemplateBlock {
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
			consul := b.GroupConsul(item)
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

func (b *BlockBuilder) GroupConsul(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Identity(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Job(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Lifecycle(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Logs(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Meta(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		BlockName: "meta",
		Parameter: block.Parameter,
	}

	return templateBlock
}

func (b *BlockBuilder) Migrate(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Multiregion(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Name {
		case "strategy":
			strategy := b.MultiregionStrategy(item)
			internalBlock = append(internalBlock, strategy)
		case "region":
			region := b.MultiregionRegion(item)
			internalBlock = append(internalBlock, region)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "multiregion",
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) MultiregionStrategy(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) MultiregionRegion(block model.ConfigBlock) model.TemplateBlock {
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
			meta := b.CustomBlock(item)
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

func (b *BlockBuilder) Network(block model.ConfigBlock) model.TemplateBlock {
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
			port := b.NetworkPort(item)
			internalBlock = append(internalBlock, port)
		case "dns":
			dns := b.NetworkDNS(item)
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

func (b *BlockBuilder) NetworkPort(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) NetworkDNS(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Parameterized(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Periodic(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Proxy(block model.ConfigBlock) model.TemplateBlock {
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
			config := b.CustomBlock(item)
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

func (b *BlockBuilder) Reschedule(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Resources(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Restart(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Scaling(block model.ConfigBlock) model.TemplateBlock {
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
			policy := b.CustomBlock(item)
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

func (b *BlockBuilder) Service(block model.ConfigBlock) model.TemplateBlock {
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
			block := b.CustomBlock(item)
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

func (b *BlockBuilder) SidecarService(block model.ConfigBlock) model.TemplateBlock {
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
			meta := b.CustomBlock(item)
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

func (b *BlockBuilder) SidecarTask(block model.ConfigBlock) model.TemplateBlock {
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
			block := b.CustomBlock(item)
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

func (b *BlockBuilder) Spread(block model.ConfigBlock) model.TemplateBlock {
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
			target := b.SpreadTarget(item)
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

func (b *BlockBuilder) SpreadTarget(block model.ConfigBlock) model.TemplateBlock {
	var name string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "value":
				name = v.(string)
			case "percent":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "target",
		Name:      name,
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Task(block model.ConfigBlock) model.TemplateBlock {
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
			config := b.CustomBlock(item)
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

func (b *BlockBuilder) Template(
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
				filePath := filepath.Join(projectPath, "files", v.(string))

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
			wait := b.CustomBlock(item)
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

func (b *BlockBuilder) Update(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Upstreams(block model.ConfigBlock) model.TemplateBlock {
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
			meshGateway := b.UpstreamMeshGateway(item)
			internalBlock = append(internalBlock, meshGateway)
		case "config":
			config := b.CustomBlock(item)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "upstreams",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) UpstreamMeshGateway(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Vault(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) Volume(block model.ConfigBlock) model.TemplateBlock {
	var name string
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
		if item.Name == "mount_options" {
			mountOptions := b.VolumeMountOptions(item)
			internalBlock = append(internalBlock, mountOptions)
		}
	}

	templateBlock := model.TemplateBlock{
		BlockName: "volume",
		Name:      name,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) VolumeMountOptions(block model.ConfigBlock) model.TemplateBlock {
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

func (b *BlockBuilder) VolumeMount(block model.ConfigBlock) model.TemplateBlock {
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
