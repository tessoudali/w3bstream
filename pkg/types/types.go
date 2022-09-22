package types

type UploadConfig struct {
	Root          string `env:""`
	FileSizeLimit int64  `env:""`
}

func (c *UploadConfig) SetDefault() {
	if c.Root == "" {
		c.Root = "./asserts"
	}
	if c.FileSizeLimit == 0 {
		c.FileSizeLimit = 100 * 1024 * 1024
	}
}

type EventChanConfig struct {
	Limit int `env:""`
}

func (v *EventChanConfig) SetDefault() {
	if v.Limit == 0 {
		v.Limit = 500
	}
}
