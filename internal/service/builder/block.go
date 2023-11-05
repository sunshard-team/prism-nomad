// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"prism/internal/model"
	"regexp"
)

type BlockBuilder struct{}

func NewBlockBuilder() *BlockBuilder {
	return &BlockBuilder{}
}

// Returns a block with any key-value parameters.
func (b *BlockBuilder) CustomBlock(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		Type:      block.Type,
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
		if item.Type == "options" || item.Type == "headers" {
			block := b.CustomBlock(item)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "artifact",
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
		Type:      "affinity",
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
		Type:      "change_script",
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
		if item.Type == "header" {
			header := b.CustomBlock(item)
			internalBlock = append(internalBlock, header)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "check",
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
		Type:      "check_restart",
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
		Type:      "connect",
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
		Type:      "constraint",
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
		Type:      "csi_plugin",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Device(block model.ConfigBlock) model.TemplateBlock {
	var label string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				label = v.(string)
			case "count":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		Label:     label,
		Type:      "device",
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
		Type:      "dispatch_payload",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Env(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		Type:      "env",
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
		Type:      "ephemeral_disk",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Expose(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Type == "path" {
			path := b.ExposePath(item)
			internalBlock = append(internalBlock, path)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:  "expose",
		Block: internalBlock,
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
		Type:      "path",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Gateway(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Type {
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
		Type:  "gateway",
		Block: internalBlock,
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
		switch item.Type {
		case "envoy_gateway_bind_addresses":
			address := b.GatewayProxyAddress(item)
			internalBlock = append(internalBlock, address)
		case "config":
			config := b.CustomBlock(item)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "proxy",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayProxyAddress(block model.ConfigBlock) model.TemplateBlock {
	var label string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				label = v.(string)
			case "address", "port":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "envoy_gateway_bind_addresses",
		Label:     label,
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayIngress(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Type {
		case "tls":
			tls := b.GatewayIngressTLS(item)
			internalBlock = append(internalBlock, tls)
		case "listener":
			listener := b.GatewayIngressListener(item)
			internalBlock = append(internalBlock, listener)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:  "ingress",
		Block: internalBlock,
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
		Type:      "tls",
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
		if item.Type == "service" {
			service := b.GatewayIngressListenerService(item)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "listener",
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
		Type:      "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayTerminating(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		if item.Type == "service" {
			service := b.GatewayTerminatingService(item)
			internalBlock = append(internalBlock, service)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:  "terminating",
		Block: internalBlock,
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
		Type:      "service",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) GatewayMesh() model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		Type: "mesh",
	}

	return templateBlock
}

func (b *BlockBuilder) Group(block model.ConfigBlock) model.TemplateBlock {
	var label string
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
				label = v.(string)
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
		if item.Type == "consul" {
			consul := b.GroupConsul(item)
			internalBlock = append(internalBlock, consul)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "group",
		Label:     label,
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
		Type:      "consul",
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
		Type:      "identity",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Job(block model.ConfigBlock) model.TemplateBlock {
	var label string
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
				label = v.(string)
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
		Type:      "job",
		Label:     label,
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
		Type:      "lifecycle",
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
		Type:      "logs",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Meta(block model.ConfigBlock) model.TemplateBlock {
	templateBlock := model.TemplateBlock{
		Type:      "meta",
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
		Type:      "migrate",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Multiregion(block model.ConfigBlock) model.TemplateBlock {
	var internalBlock []model.TemplateBlock

	for _, item := range block.Block {
		switch item.Type {
		case "strategy":
			strategy := b.MultiregionStrategy(item)
			internalBlock = append(internalBlock, strategy)
		case "region":
			region := b.MultiregionRegion(item)
			internalBlock = append(internalBlock, region)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:  "multiregion",
		Block: internalBlock,
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
		Type:      "strategy",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) MultiregionRegion(block model.ConfigBlock) model.TemplateBlock {
	var label string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				label = v.(string)
			case "count", "datacenters", "node_pool":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Type == "meta" {
			meta := b.CustomBlock(item)
			internalBlock = append(internalBlock, meta)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "region",
		Label:     label,
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
		switch item.Type {
		case "port":
			port := b.NetworkPort(item)
			internalBlock = append(internalBlock, port)
		case "dns":
			dns := b.NetworkDNS(item)
			internalBlock = append(internalBlock, dns)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "network",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) NetworkPort(block model.ConfigBlock) model.TemplateBlock {
	var label string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				label = v.(string)
			case "static", "to", "host_network":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "port",
		Label:     label,
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
		Type:      "dns",
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
		Type:      "parameterized",
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
		Type:      "periodic",
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
		if item.Type == "config" {
			config := b.CustomBlock(item)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "proxy",
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
		Type:      "reschedule",
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
		Type:      "resources",
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
		Type:      "restart",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Scaling(block model.ConfigBlock) model.TemplateBlock {
	var label string
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "name":
				label = v.(string)
			case "min", "max", "enabled":
				parameters = append(parameters, item)
			}
		}
	}

	for _, item := range block.Block {
		if item.Type == "policy" {
			policy := b.CustomBlock(item)
			internalBlock = append(internalBlock, policy)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "scaling",
		Label:     label,
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
		switch item.Type {
		case "tagged_addresses", "meta", "canary_meta":
			block := b.CustomBlock(item)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "service",
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
		if item.Type == "meta" {
			meta := b.CustomBlock(item)
			internalBlock = append(internalBlock, meta)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "sidecar_service",
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
		switch item.Type {
		case "config", "env", "meta":
			block := b.CustomBlock(item)
			internalBlock = append(internalBlock, block)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "sidecar_task",
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
		if item.Type == "target" {
			target := b.SpreadTarget(item)
			internalBlock = append(internalBlock, target)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "spread",
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) SpreadTarget(block model.ConfigBlock) model.TemplateBlock {
	var label string
	parameters := make([]map[string]interface{}, 0)

	for _, item := range block.Parameter {
		for k, v := range item {
			switch k {
			case "value":
				label = v.(string)
			case "percent":
				parameters = append(parameters, item)
			}
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "target",
		Label:     label,
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Task(block model.ConfigBlock) model.TemplateBlock {
	var label string
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
				label = v.(string)
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
		if item.Type == "config" {
			config := b.CustomBlock(item)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "task",
		Label:     label,
		Parameter: parameters,
		Block:     internalBlock,
	}

	return templateBlock
}

func (b *BlockBuilder) Template(
	block model.ConfigBlock,
	fileDirPath string,
) model.TemplateBlock {
	var internalBlock []model.TemplateBlock
	parameters := make([]map[string]interface{}, 0)

	parameterName := []string{
		"name",
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
			case "file":
				var fileFullPath string

				// Check the full file path or file name.
				separatorFormat, err := regexp.Compile(`\\|\/`)
				if err != nil {
					fmt.Printf("failed check OS separator in file path, %s", err)
					os.Exit(1)
				}

				findSeparator := separatorFormat.FindStringSubmatch(v.(string))

				if len(findSeparator) > 0 {
					fileFullPath = v.(string)
				} else {
					fileFullPath = filepath.Join(fileDirPath, v.(string))
				}

				// Read the file and add data to the "data" parameter.
				file, err := os.ReadFile(fileFullPath)
				if err != nil {
					fmt.Printf("error read template files - %v\n", err)
					os.Exit(1)
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
		if item.Type == "wait" {
			wait := b.CustomBlock(item)
			internalBlock = append(internalBlock, wait)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "template",
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
		Type:      "update",
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
		switch item.Type {
		case "mesh_gateway":
			meshGateway := b.UpstreamMeshGateway(item)
			internalBlock = append(internalBlock, meshGateway)
		case "config":
			config := b.CustomBlock(item)
			internalBlock = append(internalBlock, config)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "upstreams",
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
		Type:      "mesh_gateway",
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
		Type:      "vault",
		Parameter: parameters,
	}

	return templateBlock
}

func (b *BlockBuilder) Volume(block model.ConfigBlock) model.TemplateBlock {
	var label string
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
				label = v.(string)
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
		if item.Type == "mount_options" {
			mountOptions := b.VolumeMountOptions(item)
			internalBlock = append(internalBlock, mountOptions)
		}
	}

	templateBlock := model.TemplateBlock{
		Type:      "volume",
		Label:     label,
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
		Type:      "mount_options",
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
		Type:      "volume_mount",
		Parameter: parameters,
	}

	return templateBlock
}
