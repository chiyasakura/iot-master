package mqtt

import (
	"encoding/json"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v3/internal"
	"github.com/zgwit/iot-master/v3/model"
)

func mergeProperties(id string, properties []model.ValuePayload) {
	dev := internal.Devices.Load(id)
	if dev == nil {
		dev = internal.NewDevice(id)
		internal.Devices.Store(id, dev)
	}
	//合并数据
	for _, p := range properties {
		dev.Properties[p.Name] = p.Value
	}
}

func subscribeProperty() error {
	Client.Subscribe("up/property/+/+", 0, func(client paho.Client, message paho.Message) {
		var payload model.UpPropertyPayload
		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			return
		}

		//属性值
		if payload.Id != "" && payload.Properties != nil {
			mergeProperties(payload.Id, payload.Properties)
		}

		//子设备属性
		if payload.Devices != nil {
			for _, d := range payload.Devices {
				mergeProperties(d.Id, d.Properties)
			}
		}

	})

	return nil
}
