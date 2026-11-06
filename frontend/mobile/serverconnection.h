#pragma once

#include <QWebSocket>
#include <emscripten/val.h>
#include <QJsonObject>
#include <QJsonDocument>

#include <QObject>

class ServerConnection
{
public:
    ServerConnection();
    void tryToReconnect() const;
    void sendTextMessage(QString message) const;
private:
    QWebSocket *socket;
    QUrl lastConnectedUrl = getWebsocketUrl();
    QUrl getWebsocketUrl();
    qint64 lastReconnectTime;
};
