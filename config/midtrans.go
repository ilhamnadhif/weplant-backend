package config

import "os"

func GetMidtransKey() string {
	return os.Getenv("MIDTRANS_SERVER_KEY")
}
