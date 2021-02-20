package "github.com/sirupsen/logrus"

const {
	ServiceName = "telegram"
	dbTimeout=10
}

func main() {
	fmter := formatters.NewGelf(ServiceName)
}