package logging

type StdOutLogger struct {
    // like rect, the data structure
    Val int    
}

type Chef_Logger interface {
    // like inheritance, any object that does all the methods here can be called with <obj>.Log()
    Log() string
}

func (StdOutLogger) Log() string {
    return "Hello, world."
}

/*
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
*/