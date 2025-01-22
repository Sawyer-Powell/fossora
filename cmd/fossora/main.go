package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"go.etcd.io/etcd/server/v3/embed"
	//"context"
	//"log"
	//"tailscale.com/client/tailscale"
)

func main() {

	cfg := embed.NewConfig()
	cfg.Dir = "/tmp/fossora/default.etcd"
	
	machine_name := os.Args[1]

	metta_ip := "100.91.155.67"
	pi0_ip := "100.90.0.11"
	pi1_ip := "100.123.57.97"

	var peer_ip string 

	switch machine_name {
	case "metta":
		peer_ip = metta_ip
	case "pi0":
		peer_ip = pi0_ip
	case "pi1":
		peer_ip = pi1_ip
	}

	listen_client_local, _ := url.Parse("http://127.0.0.1:3445")
	listen_client, _ := url.Parse("http://" + peer_ip + ":3445")

	peer_url, _ := url.Parse("http://" + peer_ip + ":3444")

	cfg.ListenPeerUrls = []url.URL{*peer_url}
	cfg.ListenClientUrls = []url.URL{*listen_client, *listen_client_local}
	cfg.AdvertisePeerUrls = []url.URL{*peer_url}
	cfg.AdvertiseClientUrls = []url.URL{*listen_client}
	cfg.InitialCluster = "metta=http://100.91.155.67:3444,pi0=http://100.90.0.11:3444,pi1=http://100.123.57.97:3444"
	cfg.InitialClusterToken = "fossora"
	cfg.Name = os.Args[1]

	e, err := embed.StartEtcd(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer e.Close()

	select {
	case <-e.Server.ReadyNotify():
		log.Printf("Server is ready!")
	case <-time.After(60 * time.Second):
		e.Server.Stop() // trigger shutdown
		log.Printf("Server took too long to start")
	}

	log.Fatal(<-e.Err())

	//tailscale.I_Acknowledge_This_API_Is_Unstable = true;

	//client := tailscale.NewClient("sawyerhpowell@gmail.com", tailscale.APIKey("tskey-api-kKjpVobM9221CNTRL-2TWCQs3wgH61djgKV9ANH6FWbxEPZH2f"))
	//ctx := context.Background()

	//devices, err := client.Devices(ctx, nil)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, device := range(devices) {
	//	log.Print(device.Addresses[0])
	//}


	//cfg := embed.NewConfig()
	//cfg.Dir = "/tmp/fossora/default.etcd"
	//peerUrl, _ := url.Parse("http://localhost:3000")
	//clientUrl, _ := url.Parse("http://localhost:3001")

	//cfg.ListenPeerUrls = []url.URL{*peerUrl}
	//cfg.ListenClientUrls = []url.URL{*clientUrl}

	//e, err := embed.StartEtcd(cfg)

	//if err != nil {
	//	log.Fatal(err)
	//}

	//defer e.Close()

	//select {
	//case <-e.Server.ReadyNotify():
	//	log.Printf("Server is ready!")
	//case <-time.After(60 * time.Second):
	//	e.Server.Stop() // trigger shutdown
	//	log.Printf("Server took too long to start")
	//}

	//log.Fatal(<-e.Err())
}
