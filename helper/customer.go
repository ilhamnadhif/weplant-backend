package helper

import "strings"

func GetCustomerIdFromOrderId(orderId string) string {
	return strings.Split(orderId, "-")[0]
}
