package builder

import (
	"prism/internal/model"
)

type BlockBuilder interface {
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
	Gateway(block model.ConfigBlock) model.TemplateBlock
	Group(block model.ConfigBlock) model.TemplateBlock
	Identity(block model.ConfigBlock) model.TemplateBlock
	Job(block model.ConfigBlock, chart map[string]interface{}) model.TemplateBlock
	Lifecycle(block model.ConfigBlock) model.TemplateBlock
	Logs(block model.ConfigBlock) model.TemplateBlock
	Meta(block model.ConfigBlock) model.TemplateBlock
	Migrate(block model.ConfigBlock) model.TemplateBlock
	Multiregion(block model.ConfigBlock) model.TemplateBlock
	Network(block model.ConfigBlock) model.TemplateBlock
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
	Task(block model.ConfigBlock) model.TemplateBlock
	Template(block model.ConfigBlock) model.TemplateBlock
	Update(block model.ConfigBlock) model.TemplateBlock
	Upstream(block model.ConfigBlock) model.TemplateBlock
	Vault(block model.ConfigBlock) model.TemplateBlock
	Volume(block model.ConfigBlock) model.TemplateBlock
	VolumeMount(block model.ConfigBlock) model.TemplateBlock
}

type Block struct{}

func (b *Block) Artifact(block model.ConfigBlock) model.TemplateBlock {
	artifacat := structBuilder.Artifact(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "options" || b.Name == "headers" {
			options := structBuilder.CustomBlock(b.Name, b.Parameter)
			artifacat.Block = append(artifacat.Block, options)
		}
	}

	return artifacat
}

func (b *Block) Affinity(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Affinity(block.Parameter)
}

func (b *Block) ChangeScript(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.ChangeScript(block.Parameter)
}

func (b *Block) Check(block model.ConfigBlock) model.TemplateBlock {
	check := structBuilder.Check(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "header" {
			header := structBuilder.CustomBlock(b.Name, b.Parameter)
			check.Block = append(check.Block, header)
		}
	}

	return check
}

func (b *Block) CheckRestart(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.CheckRestart(block.Parameter)
}

func (b *Block) Connect(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Connect(block.Parameter)
}

func (b *Block) Constraint(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Constraint(block.Parameter)
}

func (b *Block) CSIPlugin(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.CSIPlugin(block.Parameter)
}

func (b *Block) Device(block model.ConfigBlock) model.TemplateBlock {
	device := structBuilder.Device(block.Parameter)

	for _, p := range block.Parameter {
		for k, v := range p {
			if k == "constraint" {
				constraint := structBuilder.Constraint(v.([]map[string]interface{}))
				device.Block = append(device.Block, constraint)
			}

			if k == "affinity" {
				affinity := structBuilder.Affinity(v.([]map[string]interface{}))
				device.Block = append(device.Block, affinity)
			}
		}
	}

	return device
}

func (b *Block) DispatchPayload(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.DispatchPayload(block.Parameter)
}

func (b *Block) Env(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Env(block.Parameter)
}

func (b *Block) EphemeralDisk(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.EphemeralDisk(block.Parameter)
}

func (b *Block) Expose(block model.ConfigBlock) model.TemplateBlock {
	expose := structBuilder.Expose()

	for _, p := range block.Parameter {
		for k, v := range p {
			if k == "path" {
				path := structBuilder.ExposePath(v.(map[string]interface{}))
				expose.Block = append(expose.Block, path)
			}
		}
	}

	return expose
}

