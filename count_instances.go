package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Создание сессии AWS
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Создание клиента EC2
	svc := ec2.New(sess)

	// Получение списка регионов
	regionsOutput, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		fmt.Println("Ошибка при получении списка регионов:", err)
		return
	}

	regions := regionsOutput.Regions

	// Переменная для хранения общего количества инстансов
	totalInstances := 0

	// Цикл по регионам
	for _, region := range regions {
		regionName := *region.RegionName

		// Создание клиента EC2 для текущего региона
		svc := ec2.New(sess, &aws.Config{
			Region: aws.String(regionName),
		})

		// Получение списка инстансов EC2
		instancesOutput, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
		if err != nil {
			fmt.Printf("Ошибка при получении списка инстансов в регионе %s: %v\n", regionName, err)
			continue
		}

		instances := instancesOutput.Reservations

		// Добавление количества инстансов в текущем регионе к общему количеству
		totalInstances += len(instances)
	}

	// Вывод общего количества инстансов
	fmt.Println("Общее количество инстансов:", totalInstances)
}

