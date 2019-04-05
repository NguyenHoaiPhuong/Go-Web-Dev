package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
)

func init() {
	fmt.Println("Init")

	debug.SetGCPercent(10)
}

func main() {
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	// err := metrics.RunCollector(metrics.DefaultConfig)
	// if err != nil {
	// 	// handle error
	// 	fmt.Println(err)
	// }

	cpuFile := StartCPUProfiling()
	memFile := StartMEMProfiling()
	runtime.MemProfileRate = 1

	// PrintMemUsage("Before alloc")
	// time.Sleep(10 * time.Second)

	a := make([]int, 1000000000)
	for i := range a {
		a[i] = i
	}

	// PrintMemUsage("After alloc")
	// time.Sleep(50 * time.Second)

	// debug.FreeOSMemory()
	// PrintMemUsage("After the first FreeOSMemory")
	// time.Sleep(5 * time.Second)

	a = nil
	debug.FreeOSMemory()
	RecordHeapProfile(memFile)
	// PrintMemUsage("After the second FreeOSMemory (Already Dealloc)")
	// time.Sleep(5 * time.Second)

	StopCPUProfiling(cpuFile)
	StopMEMProfiling(memFile)
	// time.Sleep(10 * time.Hour)
}

func f() {
	a := make([]int, 100000000)
	for i := range a {
		a[i] = i
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage(message string) {
	fmt.Println(message)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("HeapAlloc = %v MiB", bToMb(m.HeapAlloc))
	fmt.Printf("\tHeapSys = %v MiB", bToMb(m.HeapSys))
	fmt.Printf("\tHeapIdle = %v MiB", bToMb(m.HeapIdle))
	fmt.Printf("\tHeapReleased = %v MiB\n", bToMb(m.HeapReleased))
	fmt.Printf("MSpanInuse = %v MiB", bToMb(m.MSpanInuse))
	fmt.Printf("\tMSpanSys = %v MiB", bToMb(m.MSpanSys))
	fmt.Printf("\tGCSys = %v MiB\n", bToMb(m.GCSys))
	// fmt.Printf("\tNumGC = %v\n", m.NumGC)
	// fmt.Printf("\tMallocs = %v\n", m.Mallocs)
	// fmt.Printf("\tFrees = %v\n", m.Frees)
	fmt.Println("----------------------------------------------")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// StartCPUProfiling : start CPU profile
func StartCPUProfiling() *os.File {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	return f
}

// StopCPUProfiling : stop CPU profile
func StopCPUProfiling(f *os.File) {
	pprof.StopCPUProfile()
	f.Close()

	fmt.Println("Finished CPU profiling, to view the result run the following command : ")
	fmt.Println("go tool pprof cpu.prof")
	fmt.Println("Then select web, or type 'svg' to output the result as svg in your current directory")
}

// StartMEMProfiling : start Memory profile
func StartMEMProfiling() *os.File {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	return f
}

// RecordHeapProfile : write heap profile into file
func RecordHeapProfile(f *os.File) {
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

// StopMEMProfiling : stop memory profile
func StopMEMProfiling(f *os.File) {
	f.Close()

	fmt.Println("Finished memory profiling, to view the result run the following command : ")
	fmt.Println("go tool pprof mem.prof")
	fmt.Println("Then select web, or type 'svg' to output the result as svg in your current directory")
}
