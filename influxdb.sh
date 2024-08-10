#!/bin/bash
set -e

influx bucket create -n hammerBucketDetailed -o taxonim
influx bucket create -n hammerBucketIteration -o taxonim