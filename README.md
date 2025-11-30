# zaal-utils

`zaal-utils` یک پکیج Go برای اعتبارسنجی شناسه‌های ایرانی است که شامل ابزارهای زیر می‌باشد:

- **National ID Validator**: اعتبارسنجی کد ملی
- **IBAN Validator**: اعتبارسنجی شماره شبا

این پکیج می‌تواند به راحتی در پروژه‌های Go شما استفاده شود.

---

## نصب

```bash
go get github.com/sezarsaman/zaal-utils/validator
```

## استفاده

### اعتبارسنجی کد ملی

```go
package main

import (
	"fmt"
	"github.com/sezarsaman/zaal-utils/validator"
)

func main() {
	err := validator.ValidateNationalID("0013520839")
	if err != nil {
		fmt.Println("کد ملی نامعتبر است:", err)
	} else {
		fmt.Println("کد ملی معتبر است")
	}
}
```

### اعتبارسنجی IBAN

```go
package main

import (
	"fmt"
	"github.com/sezarsaman/zaal-utils/validator"
)

func main() {
	err := validator.ValidateIBAN("IR820540102680020817909002")
	if err != nil {
		fmt.Println("شماره شبا نامعتبر است:", err)
	} else {
		fmt.Println("شماره شبا معتبر است")
	}
}
```

## تست

```bash
go test ./validator
```

## لایسنس

MIT License