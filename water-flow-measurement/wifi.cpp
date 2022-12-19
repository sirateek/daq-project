#include "wifi.h"
#include "ESP8266WiFi.h"
#include "AceRoutine.h"
#include "Arduino.h"
#include <Logger.h>

WiFiHandler::WiFiHandler(const char *ssid, const char *password)
{
    _ssid = ssid;
    _password = password;
};

 bool WiFiHandler::isConnected() {
    return WiFi.isConnected();
 }

int WiFiHandler::runCoroutine()
{
    char a;
    COROUTINE_LOOP()
    {
        if (WiFi.isConnected())
        {
            Logger::notice("wifiHandler:wifiConnected :", "WiFi is connected. Skipping the connection process.");
            COROUTINE_DELAY(60000);
            continue;
        }
        Logger::warning("wifiHandler:wifiDisconnected :", "WiFi is disconnected, Reconnecting");
        Logger::notice("wifiHandler:info : SSID:", _ssid);
        Logger::verbose("wifiHandler:info : PASSWD:", _password);
        WiFi.begin(_ssid, _password);
        while (WiFi.status() != WL_CONNECTED)
        {
            Logger::notice("wifiHandler:connectWifi :", "Waiting for WiFi connection...");
            COROUTINE_DELAY(1000);
        }
        Logger::notice("wifiHandler:connectWifi :", "Connected to WiFi.");
        COROUTINE_DELAY(60000);
    }
};
