package qrzlog

import "context"

const agent = "xylo-go-1.0"

func Fetch(key *string) (*Response, error) {
	config := NewConfiguration()
	config.UserAgent = agent
	client := NewAPIClient(config)

	opts := RootPostOpts{}
	sessResp, _, err := client.DefaultApi.RootPost(context.TODO(), *key, "FETCH", &opts)
	if err != nil {
		return nil, err
	}
	return &sessResp, err
}
