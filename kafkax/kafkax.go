package kafkax

type Kafka struct {
	BootstrapServer string   `json:"boostrapserver" mapstructure:"boostrapserver"`
	Username        string   `json:"username" mapstructure:"username"`
	Password        string   `json:"password" mapstructure:"password"`
	GroupId         string   `json:"groupid" mapstructure:"groupid"`
	ListenOnTopics  []string `json:"listenontopics" mapstructure:"listenontopics"`
	ProducToTopic   string   `json:"productotopic" mapstructure:"productotopic"`
	SchemaRegistry  *SchemaRegistry
}

type SchemaRegistry struct {
	Url          string `json:"url" mapstructure:"url"`
	TopicPattern string `json:"topicpattern" mapstructure:"topicpattern"`
}
