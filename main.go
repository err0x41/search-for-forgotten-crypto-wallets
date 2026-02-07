package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"qbt-hunter/internal/derivation"
	"qbt-hunter/internal/entropy"
	"qbt-hunter/internal/networks"
	"qbt-hunter/internal/scanner"

	"github.com/tyler-smith/go-bip39"
)

type Job struct {
	Mnemonic string
	Seed     []byte
}

var (
	counter    int64
	totalGen   int64 // Общий счетчик генераций
	totalScan  int64 // Общий счетчик проверок
	mu         sync.Mutex
	logChan    = make(chan string, 15)
)

func main() {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(100)

	registry := networks.GetRegistry()
	client := scanner.NewSmartClient()
	jobs := make(chan Job, 1000)

	fmt.Print("\033[H\033[2J")

	// Мониторинг интерфейса
	go func() {
		lastLog := "Initializing..."
		startTime := time.Now()
		for {
			select {
			case msg := <-logChan:
				lastLog = msg
			default:
			}
			
			uptime := time.Since(startTime).Round(time.Second)
			
			fmt.Printf("\033[H")
			fmt.Printf("\033[31m[ QBT GENERATION ]\033[0m            | \033[34m[ MULTI-CHAIN SCANNER ]\033[0m\n")
			fmt.Printf("Speed: %d m/s               | %-40s\n", counter, lastLog)
			fmt.Printf("Total Gen: %d               | Total Scan: %d\n", totalGen, totalScan)
			fmt.Printf("Uptime: %s            | Chains Loaded: %d\n", uptime, len(registry))
			fmt.Printf("----------------------------------------------------------------------\n")
			
			mu.Lock()
			counter = 0
			mu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	// Воркеры чека
	for w := 0; w < 3; w++ {
		go func(id int) {
			for job := range jobs {
				ethAddr, _ := derivation.GetEVMAddress(job.Seed)

				// Проверка баланса (каждый 500-й для стабильности RPC)
				if time.Now().UnixNano()%200 == 0 {
					mu.Lock()
					totalScan++
					mu.Unlock()
					
					netIdx := time.Now().Unix() % int64(len(registry))
					targetNet := registry[netIdx]

					if targetNet.Type == "evm" {
						bal, err := client.CheckEVMBalance(targetNet.RPCs[0], ethAddr)
						if err == nil {
							if bal > 0 {
								f, _ := os.OpenFile("JACKPOT.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
								if f != nil {
									fmt.Fprintf(f, "[%s] NET: %s | ADDR: %s | MNEM: %s | BAL: %f\n", 
										time.Now().Format("15:04:05"), targetNet.Name, ethAddr, job.Mnemonic, bal)
									f.Sync()
									f.Close()
								}
								logChan <- fmt.Sprintf("\033[42m\033[30m WIN: %s %s \033[0m", targetNet.Name, ethAddr[:6])
								fmt.Print("\a")
							} else {
								logChan <- fmt.Sprintf("%s: %s [0.00]", targetNet.Name, ethAddr[2:8])
							}
						}
					} else if targetNet.Name == "Solana" {
						solAddr, _ := derivation.GetSolanaAddress(job.Seed)
						logChan <- fmt.Sprintf("SOL: %s.. [SCAN]", solAddr[:6])
					}
				}
			}
		}(w)
	}

	// Генератор
	for {
		ent := entropy.GetSmartEntropy()
		mnem, _ := bip39.NewMnemonic(ent)
		seed := bip39.NewSeed(mnem, "")

		select {
		case jobs <- Job{Mnemonic: mnem, Seed: seed}:
			mu.Lock()
			counter++
			totalGen++
			mu.Unlock()
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
