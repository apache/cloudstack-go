package main

import (
	"encoding/json"
	"fmt"
	"github.com/apache/cloudstack-go/v2/cloudstack"
	"log"
)

func main()  {
	cs := cloudstack.NewAsyncClient("http://localhost:8080/client/api",
		"zNeTB-XgjvqRLrlf0t_dfVc1kLPBOv88wyUKaUcD0ycCJOkkxe3iRoZcv4JGxaU1882dGuCnCkUcDsirK3M8nw",
		"JmdgVR-TXsozU1sRBYaBs1ctCWUEnvoxKYwJw_xZW4TRMFu4SU3J9tP8hEw5od26iU9HWX0w_oUWv07EcXnItw", false)
	p := cs.Host.NewAddHostParams("Simulator", "password", "5382edc2-e689-4074-bd67-0e1a236eb2bc",
		"http://sim/c0/h0", "root", "d4a81f75-5d92-415e-ab59-e85cc2ce56d9")
	resp, err := cs.Host.AddHost(p)
	if err != nil {
		fmt.Errorf("Failed to add host due to: %v", err)
	}

	b, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("Host response : %v", string(b))
}