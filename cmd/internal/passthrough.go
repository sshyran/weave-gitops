package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	HelpCmd  = "help"
	HelpFlag = "--help"

	final    = "Final"
	initial  = "Initial"
	usage    = "Usage"
	examples = "Examples"
	flags    = "Flags"
)

type PreFun func(fs *pflag.FlagSet, args []string) error

type PostFun func(fs *pflag.FlagSet, args []string, wrappedOutput []byte) error

type Wrapper struct {
	Before PreFun
	After  PostFun
}

type BackingCLI struct {
	CmdName     string
	Alias       string
	VersionPath string
	Version     string
}

type WrappedCommand struct {
	Cmd             *cobra.Command
	AdditionalFlags []string
}

func (bc BackingCLI) validate() {
	versionOut, err := bc.Invoke(bc.VersionPath, true)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s version lookup failed", bc.CmdName)
		os.Exit(1)
	}

	if strings.HasPrefix(string(versionOut), bc.Version) {
		fmt.Fprintf(os.Stderr, "%s version check failed -- expected: %q, received: %q", bc.CmdName, bc.Version, versionOut)
		os.Exit(1)
	}
}

func (bc BackingCLI) Invoke(argStr string, silent bool) ([]byte, error) {
	path, err := exec.LookPath(bc.CmdName)
	if err != nil {
		return nil, err
	}

	adjustHelpOutput := false

	if strings.HasPrefix(argStr, HelpCmd) || strings.Index(argStr, " "+HelpFlag+" ") != -1 || strings.HasSuffix(argStr, " "+HelpFlag) {
		adjustHelpOutput = true
	}

	cmd := exec.Command("sh", "-c", path+" "+argStr)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	done := make(chan bool)
	defer close(done)

	process := func(rc io.ReadCloser, stream *os.File) error {
		defer func() { done <- true }()

		r := bufio.NewReader(rc)
		state := initial

		for {
			rawLine, _, err := r.ReadLine()
			if err != nil {
				return nil
			}

			line := string(rawLine)

			if adjustHelpOutput {
				headerIndex := strings.Index(line, ":")

				if headerIndex != -1 {
					switch line[0:headerIndex] {
					case usage:
						state = usage
					case examples:
						state = examples
					case flags:
						state = final
					}
				}

				trimmed := strings.TrimLeft(line, " ")

				if (state == usage || state == examples) && strings.HasPrefix(trimmed, bc.CmdName) {
					line = strings.Repeat(" ", len(line)-len(trimmed)) + bc.Alias + strings.TrimLeft(trimmed, bc.CmdName)
				} else if state == final {
					prefix := fmt.Sprintf(`Use "%s `, bc.CmdName)

					if strings.HasPrefix(line, prefix) {
						line = fmt.Sprintf("Use %s %s", bc.Alias, strings.TrimPrefix(line, prefix))
					}
				}
			}

			_, err = b.Write([]byte(line + "\n"))
			if err != nil {
				return err
			}

			if !silent {
				fmt.Fprintf(stream, "%s\n", line)
			}
		}

		return nil
	}

	var outErr error
	var errErr error

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		outErr = process(stdout, os.Stdout)
	}()

	go func() {
		errErr = process(stderr, os.Stderr)
	}()

	<-done
	<-done

	cmd.Wait()

	if outErr != nil {
		return nil, outErr
	}

	if errErr != nil {
		return nil, errErr
	}

	return b.Bytes(), nil
}

func passThrough(fs *pflag.FlagSet, newFlags []string) string {
	var sb strings.Builder

	fs.VisitAll(func(f *pflag.Flag) {
		for _, nf := range newFlags {
			if nf == f.Name {
				return
			}
		}

		sb.WriteString("--")
		sb.WriteString(f.Name)
		sb.WriteString(" '")
		sb.WriteString(f.Value.String())
		sb.WriteString("' ")
	})

	return sb.String()
}

func WrapCommand(cmd *cobra.Command, backingCLI BackingCLI, wrapper Wrapper, additionalFlags []string) error {
	backingCLI.validate()

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := wrapper.Before(cmd.Flags(), args); err != nil {
			return err
		}

		passThroughFlags := passThrough(cmd.Flags(), additionalFlags)

		out, err := backingCLI.Invoke(cmd.Name()+" "+passThroughFlags+strings.Join(args, " "), false)
		if err != nil {
			return err
		}

		if err := wrapper.After(cmd.Flags(), args, out); err != nil {
			return err
		}

		return nil
	}

	return nil
}
