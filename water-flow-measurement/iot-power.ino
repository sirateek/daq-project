#include <ESP8266WiFi.h>
#include <Logger.h>
#include "wifi.h"
#include "mqtthandler.h"
#include "config.h"

WiFiHandler wifiHandler(WIFI_SSID, WIFI_PASSWD);
MqttHandler mqttHandler(MQTT_HOST);

void setup()
{
    Serial.begin(115200);
    Logger::setLogLevel(Logger::VERBOSE);
    Logger::notice("SYS", "Starting up the iot-power");
}

void loop()
{
    // Connect and maintain the connection.
    wifiHandler.runCoroutine();
    if (wifiHandler.isConnected())
    {
        mqttHandler.runCoroutine();
    }
}
