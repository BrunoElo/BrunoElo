package exporter

import (
		"fmt"
		"os"
		"go_export/smediafeeds"
)

// Export function exports the feeds to files
func Export(u smediafeeds.SocialMedia, filename string) error {
	f,err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY,0755)
	if err != nil {
		return err
	}
	for _, fd := range u.Feed(){
		n, err := f.Write([]byte(fd+"\n"))
		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}