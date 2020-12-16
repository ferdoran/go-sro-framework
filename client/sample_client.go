package client

func StartSampleClient() {
	client := NewClient([]byte{127, 0, 0, 1}, 15779, "SR_Client")
	client.Connect()
}
