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

		// Если в регионе есть инстансы, выводим информацию
		if len(instances) > 0 {
			fmt.Println("Регион:", regionName)

			// Цикл по инстансам
			for _, reservation := range instances {
				for _, instance := range reservation.Instances {
					instanceID := *instance.InstanceId
					fmt.Println("Идентификатор инстанса:", instanceID)
					// Добавьте необходимую обработку других полей инстанса
				}
			}

			fmt.Println()
		}
	}
}

