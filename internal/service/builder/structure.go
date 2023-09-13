package builder

import (
	"prism/internal/model"
)

type StructureBuilder struct{}

func NewStructureBuilder() *StructureBuilder {
	return &StructureBuilder{}
}

// Get configuration block by nomad block name.
func getBlockByName(name string, config model.ConfigBlock) model.ConfigBlock {
	for _, block := range config.Block {
		if block.Name == name {
			return block
		}
	}

	return model.ConfigBlock{}
}

// Get configuration template list by nomad block name.
func getBlockList(name []string, config model.ConfigBlock) []model.TemplateBlock {
	builder := blockBuilder
	var block []model.TemplateBlock

	f := make(map[string]func(model.ConfigBlock) model.TemplateBlock)

	f["artifact"] = builder.Artifact
	f["affinity"] = builder.Affinity
	f["change_script"] = builder.ChangeScript
	f["check_restart"] = builder.CheckRestart
	f["constraint"] = builder.Constraint
	f["csi_plugin"] = builder.CSIPlugin
	f["device"] = builder.Device
	f["dispatch_payload"] = builder.DispatchPayload
	f["env"] = builder.Env
	f["ephemeral_disk"] = builder.EphemeralDisk
	f["expose"] = builder.Expose
	f["gateway"] = builder.Gateway
	f["identity"] = builder.Identity
	f["lifecycle"] = builder.Lifecycle
	f["logs"] = builder.Logs
	f["meta"] = builder.Meta
	f["migrate"] = builder.Migrate
	f["multiregion"] = builder.Multiregion
	f["network"] = builder.Network
	f["parameterized"] = builder.Parameterized
	f["periodic"] = builder.Periodic
	f["reschedule"] = builder.Reschedule
	f["restart"] = builder.Restart
	f["scaling"] = builder.Scaling
	f["spread"] = builder.Spread
	f["update"] = builder.Update
	f["upstream"] = builder.Upstream
	f["vault"] = builder.Vault
	f["volume"] = builder.Volume
	f["volume_mount"] = builder.VolumeMount

	for _, n := range name {
		for k, v := range f {
			if n == k {
				b := getBlockByName(n, config)
				t := v(b)

				if len(t.Parameter) != 0 || len(t.Block) != 0 {
					block = append(block, t)
				}
			}
		}
	}

	return block
}

// Builds and returns a job configuration structure.
func (s *StructureBuilder) BuildConfigStructure(
	job, chart model.ConfigBlock,
	projectDirPath string,
) model.TemplateBlock {
	pathProjectDir = projectDirPath
	return jobStructure(job, chart)
}

func jobStructure(config, chart model.ConfigBlock) model.TemplateBlock {
	job := blockBuilder.Job(config)
	job = changes.Job(job, chart)

	block := []string{
		"affinity",
		"constraint",
		"multiregion",
		"parameterized",
		"periodic",
		"migrate",
		"reschedule",
		"spread",
		"update",
		"meta",
		"vault",
	}

	blockList := getBlockList(block, config)
	for _, item := range blockList {
		if item.BlockName == "meta" {
			item = changes.Meta(item, chart)
		}

		job.Block = append(job.Block, item)
	}

	// Group.
	group := groupStructure(config)
	job.Block = append(job.Block, group...)

	return job
}

func groupStructure(config model.ConfigBlock) []model.TemplateBlock {
	var groupList []model.TemplateBlock

	for _, item := range config.Block {
		if item.Name == "group" {
			group := blockBuilder.Group(item)

			block := []string{"affinity", "constraint", "meta"}
			group.Block = append(group.Block, getBlockList(block, item)...)

			// service block.
			for _, block := range item.Block {
				if block.Name == "service" {
					service := serviceStructure(block)

					if len(service.Parameter) != 0 || len(service.Block) != 0 {
						group.Block = append(group.Block, service)
					}
				}
			}

			block = []string{"restart", "scaling"}
			group.Block = append(group.Block, getBlockList(block, item)...)

			// task block.
			taskConfig := getBlockByName("task", item)
			task := taskStructure(taskConfig)

			if len(task.Parameter) != 0 || len(task.Block) != 0 {
				group.Block = append(group.Block, task)
			}

			block = []string{
				"vault",
				"volume",
				"migrate",
				"reschedule",
				"spread",
				"update",
			}
			group.Block = append(group.Block, getBlockList(block, item)...)

			network := networkStructure(item)
			if len(network.Block) != 0 {
				group.Block = append(group.Block, network)
			}

			groupList = append(groupList, group)
		}
	}

	return groupList
}

func networkStructure(config model.ConfigBlock) model.TemplateBlock {
	var network model.TemplateBlock
	var networkParameter []map[string]interface{}
	var networkBlock []model.ConfigBlock

	for _, p := range config.Parameter {
		for k, v := range p {
			switch k {
			case "network_mode":
				i := make(map[string]interface{})
				i["mode"] = v
				networkParameter = append(networkParameter, i)
			case "network_hostname":
				i := make(map[string]interface{})
				i["hostname"] = v
				networkParameter = append(networkParameter, i)
			}

		}
	}

	for _, b := range config.Block {
		if b.Name == "network" {
			for _, item := range b.Block {
				configBlock := model.ConfigBlock{
					Name:      item.Name,
					Parameter: item.Parameter,
					Block:     item.Block,
				}

				networkBlock = append(networkBlock, configBlock)
			}
		}
	}

	networkConfig := model.ConfigBlock{
		Name:      "network",
		Parameter: networkParameter,
		Block:     networkBlock,
	}

	network = blockBuilder.Network(networkConfig)
	return network
}

