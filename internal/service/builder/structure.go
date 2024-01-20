// Copyright (c) 2023 SUNSHARD
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package builder

import (
	"prism/internal/model"
	"prism/pkg"
)

var (
	blockBuilder BlockBuilder
	filesDirPath string
)

type StructureBuilder struct {
	blockBuilder BlockBuilder
}

func NewStructureBuilder(blockBuilder BlockBuilder) *StructureBuilder {
	return &StructureBuilder{blockBuilder: blockBuilder}
}

// Builds and returns a job configuration structure.
func (s *StructureBuilder) BuildConfigStructure(
	buildStructure model.BuildStructure,
) model.TemplateBlock {
	blockBuilder = s.blockBuilder
	filesDirPath = buildStructure.FilesDirPath
	return jobStructure(buildStructure.Config)
}

// Get configuration block by nomad block name.
func getBlockByType(
	blockType string,
	config model.ConfigBlock,
) model.ConfigBlock {
	for _, block := range config.Block {
		if block.Type == blockType {
			return block
		}
	}

	return model.ConfigBlock{}
}

// Get configuration template list by nomad block name.
func getConfigBlock(
	config model.ConfigBlock,
	configBlock map[string]func(model.ConfigBlock) model.TemplateBlock,
) []model.TemplateBlock {
	var block []model.TemplateBlock

	for k, v := range configBlock {
		b := getBlockByType(k, config)
		t := v(b)

		if len(t.Parameter) != 0 || len(t.Block) != 0 {
			block = append(block, t)
		}
	}

	return block
}

func jobStructure(config model.ConfigBlock) model.TemplateBlock {
	job := blockBuilder.Job(config)

	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["affinity"] = blockBuilder.Affinity
	configBlock["constraint"] = blockBuilder.Constraint
	configBlock["meta"] = blockBuilder.Meta
	configBlock["parameterized"] = blockBuilder.Parameterized
	configBlock["periodic"] = blockBuilder.Periodic
	configBlock["migrate"] = blockBuilder.Migrate
	configBlock["reschedule"] = blockBuilder.Reschedule
	configBlock["update"] = blockBuilder.Update
	configBlock["vault"] = blockBuilder.Vault

	blockList := getConfigBlock(config, configBlock)
	job.Block = append(job.Block, blockList...)

	// multiregion, set job block.
	multiregion := multiregionStructure(config)
	if len(multiregion.Block) != 0 {
		job.Block = append(job.Block, multiregion)
	}

	// spread, set job block.
	spread := spreadStructure(config)
	if len(spread.Block) != 0 {
		job.Block = append(job.Block, spread)
	}

	// group.
	for _, block := range config.Block {
		if block.Type == "group" {
			group := groupStructure(block)
			job.Block = append(job.Block, group)
		}
	}

	return job
}

func multiregionStructure(config model.ConfigBlock) model.TemplateBlock {
	var multiregionBlock []model.ConfigBlock

	for _, block := range config.Block {
		if block.Type == "multiregion" {
			for _, item := range block.Block {
				configBlock := model.ConfigBlock{
					Type:      item.Type,
					Parameter: item.Parameter,
					Block:     item.Block,
				}

				multiregionBlock = append(multiregionBlock, configBlock)
			}
		}
	}

	multiregionConfig := model.ConfigBlock{
		Type:  "multiregion",
		Block: multiregionBlock,
	}

	multiregion := blockBuilder.Multiregion(multiregionConfig)
	return multiregion
}

func groupStructure(config model.ConfigBlock) model.TemplateBlock {
	group := blockBuilder.Group(config)

	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["affinity"] = blockBuilder.Affinity
	configBlock["constraint"] = blockBuilder.Constraint
	configBlock["meta"] = blockBuilder.Meta
	configBlock["restart"] = blockBuilder.Restart
	configBlock["vault"] = blockBuilder.Vault
	configBlock["ephemeral_disk"] = blockBuilder.EphemeralDisk
	configBlock["migrate"] = blockBuilder.Migrate
	configBlock["reschedule"] = blockBuilder.Reschedule
	configBlock["update"] = blockBuilder.Update

	group.Block = append(
		group.Block,
		getConfigBlock(config, configBlock)...,
	)

	// scaling.
	for _, block := range config.Block {
		if block.Type == "scaling" {
			scaling := blockBuilder.Scaling(block)

			if len(scaling.Parameter) != 0 || len(scaling.Block) != 0 {
				group.Block = append(group.Block, scaling)
			}
		}
	}

	// volume block.
	for _, block := range config.Block {
		if block.Type == "volume" {
			volume := blockBuilder.Volume(block)

			if len(volume.Parameter) != 0 || len(volume.Block) != 0 {
				group.Block = append(group.Block, volume)
			}
		}
	}

	// service block.
	for _, block := range config.Block {
		if block.Type == "service" {
			service := serviceStructure(block)

			if len(service.Parameter) != 0 || len(service.Block) != 0 {
				group.Block = append(group.Block, service)
			}
		}
	}

	// task block.
	for _, block := range config.Block {
		if block.Type == "task" {
			task := taskStructure(block)

			if len(task.Parameter) != 0 || len(task.Block) != 0 {
				group.Block = append(group.Block, task)
			}
		}
	}

	// network, set group block.
	network := networkStructure(config)
	if len(network.Parameter) != 0 || len(network.Block) != 0 {
		group.Block = append(group.Block, network)
	}

	// spread, set group block.
	spread := spreadStructure(config)
	if len(spread.Block) != 0 {
		group.Block = append(group.Block, network)
	}

	return group
}

