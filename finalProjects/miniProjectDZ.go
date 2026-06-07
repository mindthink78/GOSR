package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Sensor struct {
	ID    int
	Type  string
	Value float64
	LatX  float64
	LonY  float64
}

func StartPressureSensors(
	ctx context.Context,
	wg *sync.WaitGroup,
	dataChannel chan Sensor,
	count int,

) {
	for i := 1; i < count+1; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for {
				value := 700 + rand.Intn(101)
				sensor := Sensor{
					ID:    id,
					Type:  "pressure",
					Value: float64(value),
					LatX:  55.0,
					LonY:  37.0,
				}
				select {
				case <-ctx.Done():
					fmt.Println("--->>>Работа датчика ДАВЛЕНИЯ завершена")
					return
				case dataChannel <- sensor:
					fmt.Println("Данные датчика ДАВЛЕНИЯ зафиксированы")

					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}
}

func StartHumiditySensors(
	ctx context.Context,
	wg *sync.WaitGroup,
	dataChannel chan Sensor,
	count int,
) {
	for i := 1; i < count+1; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for {
				value := 1 + rand.Intn(101)
				sensor := Sensor{
					ID:    id,
					Type:  "humidity",
					Value: float64(value),
					LatX:  55.0,
					LonY:  37.0,
				}
				select {
				case <-ctx.Done():
					fmt.Println("--->>>Работа датчика ВЛАЖНОСТИ завершена")
					return
				case dataChannel <- sensor:
					fmt.Println("Данные датчика ВЛАЖНОСТИ зафиксированы")

					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}
}
func StartSeismicSensors(
	ctx context.Context,
	wg *sync.WaitGroup,
	dataChannel chan Sensor,
	count int,
) {
	for i := 1; i < count+1; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for {
				value := rand.Intn(11)
				sensor := Sensor{
					ID:    id,
					Type:  "seismic",
					Value: float64(value),
					LatX:  55.0,
					LonY:  37.0,
				}
				select {
				case <-ctx.Done():
					fmt.Println("--->>>Работа датчика СЕЙСМИЧЕСКОЙ АКТИВНОСТИ завершена")
					return
				case dataChannel <- sensor:
					fmt.Println("Данные датчика СЕЙСМИЧЕСКОЙ АКТИВНОСТИ зафиксированы")

					time.Sleep(500 * time.Millisecond)
				}
			}
		}(i)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	pressureContext, pressureCancel := context.WithCancel(context.Background())
	humidityContext, humidityCancel := context.WithCancel(context.Background())
	seismicContext, seismicCancel := context.WithCancel(context.Background())

	dataChannel := make(chan Sensor)

	go func() {
		time.Sleep(1 * time.Second)
		pressureCancel()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		humidityCancel()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		seismicCancel()
	}()

	StartPressureSensors(pressureContext, wg, dataChannel, 3)
	StartHumiditySensors(humidityContext, wg, dataChannel, 3)
	StartSeismicSensors(seismicContext, wg, dataChannel, 3)

	go func() {
		wg.Wait()
		close(dataChannel)
	}()

	for value := range dataChannel {
		fmt.Println(value)
	}
	fmt.Println("--->>>Метеоцентр закрылся!")
}
