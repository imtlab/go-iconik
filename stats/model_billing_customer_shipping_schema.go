/*
 * iconik_stats
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stats

type BillingCustomerShippingSchema struct {
	Address *BillingCustomerShippingAddressSchema `json:"address"`
	Name string `json:"name"`
	Phone string `json:"phone,omitempty"`
}
