# promsarama

Use prometheus and sarama together.

You should use `promsarama`'s Registry and export prometheus by http.

```
config.MetricRegistry = promsarama.NewRegistry()

go func() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}()
```

More deatil in `example`.
