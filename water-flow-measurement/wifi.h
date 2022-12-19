#ifndef WIFI_H
#define WIFI_H
#include "AceRoutine.h"

class WiFiHandler : public ace_routine::Coroutine
{
public:
    WiFiHandler(const char *ssid, const char *password);
    int runCoroutine() override;
    bool isConnected();

private:
    const char *_ssid;
    const char *_password;
};

#endif
