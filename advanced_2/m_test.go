package main_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"reflect"
	"runtime/debug"
	"sync"
	"testing"
)

func Test_fmt(t *testing.T) {
	// Bytes from a png image or whatever
	pngPayload := []byte{137, 80, 78, 71, 13, 10, 26, 10, 11, 12, 14}
	buf := make([]byte, 4)

	// read the first 4 bytes into buffer - Essentially the png header
	_, err := io.ReadFull(bytes.NewReader(pngPayload), buf)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(buf)

	// alternative way to write to stdout
	io.WriteString(os.Stdout, string(buf))
	// alternative implementation
	lr := io.LimitReader(bytes.NewReader(pngPayload), 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
	tempFile, err := os.CreateTemp(".", "temp_")
	if err != nil {
		t.Fatal(err)
	}

	// defer os.Remove(tmpFile.Name) clear up
	defer func(tmpFile *os.File) {
		// close can fail
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
	}(tempFile)

	// write pid to file
	if _, err := tempFile.Write([]byte(fmt.Sprintf("%d", os.Getpid()))); err != nil {
		log.Fatal(err)
	}

	//alternatively
	os.WriteFile("pid.pid", []byte(fmt.Sprintf("%d", os.Getpid())), 0664)
}

type LowerCaseReader struct {
	text string
}

func NewLowerCaseReader(text string) *LowerCaseReader {
	return &LowerCaseReader{
		text: text,
	}
}

func (r *LowerCaseReader) Read(p []byte) (int, error) {
	buf := make([]byte, len(r.text))

	for i := 0; i < len(buf); i++ {
		buf[i] = r.text[i] | 0x20
	}
	n := copy(p, buf)

	return n, io.EOF
}
func TestReader(t *testing.T) {
	// example of implmenting a custom io.Reader
	r := NewLowerCaseReader("ALL CAPITALS")
	resp, err := io.ReadAll(r)
	if err != nil {
		t.Fatal("error: something was wrong when converting to lowecase")
	}

	fmt.Println(string(resp))
}

func TestLog(t *testing.T) {
	log.Println("Log Enter")
	log.SetFlags(0)
	for i := 0; i < 100; i++ {
		go log.Println(i)
	}

	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	log.SetOutput(io.Discard)
	log.Println("ENtry 2")
	defer log.Println("Will not be logged")
	//log.Fatal("Exit")
}

func TestOS(t *testing.T) {
	// command-line arguments
	fmt.Println(os.Args)

	// file existence check
	_, err := os.Stat("i-dont-exist")
	fmt.Printf("%v\n", os.IsNotExist(err))

	// print current directory
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current dir is %s\n", currDir)

	// Get Environment Variable
	fmt.Printf("Current $PATH is %s\n", os.Getenv("PATH"))

	// Get current username
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current User name is %s", currUser.Username)

	// Note: Platform dependant command
	cmd := exec.Command("ls", "-ltr")
	// Convenient wrapper
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	// Will not get called
	defer fmt.Println("Call Exit")

	//os.Exit(0)
}

var SomeError = errors.New("error:description")
var EOFError = errors.New("EOF")

// implements error interface
type CommandError struct {
	err string
}

func (e CommandError) Error() string {
	return e.err
}

func Avoid() error {
	return &CommandError{
		err: "It's an error",
	}
}

func TestError(t *testing.T) {
	if EOFError == io.EOF {
		log.Fatal("should not happen")
	}

	// prints *errors.errorString so its not ideal for switch statements
	fmt.Println(reflect.TypeOf(SomeError))

	switch SomeError.(type) {
	case error:
		fmt.Println("it's is error")
	}

	switch SomeError {
	case SomeError:
		fmt.Println("it's is same error")
	}

	// type hint is required for switch
	var invalidCommand error = CommandError{"Invalid Command"}
	switch invalidCommand.(type) {
	case CommandError:
		fmt.Println(invalidCommand)
	}

	err := Avoid()
	// runtime time check
	if err, ok := err.(*CommandError); ok {
		log.Fatal(err)
	}
}

func callbackSafe(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	panic("Panic!!")
}

func callbackUnsafe(m *sync.Mutex) {
	m.Lock()
	panic("Panic!!")
	// will not unlock
	m.Unlock()
}

func TestPanic(t *testing.T) {
	// not all defers will work. will not compile
	// defer append([]string{"1", "2", "3"}, "4")

	// be careful when assigning and comparing interfaces
	var v interface{} = nil
	var arr []int = nil
	v = arr

	// will panic as []int is an uncomparable type
	fmt.Println(v == v)
	defer fmt.Println("will print string.")
	panic("Panic")
	defer fmt.Println("will not print string.")

	var m = &sync.Mutex{}

	defer func() {
		if v := recover(); v != nil {
			log.Println("Revovered from:", v)
			log.Printf("Lock is: %v\n", m)
		}
	}()

	callbackSafe(m)
	callbackUnsafe(m)

	// Panic suppression. FIFO applies here
	defer panic(1)
	defer panic(2)
	panic(3)
}

func TestStackTrace(t *testing.T) {
	// main go routine
	debug.PrintStack()
	defer debug.PrintStack()
	done := make(chan bool)

	//from go routine
	go func(done chan bool) {
		debug.PrintStack()
		done <- true
		close(done)
	}(done)

	<-done

	// use stack()
	stack_trace := debug.Stack()
	fmt.Printf("%s", string(stack_trace))
}
