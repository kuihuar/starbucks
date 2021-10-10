package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"starbucks/suanfa/search"
	"testing"
)

////Abstract Interface
//type iAlpha interface {
//	work()
//	common()
//}
//
////Abstract Concrete Type
//type alpha struct {
//	name string
//	work func()
//}
//
//func (a *alpha) common() {
//	fmt.Println("common called")
//	a.work()
//}
//
////Implementing Type
//type beta struct {
//	alpha
//}
//
//func (b *beta) work() {
//	fmt.Println("work called")
//	fmt.Printf("Name is %s\n", b.name)
//}
//func check(i iAlpha)  {
//	i.common()
//}
func TestCase1(t *testing.T) {
	name := "Bob"
	age := 10

	assert.Equal(t, "bob", name)
	assert.Equal(t, 20, age)
}
func aaa(){
	//for{
	//	if left > right {
	//		break
	//	}
	//	for i:=left; i<right;i++{
	//		//fmt.Printf("%d-%d\t",top,i)
	//		res = append(res,fmt.Sprintf("%d-%d",top,i))
	//	}
	//	top++
	//	for i:=top; i<button;i++{
	//		//fmt.Printf("%d-%d\t",i,right)
	//		res = append(res,fmt.Sprintf("%d-%d",i,right))
	//	}
	//	right--
	//	for i:= right;i>left;i--{
	//		//fmt.Printf("%d-%d\t",button,i)
	//		res = append(res,fmt.Sprintf("%d-%d",button,i))
	//	}
	//	button--
	//	for i:= button;i>top;i--{
	//		//fmt.Printf("%d-%d\t",i,left)
	//		res = append(res,fmt.Sprintf("%d-%d",i,left))
	//	}
	//	left++
	//}
}
func main() {
	//res := array_style.ThreeSum4([]int{-1,0,1,2,-1,4})
	//sort.TestBubbleSort()
	search.TestSPBM()
	fmt.Println("----------")
	//m  := new(map[string]int)
	//
	//*m = map[string]int{}
	//(*m)["a"] = 5
	//fmt.Printf("%+v\n",m)
	//fmt.Println("end")
	//
	//s  := new([]int)
	//*s = append(*s, 0)
	//(*s)[0] = 3
	//fmt.Printf("%+v\n",s)
	//fmt.Println("end")
	return
	//chain.SeeDoctor()
	//command.CommandRunTv()
	//iterator.IntertorUser()
	//mediator.TestMediator()
	//memento.TestMemento()
	//nullobject.TestNullObj()
	//observer.TestObserver()
	//state.TestState()
	//strategy.TestStrategy()
	//template.TestTemplate()
	//visitor.TestVisitor()
	//bridge.TestBridge()
	//decorator.TestDecorator()
	//facade.TestFacade()
	//flyweight.TestFlyweight()
	//proxy.TestProxy()
	//filesop.WriteFile()
	//filesop.WriteFileWithSize()
	//validate.ValidateStruct1()
	//time.TestTime()
	//log.TestLog()
	//os.TestOS()
	//getProductsHandler := http.HandlerFunc(httpptest.GetProducts)
	//http.Handle("/products", getProductsHandler)
	//http.ListenAndServe("localhost:8080", nil)
	//httpptest.DetectTimeout()
	//fmt.Println(13/ 2)
	//datastructures.TestMinheap()
	//regexptest.TestRegexp()
	//cond.TestCond()
//	atomic.TestConfig()
//	c := sarama.Config{
//		Admin: struct {
//			Retry struct {
//				Max     int
//				Backoff time.Duration
//			}
//			Timeout time.Duration
//		}{},
//		Net: struct {
//    MaxOpenRequests int
//    DialTimeout     time.Duration
//    ReadTimeout     time.Duration
//    WriteTimeout    time.Duration
//    TLS             struct {
//        Enable bool
//        Config *tls.Config
//    }
//    SASL struct {
//        Enable                   bool
//        Mechanism                sarama.SASLMechanism
//        Version                  int16
//        Handshake                bool
//        AuthIdentity             string
//        User                     string
//        Password                 string
//        SCRAMAuthzID             string
//        SCRAMClientGeneratorFunc func() sarama.SCRAMClient
//        TokenProvider            sarama.AccessTokenProvider
//        GSSAPI                   sarama.GSSAPIConfig
//    }
//    KeepAlive time.Duration
//    LocalAddr net.Addr
//    Proxy     struct {
//        Enable bool
//        Dialer proxy.Dialer
//    }
//}{},
//		Metadata: struct {
//    Retry struct {
//        Max         int
//        Backoff     time.Duration
//        BackoffFunc func(retries int, maxRetries int) time.Duration
//    }
//    RefreshFrequency time.Duration
//    Full             bool
//    Timeout          time.Duration
//}{},
//		Producer: struct {
//    MaxMessageBytes  int
//    RequiredAcks     sarama.RequiredAcks
//    Timeout          time.Duration
//    Compression      sarama.CompressionCodec
//    CompressionLevel int
//    Partitioner      sarama.PartitionerConstructor
//    Idempotent       bool
//    Return           struct {
//        Successes bool
//        Errors    bool
//    }
//    Flush struct {
//        Bytes       int
//        Messages    int
//        Frequency   time.Duration
//        MaxMessages int
//    }
//    Retry struct {
//        Max         int
//        Backoff     time.Duration
//        BackoffFunc func(retries int, maxRetries int) time.Duration
//    }
//    Interceptors []sarama.ProducerInterceptor
//}{},
//		Consumer: struct {
//    Group struct {
//        Session struct {
//            Timeout time.Duration
//        }
//        Heartbeat struct {
//            Interval time.Duration
//        }
//        Rebalance struct {
//            Strategy sarama.BalanceStrategy
//            Timeout  time.Duration
//            Retry    struct {
//                Max     int
//                Backoff time.Duration
//            }
//        }
//        Member struct {
//            UserData []byte
//        }
//    }
//    Retry struct {
//        Backoff     time.Duration
//        BackoffFunc func(retries int) time.Duration
//    }
//    Fetch struct {
//        Min     int32
//        Default int32
//        Max     int32
//    }
//    MaxWaitTime       time.Duration
//    MaxProcessingTime time.Duration
//    Return            struct {
//        Errors bool
//    }
//    Offsets struct {
//        CommitInterval time.Duration
//        AutoCommit     struct {
//            Enable   bool
//            Interval time.Duration
//        }
//        Initial   int64
//        Retention time.Duration
//        Retry     struct {
//            Max int
//        }
//    }
//    IsolationLevel sarama.IsolationLevel
//    Interceptors   []sarama.ConsumerInterceptor
//}{},
//		ClientID:          "",
//		RackID:            "",
//		ChannelBufferSize: 0,
//		Version:           sarama.KafkaVersion{},
//		MetricRegistry:    nil,
//	}
//	fmt.Println(c)

	//datastructures.TestTrie()
	//datastructures.TestSingleList()
	//for i :=0; i< 20; i++ {
	//	go single.GetInstance2()
	//}
	//fmt.Scanln()
	//file1 := &prototype.File{Name:"File1"}
	//file2 := &prototype.File{Name:"File2"}
	//file3 := &prototype.File{Name:"File3"}
	//
	//folder1 := &prototype.Folder{
	//	Childerens: []prototype.Inode{file1},
	//	Name: "Folder1",
	//}
	//folder2 := &prototype.Folder{
	//	Childerens: []prototype.Inode{folder1, file2, file3},
	//	Name:       "Folder2",
	//}
	//folder2.Print(" ")
	//cloneFolder := folder2.Clone()
	//fmt.Println("+++++++")
	//cloneFolder.Print(" ")
	//connections := make([]poolObject.IPoolObject, 0)
	//
	//for i:=0; i< 3; i++{
	//	c := &poolObject.Connection{Id:strconv.Itoa(i)}
	//	connections = append(connections, c)
	//}
	//pool , err := poolObject.InitPool(connections)
	//
	//if err != nil {
	//	log.Fatalf("Init Pool Error: %s", err)
	//}
	//conn1, err := pool.Loan()
	//
	//conn2, err := pool.Loan()
	//if err != nil {
	//	log.Fatal("Pool Load Error: %s", err)
	//}
	//
	//pool.Receive(conn1)
	//pool.Receive(conn2)
	//ak47 , _ := factorygun.GetGun("ak47")
	//maverick, _ := factorygun.GetGun("moverick")
	//
	//factorygun.PrintDetails(ak47)
	//factorygun.PrintDetails(maverick)
	//normalBuilder := builder.GetBuilder("normal")
	//
	//director := builder.NewDirector(normalBuilder)
	//normalHouse := director.BuildHouse()
	//fmt.Println(normalHouse)
	//adidasFactory,_ := abstractfactory.GetSportsFactory("adidas")
	//nikeFactory,_ := abstractfactory.GetSportsFactory("nike")
	//
	//adidasShoe := adidasFactory.MakeShoe()
	//adidasShort := adidasFactory.MakeShort()
	//
	//nikeShoe := nikeFactory.MakeShoe()
	//nikeShort := nikeFactory.MakeShort()
	//
	//
	//adidasShoe.PrintDetail()
	//adidasShort.PrintDetail()
	//nikeShoe.PrintDetail()
	//nikeShort.PrintDetail()
	//
	//a := alpha{
	//	name: "test",
	//}
	//
	////var some iAlpha
	//some := &beta{
	//	alpha: a,
	//}
	//some.alpha.work = some.work
	//some.common()
}