func (b *Block) Gateway(block model.ConfigBlock) model.TemplateBlock {
	gateway := structBuilder.Gateway()

	for _, b := range block.Block {
		if b.Name == "proxy" {
			proxy := structBuilder.GatewayProxy(b.Parameter)

			for _, b := range b.Block {
				if b.Name == "envoy_gateway_bind_addresses" {
					address := structBuilder.GatewayProxyAddress(
						b.Parameter,
					)
					proxy.Block = append(proxy.Block, address)
				}

				if b.Name == "config" {
					config := structBuilder.CustomBlock(
						b.Name,
						b.Parameter,
					)
					proxy.Block = append(proxy.Block, config)
				}
			}

			gateway.Block = append(gateway.Block, proxy)
		}

		if b.Name == "ingress" {
			ingress := structBuilder.GatewayIngress()

			for _, b := range b.Block {
				if b.Name == "tls" {
					tls := structBuilder.GatewayIngressTLS(b.Parameter)
					ingress.Block = append(ingress.Block, tls)
				}

				if b.Name == "listener" {
					for _, b := range b.Block {
						listener := structBuilder.GatewayIngressListener(
							b.Parameter,
						)

						for _, b := range b.Block {
							if b.Name == "service" {
								service := structBuilder.GatewayIngressListenerService(
									b.Parameter,
								)
								listener.Block = append(listener.Block, service)
							}
						}

						ingress.Block = append(ingress.Block, listener)
					}
				}
			}

			gateway.Block = append(gateway.Block, ingress)
		}

		if b.Name == "terminating" {
			terminating := structBuilder.GatewayTerminating()

			for _, b := range b.Block {
				if b.Name == "service" {
					service := structBuilder.GatewayTerminatingService(
						b.Parameter,
					)
					terminating.Block = append(terminating.Block, service)
				}
			}

			gateway.Block = append(gateway.Block, terminating)
		}

		if b.Name == "mesh" {
			mesh := structBuilder.GatewayMesh()
			gateway.Block = append(gateway.Block, mesh)
		}
	}

	return gateway
}

func (b *Block) Group(block model.ConfigBlock) model.TemplateBlock {
	group := structBuilder.Group(block.Parameter)
Loop:
	for _, p := range group.Parameter {
		for k := range p {
			if k == "consul" {
				consul := structBuilder.GroupConsul(p)
				group.Block = append(group.Block, consul)
				break Loop
			}
		}
	}

	return group
}

func (b *Block) Identity(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Identity(block.Parameter)
}

func (b *Block) Job(
	block model.ConfigBlock,
	chart map[string]interface{},
) model.TemplateBlock {
	job := structBuilder.Job(block)

	if len(chart) > 0 {
	Loop:
		for key, value := range chart {
			if key == "type" {
				for i := range job.Parameter {
					for k := range job.Parameter[i] {
						if k == key {
							job.Parameter[i][k] = value
							break Loop
						}
					}
				}
			}
		}
	}

	return job
}

func (b *Block) Lifecycle(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Lifecycle(block.Parameter)
}

func (b *Block) Logs(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Logs(block.Parameter)
}

func (b *Block) Meta(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Meta(block.Parameter)
}

func (b *Block) Migrate(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Meta(block.Parameter)
}

func (b *Block) Multiregion(block model.ConfigBlock) model.TemplateBlock {
	multiregion := structBuilder.Multiregion()

	for _, p := range block.Block {
		for _, item := range p.Parameter {
			for k, v := range item {

				if k == "strategy" {
					strategy := structBuilder.MultiregionStrategy(item)
					multiregion.Block = append(multiregion.Block, strategy)
				}

				if k == "region" {
					for _, i := range v.([]interface{}) {
						region := structBuilder.MultiregionRegion(
							i.(map[string]interface{}),
						)

						for _, p := range region.Parameter {
							for k, v := range p {
								if k == "meta" {
									meta := structBuilder.Meta(
										v.([]map[string]interface{}),
									)
									region.Block = append(region.Block, meta)
								}
							}
						}

						multiregion.Block = append(multiregion.Block, region)
					}
				}
			}
		}
	}

	return multiregion
}

func (b *Block) Network(block model.ConfigBlock) model.TemplateBlock {
	network := structBuilder.Network(block.Parameter)

	for _, item := range block.Block {
		if item.Name == "port" {
			port := structBuilder.NetworkPort(item.Parameter)
			network.Block = append(network.Block, port)
		}

		if item.Name == "dns" {
			DNS := structBuilder.NetworkDNS(item.Parameter)
			network.Block = append(network.Block, DNS)
		}
	}

	return network
}

func (b *Block) Parameterized(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Parameterized(block.Parameter)
}

func (b *Block) Periodic(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Periodic(block.Parameter)
}

