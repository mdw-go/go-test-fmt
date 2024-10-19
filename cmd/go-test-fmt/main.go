package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdw-go/go-test-fmt/gotest"
)

var Version = "dev"

func main() {
	log.SetFlags(0)
	flags := flag.NewFlagSet(fmt.Sprintf("`%s` @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintln(flags.Output(),
			"Read go test ./... output from stdin and format it (kind of like piping to column -t, but better).")
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])
	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintln(&builder, line)
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	formatted := gotest.Format(builder.String())
	_, err = io.WriteString(os.Stdout, formatted)
	if err != nil {
		log.Fatal(err)
	}
}