func spreadStructure(config model.ConfigBlock) model.TemplateBlock {
	var spreadParameter []map[string]interface{}
	var spreadBlock []model.ConfigBlock

	for _, block := range config.Block {
		if block.Type == "spread" {
			spreadParameter = append(spreadParameter, block.Parameter...)

			for _, item := range block.Block {
				configBlock := model.ConfigBlock{
					Type:      item.Type,
					Parameter: item.Parameter,
					Block:     item.Block,
				}

				spreadBlock = append(spreadBlock, configBlock)
			}
		}
	}

	spreadConfig := model.ConfigBlock{
		Type:      "spread",
		Parameter: spreadParameter,
		Block:     spreadBlock,
	}

	spread := blockBuilder.Spread(spreadConfig)
	return spread
}

func networkStructure(config model.ConfigBlock) model.TemplateBlock {
	var networkParameter []map[string]interface{}
	var networkBlock []model.ConfigBlock

	for _, block := range config.Block {
		if block.Type == "network" {
			networkParameter = append(networkParameter, block.Parameter...)

			for _, item := range block.Block {
				configBlock := model.ConfigBlock{
					Type:      item.Type,
					Parameter: item.Parameter,
					Block:     item.Block,
				}

				networkBlock = append(networkBlock, configBlock)
			}
		}
	}

	networkConfig := model.ConfigBlock{
		Type:      "network",
		Parameter: networkParameter,
		Block:     networkBlock,
	}

	network := blockBuilder.Network(networkConfig)
	return network
}

func taskStructure(config model.ConfigBlock) model.TemplateBlock {
	task := blockBuilder.Task(config)

	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["artifact"] = blockBuilder.Artifact
	configBlock["affinity"] = blockBuilder.Affinity
	configBlock["constraint"] = blockBuilder.Constraint
	configBlock["csi_plugin"] = blockBuilder.CSIPlugin
	configBlock["dispatch_payload"] = blockBuilder.DispatchPayload
	configBlock["env"] = blockBuilder.Env
	configBlock["identity"] = blockBuilder.Identity
	configBlock["lifecycle"] = blockBuilder.Lifecycle
	configBlock["logs"] = blockBuilder.Logs
	configBlock["meta"] = blockBuilder.Meta
	configBlock["restart"] = blockBuilder.Restart
	configBlock["vault"] = blockBuilder.Vault

	task.Block = append(
		task.Block,
		getConfigBlock(config, configBlock)...,
	)

	// scaling.
	for _, block := range config.Block {
		if block.Type == "scaling" {
			scaling := blockBuilder.Scaling(block)

			if len(scaling.Parameter) != 0 || len(scaling.Block) != 0 {
				task.Block = append(task.Block, scaling)
			}
		}
	}

	// volume mount.
	for _, block := range config.Block {
		if block.Type == "volume_mount" {
			volumeMount := blockBuilder.VolumeMount(block)

			if len(volumeMount.Parameter) != 0 || len(volumeMount.Block) != 0 {
				task.Block = append(task.Block, volumeMount)
			}
		}
	}

	// template.
	for _, block := range config.Block {
		if block.Type == "template" {
			template := templateStructure(block, filesDirPath)

			if len(template.Parameter) != 0 || len(template.Block) != 0 {
				task.Block = append(task.Block, template)
			}
		}
	}

	// service.
	for _, block := range config.Block {
		if block.Type == "service" {
			service := serviceStructure(block)

			if len(service.Parameter) != 0 || len(service.Block) != 0 {
				task.Block = append(task.Block, service)
			}
		}
	}

	// resources.
	resourcesConfig := getBlockByType("resources", config)
	resources := resourcesStructure(resourcesConfig)

	if len(resources.Parameter) != 0 || len(resources.Block) != 0 {
		task.Block = append(task.Block, resources)
	}

	return task
}

func templateStructure(
	config model.ConfigBlock,
	projectPath string,
) model.TemplateBlock {
	template := blockBuilder.Template(config, projectPath)

	// change script.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["change_script"] = blockBuilder.ChangeScript

	template.Block = append(
		template.Block,
		getConfigBlock(config, configBlock)...,
	)

	return template
}

