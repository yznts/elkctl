package elkd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// Options for elkd wrapper.
type Options struct {
	Path           string        // path to elkd binary
	TimeoutDefault time.Duration // default timeout for operations (will be overridden by specific timeouts)
}

// Elk acts as a wrapper around elkd tool.
// It spins up a process on a given ID/MAC address and provides methods to interact with it,
// like power controls, color/brightness setting, automatic retries, etc.
type Elk struct {
	// Address of the target device
	Address string
	// Options for elkd
	Options Options
	// elkd process bindings
	process *os.Process
	stdin   io.WriteCloser
	stdout  io.Reader
}

// Start the elkd process and ensure it's running (waiting for first OK interaction).
func (e *Elk) Start() error {
	// Stop any existing process
	if e.process != nil {
		e.process.Kill()
		e.process = nil
	}
	// Start new elkd process
	cmd := exec.Command(e.Options.Path, e.Address)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}
	// Run
	cmd.Start()
	// Update object
	e.process = cmd.Process
	e.stdin = stdin
	e.stdout = io.MultiReader(stdout, stderr)
	// Wait for OK response
	_, err = e.scan(5 * time.Second)
	if err != nil {
		return err
	}
	// All good
	return nil
}

// Exec sends a command to elkd and waits for a response with a timeout.
func (e *Elk) Exec(command string, timeout time.Duration) (string, error) {
	// Send command
	_, err := e.stdin.Write([]byte(command + "\n"))
	if err != nil {
		out, _ := io.ReadAll(e.stdout)
		return "", fmt.Errorf("stdin write error: %w %s", err, string(out))
	}
	// Read response
	return e.scan(timeout)
}

// scan reads a single line response from elkd with a timeout.
func (e *Elk) scan(timeout time.Duration) (string, error) {
	// Read response
	scanner := bufio.NewScanner(e.stdout)
	done := make(chan struct{})
	var line string
	var scanErr error
	go func() {
		if scanner.Scan() {
			line = scanner.Text()
		} else {
			scanErr = scanner.Err()
		}
		log.Println(line, scanErr)
		close(done)
	}()
	// Wait for response or timeout
	select {
	case <-done:
		if scanErr != nil {
			return "", fmt.Errorf("scan error: %w", scanErr)
		}
		return line, nil
	case <-time.After(timeout):
		return "", errors.New("command timeout")
	}
}

// New creates a new Elk instance with the given address and options.
func New(address string, options Options) *Elk {
	// Default options
	if options.Path == "" {
		options.Path = "elkd"
	}
	// Create instance
	return &Elk{
		Address: address,
		Options: options,
	}
}
