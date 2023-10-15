# Device

```go
type Device struct {
	IP      string  `json:"ip"`
	Company Company `json:"company"`
	Branch  Branch  `json:"branch"`
	User    User    `json:"user"`
	Type    string  `json:"type"`
}
```
