package logging

import (
	"os"
	"testing"

	"github.com/Kaspetti/Kaspebiten/internal/logging"
)


func createTestEnvironment() {
    os.Mkdir("./testenv", 700)
    os.Chdir("./testenv")
}


func deleteTestEnvironment() {
    os.Chdir("..")
    os.RemoveAll("./testenv")
}


func TestInitialize(t *testing.T) {
    createTestEnvironment()
    defer deleteTestEnvironment()

    logging.InitializeLogger()
    if _, err := os.Stat("./testenv/log"); os.IsNotExist(err) {
        t.Fatalf("")
    }
}


func TestLogType(t *testing.T) {
}
