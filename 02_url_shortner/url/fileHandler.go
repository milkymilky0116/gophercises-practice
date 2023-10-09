package urlshortner

import "os"

func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}