func serviceStructure(config model.ConfigBlock) model.TemplateBlock {
	service := blockBuilder.Service(config)

	// check.
	for _, block := range config.Block {
		if block.Type == "check" {
			check := checkStructure(block)

			if len(check.Parameter) != 0 || len(check.Block) != 0 {
				service.Block = append(service.Block, check)
			}
		}
	}

	// check restart.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["check_restart"] = blockBuilder.CheckRestart

	service.Block = append(
		service.Block,
		getConfigBlock(config, configBlock)...,
	)

	// connect.
	connectConfig := getBlockByType("connect", config)
	connect := connectStructure(connectConfig)

	if len(connect.Parameter) != 0 || len(connect.Block) != 0 {
		service.Block = append(service.Block, connect)
	}

	return service
}

func checkStructure(config model.ConfigBlock) model.TemplateBlock {
	check := blockBuilder.Check(config)

	// check restart.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["check_restart"] = blockBuilder.CheckRestart

	check.Block = append(
		check.Block,
		getConfigBlock(config, configBlock)...,
	)

	return check
}

func connectStructure(config model.ConfigBlock) model.TemplateBlock {
	connect := blockBuilder.Connect(config)

	for _, item := range connect.Parameter {
		for k, v := range item {
			if k == "open_sidecar_service" {
				if v.(bool) {
					connect := model.TemplateBlock{
						Type: "connect",
					}

					sidecarService := model.TemplateBlock{
						Type: "sidecar_service",
					}

					connect.Block = append(connect.Block, sidecarService)
					return connect
				} else {
					pkg.RemoveParameter(&connect, "open_sidecar_service")
				}
			}
		}
	}

	// sidecar service.
	sidecarServiceConfig := getBlockByType("sidecar_service", config)
	sidecarService := sidecarServiceStructure(sidecarServiceConfig)

	if len(sidecarService.Parameter) != 0 || len(sidecarService.Block) != 0 {
		connect.Block = append(connect.Block, sidecarService)
	}

	// sidecar task.
	sidecarTaskConfig := getBlockByType("sidecar_task", config)
	sidecarTask := sidecarTaskStructure(sidecarTaskConfig)

	if len(sidecarTask.Parameter) != 0 || len(sidecarTask.Block) != 0 {
		connect.Block = append(connect.Block, sidecarTask)
	}

	// gateway.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["gateway"] = blockBuilder.Gateway

	connect.Block = append(
		connect.Block,
		getConfigBlock(config, configBlock)...,
	)

	return connect
}

func sidecarServiceStructure(config model.ConfigBlock) model.TemplateBlock {
	sidecarService := blockBuilder.SidecarService(config)

	// proxy.
	proxyConfig := getBlockByType("proxy", config)
	proxy := proxyStructure(proxyConfig)

	if len(proxy.Parameter) != 0 || len(proxy.Block) != 0 {
		sidecarService.Block = append(sidecarService.Block, proxy)
	}

	return sidecarService
}

func proxyStructure(config model.ConfigBlock) model.TemplateBlock {
	proxy := blockBuilder.Proxy(config)

	// expose, upstreams.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["expose"] = blockBuilder.Expose
	configBlock["upstreams"] = blockBuilder.Upstreams

	proxy.Block = append(
		proxy.Block,
		getConfigBlock(config, configBlock)...,
	)

	return proxy
}

func sidecarTaskStructure(config model.ConfigBlock) model.TemplateBlock {
	sidecarTask := blockBuilder.SidecarTask(config)

	// logs.
	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["logs"] = blockBuilder.Logs

	sidecarTask.Block = append(
		sidecarTask.Block,
		getConfigBlock(config, configBlock)...,
	)

	// resources.
	resourcesConfig := getBlockByType("resources", config)
	resources := resourcesStructure(resourcesConfig)

	if len(resources.Parameter) != 0 || len(resources.Block) != 0 {
		sidecarTask.Block = append(sidecarTask.Block, resources)
	}

	return sidecarTask
}

func resourcesStructure(config model.ConfigBlock) model.TemplateBlock {
	resources := blockBuilder.Resources(config)

	// device.
	for _, block := range config.Block {
		if block.Type == "device" {
			device := deviceStructure(block)

			if len(device.Parameter) != 0 || len(device.Block) != 0 {
				resources.Block = append(resources.Block, device)
			}
		}
	}

	return resources
}

func deviceStructure(config model.ConfigBlock) model.TemplateBlock {
	device := blockBuilder.Device(config)

	configBlock := make(
		map[string]func(model.ConfigBlock) model.TemplateBlock,
	)

	configBlock["affinity"] = blockBuilder.Affinity
	configBlock["constraint"] = blockBuilder.Constraint

	device.Block = append(
		device.Block,
		getConfigBlock(config, configBlock)...,
	)

	return device
}
