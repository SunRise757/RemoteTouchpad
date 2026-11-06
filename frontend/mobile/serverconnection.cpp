#include <QDateTime>

#include "serverconnection.h"

QUrl ServerConnection::getWebsocketUrl(){
    const auto location = emscripten::val::global("location");
    const auto protocol = location["protocol"].as<std::string>();
    const auto host = location["host"].as<std::string>();

    const std::string wsScheme = protocol == "https:" ? "wss://" : "ws://";
    return QUrl(QString::fromStdString(wsScheme + host + "/ws"));
}

ServerConnection::ServerConnection() {
    socket = new QWebSocket();
    QObject::connect(socket, &QWebSocket::connected,
        [this](){
        qDebug("Socket connected");
        socket->sendTextMessage("Hello from client!");
    });

    QObject::connect(socket, &QWebSocket::textMessageReceived,
        [this](QString message) {
        qDebug() << "Message received:" << message;
    });

    QObject::connect(socket, &QWebSocket::errorOccurred,
        [this](QAbstractSocket::SocketError error) {
        qDebug() << "Socket error:" << this->socket->errorString();
    });

    auto disconnection = QObject::connect(socket, &QWebSocket::disconnected,
        [this](){
            tryToReconnect();
    });

    lastConnectedUrl = getWebsocketUrl();
    socket->open(lastConnectedUrl);
}

void ServerConnection::tryToReconnect() const{

    if(socket->isValid())

        return;
    if(QDateTime::currentSecsSinceEpoch() < lastReconnectTime + 2)
        return;
    socket->open(lastConnectedUrl);
}

void ServerConnection::sendTextMessage(QString message) const{
    socket->sendTextMessage(message);

}