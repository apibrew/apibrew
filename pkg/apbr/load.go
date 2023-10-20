package apbr

import (
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/pkg/apbr/flags"
	"github.com/apibrew/apibrew/pkg/formats/apply"
	"github.com/spf13/cobra"
	"sync"
	"sync/atomic"
	"text/tabwriter"
	"time"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load test",
	Long:  `Load test apibrew`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if loadNumberOfRequests == nil {
			return errors.New("please specify number of requests")
		}

		if *loadNumberOfRequests <= 0 {
			return errors.New("number of requests must be greater than 0")
		}

		if loadConcurrency == nil {
			return errors.New("please specify concurrency")
		}

		if *loadConcurrency <= 0 {
			return errors.New("concurrency must be greater than 0")
		}

		if loadAction == nil {
			return errors.New("please specify action")
		}

		if loadPayloadFile == nil {
			return errors.New("please specify payload file")
		}

		parseRootFlags(cmd)

		// locating payload file

		var loadFn func() error

		switch *loadAction {
		case "APPLY":
			loadFn = loadApply(cmd, args, *loadPayloadFile)
		default:
			return errors.New("invalid action or not implemented yet")
		}

		var sem = make(chan int, *loadConcurrency)
		var errorCount = 0
		var timeTakenNs []int64
		var sumTimeTakenNs int64 = 0
		overallStartTime := time.Now()
		var executedCount int32 = 0

		fmt.Println("Load test results")
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 1, ' ', tabwriter.TabIndent)
		_, _ = fmt.Fprintf(w, "Number of requests\t%d\n", *loadNumberOfRequests)
		_, _ = fmt.Fprintf(w, "Concurrency\t%d\n", *loadConcurrency)

		var printMutex = sync.Mutex{}
		var lastPrintTime = time.Now()
		var mustPrint = func() {
			overallDiff := time.Now().Sub(overallStartTime).Milliseconds()

			var sumTimeTaken int64 = sumTimeTakenNs / int64(time.Millisecond)
			_, _ = fmt.Fprintf(w, "Time taken [sum]\t%dms\n", sumTimeTaken)
			_, _ = fmt.Fprintf(w, "Time taken \t%dms\n", overallDiff)
			_, _ = fmt.Fprintf(w, "Average time taken\t%dms\n", sumTimeTaken/int64(executedCount))
			_, _ = fmt.Fprintf(w, "Average time taken\t%dmks\n", (sumTimeTakenNs/int64(time.Microsecond))/int64(executedCount))
			_, _ = fmt.Fprintf(w, "Error count\t%d\n", errorCount)
			// request per second
			_, _ = fmt.Fprintf(w, "Request per second\t%d\n", int64(executedCount*1000)/overallDiff)
			fmt.Println("")
			fmt.Println("")
			_ = w.Flush()
		}
		var printResults = func() {
			if time.Now().Sub(lastPrintTime).Seconds() < 3 {
				return
			}

			lastPrintTime = time.Now()
			printMutex.Lock()
			defer printMutex.Unlock()
			mustPrint()
		}

		var wg = sync.WaitGroup{}

		for i := 0; i < int(*loadNumberOfRequests); i++ {
			sem <- 1
			wg.Add(1)
			go func() {
				defer func() { <-sem; wg.Done() }()

				startTime := time.Now()
				if err := loadFn(); err != nil {
					errorCount++
				}
				endTime := time.Now()
				diff := endTime.Sub(startTime).Nanoseconds()
				timeTakenNs = append(timeTakenNs, diff)
				sumTimeTakenNs += diff

				atomic.AddInt32(&executedCount, 1)

				printResults()
			}()
		}

		wg.Wait()

		mustPrint()

		// write results

		return nil
	},
}

func loadApply(cmd *cobra.Command, args []string, payloadFile string) func() error {
	applier := apply.NewApplier(GetClient(), false, false, false, flags.OverrideConfig{})

	return func() error {
		return applier.Apply(cmd.Context(), payloadFile, "")
	}
}

var loadNumberOfRequests *int32
var loadConcurrency *int32
var loadTimeLimit *int32
var loadTimeout *int32
var loadAction *string
var loadPayloadFile *string

func init() {
	loadNumberOfRequests = loadCmd.PersistentFlags().Int32P("requests", "n", 0, "Number of requests to perform")
	loadConcurrency = loadCmd.PersistentFlags().Int32P("concurrency", "c", 0, "Number of multiple requests to make at a time")
	loadTimeLimit = loadCmd.PersistentFlags().Int32P("timelimit", "t", 0, "Seconds to max. to spend on benchmarking\n                  This implies -n 50000")
	loadTimeout = loadCmd.PersistentFlags().Int32P("timeout", "s", 0, "Seconds to max. wait for each response\n                    Default is 30 seconds")
	loadAction = loadCmd.PersistentFlags().StringP("action", "a", "APPLY", "Load action, indicates which endpoint will be called\n                    Default is 30 seconds")
	loadPayloadFile = loadCmd.PersistentFlags().StringP("payload", "p", "", "Load payload file, indicates which payload will be sent\n                    Default is 30 seconds")
}
