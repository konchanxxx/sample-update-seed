package request

// PostPubSubParams はCloud Pub/Subからのリクエストボディ
type PostPubSubParams struct {
	Message      *Message `json:"message"`
	Subscription string   `json:"subscription"`
}

// Message はCloud Pub/Subからのリクエストボディに含まれるパラメータ
type Message struct {
	Data        string      `json:"data"`
	Attributes  *Attributes `json:"attributes"`
	MessageID   string      `json:"messageId"`
	PublishTime string      `json:"publishTime"`
}

// Attributes はCloud Pub/Subからのリクエストボディに含まれるパラメータ
type Attributes struct {
	ObjectGeneration   string `json:"objectGeneration"`
	BucketID           string `json:"bucketId"`
	EventType          string `json:"eventType"`
	NotificationConfig string `json:"notificationConfig"`
	PayloadFormat      string `json:"payloadFormat"`
	ObjectID           string `json:"objectId"`
}

// GetFileName はGCSで更新があったファイル名を返す
func (p *PostPubSubParams) GetFileName() string {
	fileName := p.Message.Attributes.ObjectID

	return fileName
}
