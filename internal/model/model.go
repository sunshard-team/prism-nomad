package model

type Update struct {
	MaxParallel      int    `yaml:"max_parallel"`
	MinHealthyTime   string `yaml:"min_healthy_time"`
	HealthyDeadline  string `yaml:"healthy_deadline"`
	ProgressDeadline string `yaml:"progress_deadline"`
	AutoRevert       bool   `yaml:"auto_revert"`
	Canary           int    `yaml:"canary"`
}

type Migrate struct {
	MaxParallel     int    `yaml:"max_parallel"`
	HealthCheck     string `yaml:"health_check"`
	MinHealthyTime  string `yaml:"min_healthy_time"`
	HealthyDeadline string `yaml:"healthy_deadline"`
}

type Check struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Interval string `yaml:"interval"`
	Timeout  string `yaml:"timeout"`
}

type Service struct {
	Provider string `yaml:"provider"`
	Check    Check  `yaml:"check"`
}

type Restart struct {
	Attempts int    `yaml:"attempts"`
	Interval string `yaml:"interval"`
	Delay    string `yaml:"delay"`
	Mode     string `yaml:"mode"`
}

type EphemeralDisk struct {
	Sticky  bool `yaml:"sticky"`
	Migrate bool `yaml:"migrate"`
	Size    int  `yaml:"size"`
}

type Task struct {
	Args         []string `yaml:"args"`
	Bind         string   `yaml:"bind"`
	Datacenter   string   `yaml:"datacenter"`
	MyEnv        string   `yaml:"my_env"`
	MetaFoo      string   `yaml:"meta_foo"`
	Command      string   `yaml:"command"`
	Ports        string   `yaml:"ports"`
	Volumes      []string `yaml:"volumes"`
	AuthSoftFail bool     `yaml:"auth_soft_fail"`
}

type Logs struct {
	MaxFiles    int `yaml:"max_files"`
	MaxFileSize int `yaml:"max_file_size"`
}

type Templates struct {
	Data         string            `yaml:"data"`
	Destination  string            `yaml:"destination"`
	ChangeMode   string            `yaml:"change_mode"`
	Perms        string            `yaml:"perms"`
	EmbeddedTmpl string            `yaml:"embedded_tmpl"`
	Envsubst     bool              `yaml:"envsubst"`
	Env          map[string]string `yaml:"env"`
}

type Config struct {
	InstancesCount  int           `yaml:"instances_count"`
	Image           string        `yaml:"image"`
	EnvironmentVar1 string        `yaml:"environment_var1"`
	EnvironmentVar2 string        `yaml:"environment_var2"`
	CPU             int           `yaml:"cpu"`
	Memory          int           `yaml:"memory"`
	Update          Update        `yaml:"update"`
	Migrate         Migrate       `yaml:"migrate"`
	Service         Service       `yaml:"service"`
	Restart         Restart       `yaml:"restart"`
	EphemeralDisk   EphemeralDisk `yaml:"ephemeral_disk"`
	Task            Task          `yaml:"task"`
	Logs            Logs          `yaml:"logs"`
	Templates       []Templates   `yaml:"templates"`
}
