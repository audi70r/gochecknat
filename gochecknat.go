package gochecknat

import (
	"fmt"
	"github.com/pion/webrtc/v3"
	"sync"
)

type NATInfo struct {
	Symmetric  bool
	IP         string
	Port       uint16
	Candidates []webrtc.ICECandidate
}

func GetNATInfo() (info NATInfo, err error) {
	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun1.l.google.com:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if cErr := peerConnection.Close(); cErr != nil {
			fmt.Printf("cannot close peerConnection: %v\n", cErr)
		}
	}()

	_, err = peerConnection.CreateDataChannel("nat-checker", &webrtc.DataChannelInit{})

	if err != nil {
		return info, err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate == nil {
			wg.Done()
			return
		}

		if candidate.Typ != webrtc.ICECandidateTypeSrflx {
			return
		}

		if candidate.Port != candidate.RelatedPort {
			info.Symmetric = true
		}

		info.IP = candidate.Address
		info.Port = candidate.Port
		info.Candidates = append(info.Candidates, *candidate)
	})

	// Create an offer to send to the other process
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}

	// Sets the LocalDescription, and starts our UDP listeners
	// Note: this will start the gathering of ICE candidates
	if err = peerConnection.SetLocalDescription(offer); err != nil {
		panic(err)
	}

	wg.Wait()

	return info, err
}
