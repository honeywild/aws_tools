


#!/bin/bash

regions=$(aws ec2 describe-regions --query 'Regions[].RegionName' --output text)

for region in $regions; do
    echo "Region: $region"
    aws ec2 describe-instances --region $region --query 'Reservations[].Instances[].{InstanceId: InstanceId, Region: Placement.AvailabilityZone}' --output table
    echo
done
