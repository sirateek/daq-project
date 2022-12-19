# Happy E Water

## Team Members

1. Siratee KITTIWITCHAOWAKUL - 6310545426 - SKE - KU
2. Tanin PEWLUANGSAWAT - 6310545990 - SKE - KU

## Overview

Our project has a purpose to find the relation between temperature and water/electricity usage. The solution to measure the electricity is by using the SCT-013
and water by using the YF-S201 modules . Then we will get the temperature by using data from the third-party api.
Lastly, We will convert data in to the graph to show the relationship of how temperature effect used of electricity and water usage.

## Required libraries and tools

### Tools

1. NodeMCU 0.9
2. Kidbright (with micro python installed.)
3. SCT-013 - Clamp amp meters.
4. YF-S201 - Water flow rate meter.

### Libraries.

1. Micropython
2. Arduino MQTT
3. AceRoutine
4. ESP8266WiFi
5. GO 1.19
6. Vue.js


# Instructions for building and running

## Frontend     
> For dev run
```
make run-frontend 
```
> For Type-checking, complie and Minify for Production
```
make build-frontend
```
More detail [Frontend](./frontend/README.md)

## Backend
> Run program
```
make run-backend
```
