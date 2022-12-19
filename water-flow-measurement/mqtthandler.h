#ifndef MQTT_HANDLER_H
#define MQTT_HANDLER_H
#include "AceRoutine.h"

class MqttHandler : public ace_routine::Coroutine
{
public:
    MqttHandler(const char *host);
    int runCoroutine() override;

private:
    const char *_host;
};

#endif