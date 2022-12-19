#include "mqtthandler.h"
#include "Arduino.h"
#include <Logger.h>
#include <MQTT.h>
#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <EmonLib.h>

#define flow_sensor 14

WiFiClient espClient;
PubSubClient client(espClient);
EnergyMonitor emon1;

void callback(char* topic, byte* payload, unsigned int length) {
  Serial.print("Message arrived [");
  Serial.print(topic);
  Serial.print("] ");
  String msg = "";
  int i=0;
  while (i<length) msg += (char)payload[i++];
  Serial.println(msg);
}

MqttHandler::MqttHandler(const char *host)
{
    _host = host;
    client.setServer("iot.cpe.ku.ac.th", 1883);
    client.setCallback(callback);
    pinMode(flow_sensor, INPUT);
    pinMode(BUILTIN_LED, OUTPUT);
}

int MqttHandler::runCoroutine()
{
    char a;
    COROUTINE_LOOP()
    {
        if (!client.connected()) {
            Logger::notice("MQTT: ", "Client is not connected");
            if (client.connect("ESP8266Client", "b6310545426", "siratee.k@ku.th")) {
                Logger::notice("MQTT: ", "Client is connected");
                client.subscribe("b6310545426/test/");
            } else {
                Logger::warning("MQTT: ", "Connection Failed");
            }
        }
        client.loop();

        uint32_t pulse = pulseIn(flow_sensor,HIGH);
        float flow;
        if(pulse >= 1){
            float Hz = 1/(2*pulse*pow(10,-6));
            flow = 7.2725*(float)Hz + 3.2094;
        } else {
            flow = 0;
        }

        char data[100];
        Serial.println(flow/60);
        sprintf(data, "{\"flowRate\": %.2f}", flow/60);
        digitalWrite(BUILTIN_LED, LOW);
        client.publish("b6310545426/happye/water", data);
        digitalWrite(BUILTIN_LED, HIGH);
        Logger::notice("Data: ", data);
        COROUTINE_DELAY(1000);
    }
}