func taskStructure(config model.ConfigBlock) model.TemplateBlock {
	task := blockBuilder.Task(config)

	block := []string{
		"affinity",
		"artifact",
		"constraint",
		"csi_plugin",
		"dispatch_payload",
		"env",
		"ephemeral_disk",
		"identity",
		"lifecycle",
		"logs",
		"meta",
		"restart",
		"scaling",
		"vault",
		"volume_mount",
	}

	task.Block = append(task.Block, getBlockList(block, config)...)

	// template.
	for _, item := range config.Block {
		if item.Name == "template" {
			template := templateStructure(item, pathProjectDir)

			if len(template.Parameter) != 0 || len(template.Block) != 0 {
				task.Block = append(task.Block, template)
			}
		}
	}

	// service.
	for _, block := range config.Block {
		if block.Name == "service" {
			service := serviceStructure(block)

			if len(service.Parameter) != 0 || len(service.Block) != 0 {
				task.Block = append(task.Block, service)
			}
		}

	}

	// resources.
	resourcesConfig := getBlockByName("resources", config)
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
	block := []string{"change_script"}
	template.Block = append(template.Block, getBlockList(block, config)...)

	return template
}

func serviceStructure(config model.ConfigBlock) model.TemplateBlock {
	service := blockBuilder.Service(config)

	// check.
	checkConfig := getBlockByName("check", config)
	check := checkStructure(checkConfig)

	if len(check.Parameter) != 0 || len(check.Block) != 0 {
		service.Block = append(service.Block, check)
	}

	// check restart.
	block := []string{"check_restart"}
	service.Block = append(service.Block, getBlockList(block, config)...)

	// connect.
	connectConfig := getBlockByName("connect", config)
	connect := connectStructure(connectConfig)

	if len(connect.Parameter) != 0 || len(connect.Block) != 0 {
		service.Block = append(service.Block, connect)
	}

	return service
}

func checkStructure(config model.ConfigBlock) model.TemplateBlock {
	check := blockBuilder.Check(config)

	// check restart.
	block := []string{"check_restart"}
	check.Block = append(check.Block, getBlockList(block, config)...)

	return check
}

func connectStructure(config model.ConfigBlock) model.TemplateBlock {
	connect := blockBuilder.Connect(config)

	for _, item := range connect.Parameter {
		for k, v := range item {
			if k == "open_sidecar_service" {
				if v.(bool) {
					connect := model.TemplateBlock{
						BlockName: "connect",
					}

					sidecarService := model.TemplateBlock{
						BlockName: "sidecar_service",
					}

					connect.Block = append(connect.Block, sidecarService)
					return connect
				}
			}
		}
	}

	// sidecar service.
	sidecarServiceConfig := getBlockByName("sidecar_service", config)
	sidecarService := sidecarServiceStructure(sidecarServiceConfig)

	if len(sidecarService.Parameter) != 0 || len(sidecarService.Block) != 0 {
		connect.Block = append(connect.Block, sidecarService)
	}

	// sidecar task.
	sidecarTaskConfig := getBlockByName("sidecar_task", config)
	sidecarTask := sidecarTaskStructure(sidecarTaskConfig)

	if len(sidecarTask.Parameter) != 0 || len(sidecarTask.Block) != 0 {
		connect.Block = append(connect.Block, sidecarTask)
	}

	// gateway.
	block := []string{"gateway"}
	connect.Block = append(connect.Block, getBlockList(block, config)...)

	return connect
}

func sidecarServiceStructure(config model.ConfigBlock) model.TemplateBlock {
	sidecarService := blockBuilder.SidecarService(config)

	// proxy.
	proxyConfig := getBlockByName("proxy", config)
	proxy := proxyStructure(proxyConfig)

	if len(proxy.Parameter) != 0 || len(proxy.Block) != 0 {
		sidecarService.Block = append(sidecarService.Block, proxy)
	}

	return sidecarService
}

func proxyStructure(config model.ConfigBlock) model.TemplateBlock {
	proxy := blockBuilder.Proxy(config)

	// expose, upstream.
	block := []string{"expose", "upstream"}
	proxy.Block = append(proxy.Block, getBlockList(block, config)...)

	return proxy
}

func sidecarTaskStructure(config model.ConfigBlock) model.TemplateBlock {
	sidecarTask := blockBuilder.SidecarTask(config)

	// logs.
	block := []string{"logs"}
	sidecarTask.Block = append(sidecarTask.Block, getBlockList(block, config)...)

	// resources.
	resourcesConfig := getBlockByName("resources", config)
	resources := resourcesStructure(resourcesConfig)

	if len(resources.Parameter) != 0 || len(resources.Block) != 0 {
		sidecarTask.Block = append(sidecarTask.Block, resources)
	}

	return sidecarTask
}

func resourcesStructure(config model.ConfigBlock) model.TemplateBlock {
	resources := blockBuilder.Resources(config)

	// device.
	deviceConfig := getBlockByName("device", config)
	device := deviceStructure(deviceConfig)

	if len(deviceConfig.Parameter) != 0 || len(deviceConfig.Block) != 0 {
		resources.Block = append(resources.Block, device)
	}

	return resources
}

func deviceStructure(config model.ConfigBlock) model.TemplateBlock {
	device := blockBuilder.Device(config)

	block := []string{"affinity", "constraint"}
	device.Block = append(device.Block, getBlockList(block, config)...)

	return device
}
