package demos

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"time"

	myutils "tools.app/utils"
)

// demo, bits operation
func testBitsOperation() {
	fmt.Println("左移 1<<2:", 1<<2)
	fmt.Println("右移 10>>2:", 10>>2)
	// 异或：两个二进位，不同则该位为1, 否则该位为0
	fmt.Println("异或 10^2:", 10^2)
	fmt.Println("异或 10^10:", 10^10)
	// 或：两个二进位中只要有一个为1, 该位的结果值为1
	fmt.Println("或 10|2:", 10|2)
	// 与：两个二进位都为1, 该位的结果值才为1, 否则为0
	fmt.Println("与 10&2:", 10&2)
}

// demo, base64 encode for bytes
func testBase64Encode() {
	size := 16
	b := make([]byte, size, size)
	for i := 0; i < size; i++ {
		b[i] = uint8(60 + i)
	}
	fmt.Println("\nbytes:", string(b))
	fmt.Println("base64 encode string:", myutils.GetBase64Text(b))
}

// demo, init slice by make, and recover func
func testInitSliceAndRecovery() {
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(runtime.Error); ok {
				fmt.Println("\n[runtime error]", err)
			} else {
				fmt.Println("\n[error]", p.(error))
			}
		}
	}()

	b := make([]byte, 0, 10)
	// b := make([]byte, 10, 10)
	for i := 0; i < 10; i++ {
		b[i] = uint8(60 + i)
	}
	fmt.Println("\nbytes:", string(b))
}

// demo, interface and type assert
type mockError struct {
	message string
}

func (e mockError) Error() string {
	return "mock " + e.message
}

func printError01(err interface{}) {
	if e, ok := err.(mockError); ok { // type assert
		fmt.Println("Mock Error:", e.Error())
		return
	}
	fmt.Println("not an error!")
}

func printError02(err interface{}) {
	if e, ok := err.(interface{ Error() string }); ok { // interface assert
		fmt.Println("Error:", e.Error())
		return
	}
	fmt.Println("not an error!")
}

func testInterfaceTypeAssert() {
	mockErr := mockError{
		message: "file write error!",
	}
	fmt.Println("\ntype assert:")
	printError01(mockErr)
	printError01("type")

	fmt.Println("\ninterface assert:")
	printError02(mockErr)
	printError02("interface")
}

// demo, point type assert
func isPointer(object interface{}) bool {
	t := reflect.TypeOf(object)
	fmt.Println("\nobject kind:", t.Kind())
	return t.Kind() == reflect.Ptr
}

func testPointTypeAssert() {
	mockErr := mockError{
		message: "file write error!",
	}
	fmt.Println("is point:", isPointer(mockErr))
	fmt.Println("is point:", isPointer(&mockErr))
}

// demo, event bus by channel
type dataEvent struct {
	Data  interface{}
	Topic string
}

type dataChannel chan dataEvent
type dataChannelSlice []dataChannel

type eventBus struct {
	subscribers map[string]dataChannelSlice
	rm          sync.RWMutex
}

func (eb *eventBus) Subscribe(topic string, ch dataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()

	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = []dataChannel{ch}
	}
}

func (eb *eventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()

	if chans, found := eb.subscribers[topic]; found {
		channels := append(dataChannelSlice{}, chans...)
		go func(data dataEvent, channels dataChannelSlice) {
			for _, ch := range channels {
				ch <- data
			}
		}(dataEvent{Data: data, Topic: topic}, channels)
	}
}

func publisTo(eb *eventBus, topic string, data string) {
	for i := 0; i < 10; i++ {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func printDataEvent(ch string, data dataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func testEventBus() {
	eb := &eventBus{
		subscribers: map[string]dataChannelSlice{},
	}

	ch1 := make(chan dataEvent)
	ch2 := make(chan dataEvent)
	ch3 := make(chan dataEvent)
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)

	go publisTo(eb, "topic1", "Hi topic 1")
	go publisTo(eb, "topic2", "Welcome to topic 2")

	timeout := time.After(time.Duration(5) * time.Second)
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		case <-timeout:
			fmt.Println("timeout, and exit.")
			return
		}
	}
}

// demo, 信号量
func testSemaphore() {
	ch := make(chan struct{}, 3)
	for i := 0; i < 20; i++ {
		go func(i int) {
			ch <- struct{}{}
			defer func() {
				<-ch
			}()
			fmt.Printf("routine [%d] is running ...\n", i)
			time.Sleep(time.Duration(1) * time.Second)
		}(i)
	}

	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		size := len(ch)
		fmt.Println("channel size:", size)
		if size == 0 {
			break
		}
	}
	close(ch)
	fmt.Println("test semaphore done.")
}

// demo, sync.Once
func testSyncOnce() {
	var once sync.Once

	onceBody := func() {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("run only once.")
	}

	done := make(chan bool)
	for i := 1; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 1; i < 10; i++ {
		<-done
	}
	fmt.Println("sync once test done.")
}

// demo, reflect set value
func testReflectSetValue01() {
	var i = int(1)

	// 需要传递变量的地址
	value := reflect.ValueOf(&i)
	// value是一个指针, 获取该指针指向的值, 相当于value.Elem()
	value = reflect.Indirect(value)
	// value = value.Elem()
	fmt.Println("value:", value.Interface())

	if value.Kind() == reflect.Int {
		fmt.Println("int value:", value.Int())
		value.SetInt(2)
	}
	fmt.Println("value:", value.Interface())
}

type testStruct struct {
	// 只有大写开头的成员变量可以Set
	Str string `json:"tag_str"`
}

// demo, reflect set value (struct)
func testReflectSetValue02() {
	s := testStruct{Str: "init"}

	value := reflect.ValueOf(&s)
	value = reflect.Indirect(value)
	fmt.Println("struct value:", value.Interface())

	f := value.FieldByName("Str")
	fmt.Println("Str field:", f.Interface())

	if f.Kind() == reflect.String && f.CanSet() {
		fmt.Println("String field:", f.String())
		f.SetString("updated")
		fmt.Println("String field:", f.String())
	}
	fmt.Println("struct:", s)
	fmt.Println("struct value:", value.Interface())
}

// demo, reflect get struct fields tag
func testReflectGetTags() {
	s := new(testStruct)

	t := reflect.TypeOf(s)
	fmt.Println("type:", t)

	value := reflect.ValueOf(s)
	value = reflect.Indirect(value)
	fmt.Println("value:", value)

	valueType := value.Type()
	fmt.Println("value type:", valueType)

	if f, ok := valueType.FieldByName("Str"); ok {
		fmt.Println("tag:", f.Tag.Get("json"))
	}
}

// MainDemo06 main for golang demo06.
func MainDemo06() {
	// testBitsOperation()
	// testBase64Encode()
	// testInitSliceAndRecovery()
	// testInterfaceTypeAssert()
	// testPointTypeAssert()
	// testEventBus()

	// testSemaphore()
	// testSyncOnce()

	// testReflectSetValue01()
	// testReflectSetValue02()
	// testReflectGetTags()

	fmt.Println("golang demo06 DONE.")
}
