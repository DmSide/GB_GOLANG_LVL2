package _defer

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"time"
)

type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

type ErrorWithTrace struct {
	text  string
	trace string
	time  time.Time
}

func NewErrorWithTrace(text string) error {
	return &ErrorWithTrace{
		text:  text,
		trace: string(debug.Stack()),
		time:  time.Now().UTC(),
	}
}

func (e *ErrorWithTrace) Error() string {
	return fmt.Sprintf("error: %s\ntrace:\n%stime:%s\n", e.text, e.trace, e.time)
}

func lessonExamples() {
	var err error
	myErr := NewErrorWithTrace("my error")
	// fmt.Println(err)

	err = fmt.Errorf("found error: %w", myErr)
	// fmt.Println(err)
	err2 := fmt.Errorf("found error: %w", myErr)

	fmt.Println(errors.Is(err, myErr)) // true
	fmt.Println(errors.Is(err, &ErrorWithTrace{}))
	fmt.Println(errors.Is(myErr, &ErrorWithTrace{}))
	fmt.Println(errors.Is(err, err2))

	fmt.Println(errors.As(err, &myErr)) // true
}

func SafeFileWriter() {
	var err error

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString("test")
}

func SafeDivision() error {
	var err error

	defer func() {
		if v := recover(); v != nil {
			err = NewErrorWithTrace("division by zero")
		}
	}()

	var a int
	fmt.Println(1 / a)

	return err
}
