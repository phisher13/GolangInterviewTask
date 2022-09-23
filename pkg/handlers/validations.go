package handlers

import "net/url"

func IsValidUrl(url_addr string) bool {
	_, err := url.ParseRequestURI(url_addr)
	
	return err == nil
}