/*
 * iconik_stats
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stats

type BillingSettingsSchema struct {
	AutoRefillAmount int32 `json:"auto_refill_amount,omitempty"`
	EnableAutoTopUp bool `json:"enable_auto_top_up,omitempty"`
	LowBalanceTrigger int32 `json:"low_balance_trigger,omitempty"`
}
