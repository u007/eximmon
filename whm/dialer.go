package whm

import (
	"crypto/tls"
	"net"
	"time"
)

func WHMDialer() (net.Conn, error) {
	dial, err := (&net.Dialer{
		Timeout:   time.Minute * 5,
		KeepAlive: 30 * time.Second,
	}).Dial("tcp", ApiHost+":2087")

	if err != nil {
		return dial, err
	}

	dial.SetReadDeadline(time.Now().Add(time.Minute * 5))
	conn := tls.Client(dial, &tls.Config{InsecureSkipVerify: true})

	// 	Transport: &http.Transport{
	// 		DisableKeepAlives: true,
	// },

	return conn, nil
}

func CPanelDialer() (net.Conn, error) {
	dial, err := (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial("tcp", ApiHost+":2083")

	if err != nil {
		return dial, err
	}

	dial.SetReadDeadline(time.Now().Add(time.Minute * 5))
	conn := tls.Client(dial, &tls.Config{InsecureSkipVerify: true})

	return conn, nil
}
