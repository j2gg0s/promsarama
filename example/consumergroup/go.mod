module github.com/Shopify/sarama/examples/consumer

go 1.15

require (
	github.com/Shopify/sarama v1.27.0
	github.com/j2gg0s/promsarama v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.9.0
)

replace github.com/j2gg0s/promsarama => ../..
