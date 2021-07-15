package main

import (
	"fyne.io/fyne/v2"
)

type appData struct {
	servers []*server
}

type server struct {
	id            int
	name, iconURL string
	channels      []*channel
	service       service
}

func (s *server) icon() fyne.Resource {
	// TODO cache this resource
	icon, err := fyne.LoadResourceFromURLString(s.iconURL)
	if err != nil {
		fyne.LogError("Failed to read icon "+s.iconURL, err)
		return nil
	}

	return icon
}

type channel struct {
	id       int
	name     string
	messages []*message
}

type message struct {
	content, author string
	avatar          string
}

func findChan(d *appData, sID, cID int) *channel {
	for _, s := range d.servers {
		if s.id == sID {
			for _, c := range s.channels {
				if c.id == cID {
					return c
				}
			}
		}
	}
	return nil
}