func (b *Block) Proxy(block model.ConfigBlock) model.TemplateBlock {
	proxy := structBuilder.Proxy(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "config" {
			config := structBuilder.CustomBlock(b.Name, b.Parameter)
			proxy.Block = append(proxy.Block, config)
		}
	}

	return proxy
}

func (b *Block) Reschedule(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Reschedule(block.Parameter)
}

func (b *Block) Resources(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Resources(block.Parameter)
}

func (b *Block) Restart(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Restart(block.Parameter)
}

func (b *Block) Scaling(block model.ConfigBlock) model.TemplateBlock {
	scaling := structBuilder.Scaling(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "policy" {
			policy := structBuilder.CustomBlock(b.Name, b.Parameter)
			scaling.Block = append(scaling.Block, policy)
		}
	}

	return scaling
}

func (b *Block) Service(block model.ConfigBlock) model.TemplateBlock {
	service := structBuilder.Service(block.Parameter)

	for _, b := range block.Block {
		switch b.Name {
		case "tagged_addresses", "meta", "canary_meta":
			block := structBuilder.CustomBlock(b.Name, b.Parameter)
			service.Block = append(service.Block, block)
		}
	}

	return service
}

func (b *Block) SidecarService(block model.ConfigBlock) model.TemplateBlock {
	sidecarService := structBuilder.SidecarService(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "meta" {
			meta := structBuilder.CustomBlock(b.Name, b.Parameter)
			sidecarService.Block = append(sidecarService.Block, meta)
		}
	}

	return sidecarService
}

func (b *Block) SidecarTask(block model.ConfigBlock) model.TemplateBlock {
	sidecarTask := structBuilder.SidecarTask(block.Parameter)

	for _, b := range block.Block {
		switch b.Name {
		case "config", "env", "meta":
			block := structBuilder.CustomBlock(b.Name, b.Parameter)
			sidecarTask.Block = append(sidecarTask.Block, block)
		}
	}

	return sidecarTask
}

func (b *Block) Spread(block model.ConfigBlock) model.TemplateBlock {
	spread := structBuilder.Spread(block.Parameter)

	for _, p := range block.Block {
		for _, item := range p.Parameter {
			for k, v := range item {
				if k == "target" {
					target := structBuilder.SpreadTarget(v.(map[string]interface{}))
					spread.Block = append(spread.Block, target)
				}
			}
		}
	}

	return spread
}

func (b *Block) Task(block model.ConfigBlock) model.TemplateBlock {
	task := structBuilder.Task(block.Parameter)

	for _, p := range block.Block {
		if p.Name == "config" {
			target := structBuilder.CustomBlock(p.Name, p.Parameter)
			task.Block = append(task.Block, target)
		}
	}

	return task
}

func (b *Block) Template(block model.ConfigBlock) model.TemplateBlock {
	template := structBuilder.Template(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "wait" {
			wait := structBuilder.CustomBlock(b.Name, b.Parameter)
			template.Block = append(template.Block, wait)
		}
	}

	return template
}

func (b *Block) Update(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Update(block.Parameter)
}

func (b *Block) Upstream(block model.ConfigBlock) model.TemplateBlock {
	upstream := structBuilder.Upstream(block.Parameter)

	for _, b := range block.Block {
		if b.Name == "mesh_gateway" {
			meshGateway := structBuilder.MeshGateway(b.Parameter)
			upstream.Block = append(upstream.Block, meshGateway)
		}

		if b.Name == "config" {
			config := structBuilder.CustomBlock(b.Name, b.Parameter)
			upstream.Block = append(upstream.Block, config)
		}
	}

	return upstream
}

func (b *Block) Vault(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.Vault(block.Parameter)
}

func (b *Block) Volume(block model.ConfigBlock) model.TemplateBlock {
	volume := structBuilder.Volume(block.Parameter)

	for _, p := range block.Block {
		for _, item := range p.Parameter {
			for k, v := range item {
				if k == "mount_options" {
					mountOptions := structBuilder.MountOptions(v.(map[string]interface{}))
					volume.Block = append(volume.Block, mountOptions)
				}
			}
		}
	}

	return volume
}

func (b *Block) VolumeMount(block model.ConfigBlock) model.TemplateBlock {
	return structBuilder.VolumeMount(block.Parameter)
}
