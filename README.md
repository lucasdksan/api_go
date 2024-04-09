# Primeira API usando GO puro

```go
func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	string_64_base := base64.StdEncoding.EncodeToString(key)

	fmt.Print(string_64_base)
}
```