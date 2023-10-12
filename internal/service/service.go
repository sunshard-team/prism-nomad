package service

import (
	"prism/internal/model"
	"prism/internal/service/builder"
	"prism/internal/service/deployment"
	"prism/internal/service/output"
	"prism/internal/service/parser"
	"prism/internal/service/project"

	"github.com/hashicorp/nomad/api"
)

type Project interface {
	// Creates a project.
	Create(name string) (string, error)
}

type Parser interface {
	// Parsing the YAML configuration file.
	ParseYAML(file []byte) (map[string]interface{}, error)

	// Parsing the configuration map. Assembles a block structure.
	ParseConfig(blockType string, config map[string]interface{}) model.ConfigBlock
}

type BlockBuilder interface {
	CustomBlock(block model.ConfigBlock) model.TemplateBlock

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
	Network(block model.ConfigBlock) model.TemplateBlock
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
	Upstreams(block model.ConfigBlock) model.TemplateBlock
	UpstreamMeshGateway(block model.ConfigBlock) model.TemplateBlock
	Vault(block model.ConfigBlock) model.TemplateBlock
	Volume(block model.ConfigBlock) model.TemplateBlock
	VolumeMountOptions(block model.ConfigBlock) model.TemplateBlock
	VolumeMount(block model.ConfigBlock) model.TemplateBlock
}

type StructureBuilder interface {
	// Builds and returns a job configuration structure.
	BuildConfigStructure(buildStructure model.BuildStructure) model.TemplateBlock
}

type Deployment interface {
	// Returns the configuration structure.
	CreateConfigStructure(
		parameter model.ConfigParameter,
	) (model.TemplateBlock, error)

	// Checking if the namespace exists in the nomad cluster.
	CheckNamespace(namespace model.CheckNamespace) error

	// Job configuration deployment in the nomad cluster.
	Deployment(client *api.Client, config string) (string, error)
}

type Changes interface {
	SetChanges(config *model.TemplateBlock, changes *model.Changes) error
}

type Output interface {
	// Returns the formated job configuration of the nomad.
	OutputConfig(config model.TemplateBlock) (string, error)

	// Creates a nomad configuration file in .nomad.hcl format.
	CreateConfigFile(name, path string, config model.TemplateBlock) error
}

type Service struct {
	Project          Project
	Output           Output
	Parser           Parser
	BlockBuilder     BlockBuilder
	StructureBuilder StructureBuilder
	Changes          Changes
	Deployment       Deployment
}

func NewService(
	p *parser.Parser,
	bb *builder.BlockBuilder,
	sb *builder.StructureBuilder,
	c *builder.Changes,
	o *output.Output,
) *Service {
	return &Service{
		Project:          project.NewProject(),
		Output:           output.NewOutput(),
		Parser:           parser.NewParser(),
		BlockBuilder:     builder.NewBlockBuilder(),
		StructureBuilder: builder.NewStructureBuilder(*bb),
		Changes:          builder.NewChanges(),
		Deployment:       deployment.NewDeployment(*p, *sb, *c),
	}
}
