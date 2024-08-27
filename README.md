# SALUDOS EN GO
Este paquete proporciona saludos aleatorios

# Intalacion 
ejecuta el siguiente codigo para instalar

```bash
go get -u github.com/CristianLadino/greetings
```

### Uso 
Aqui tiene un ejemplo de como usar el codigo

```go
import (
	"fmt"
	"github.com/CristianLadino/greetings"
)

func main() {
	message, err := greetings.Hello("Caeld")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```

Este es un ejemplo de como se importa el paque y lo usa para saludar